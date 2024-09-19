package logger

import (
	"errors"
	"net/http"
	"runtime/debug"
	"time"

	"gitlab.toledo24.ru/web/ms_layout/internal/ui/middleware/request_id"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

func New(log *zerolog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Info().
				Str("method", r.Method).
				Str("url", r.URL.RequestURI()).
				Str("user_agent", r.UserAgent()).
				Str("request_id", request_id.GetReqID(r.Context())).
				Msg("Incoming request")

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			defer func() {
				if rec := recover(); rec != nil {
					log.Error().
						Int("status", ww.Status()).
						Interface("recover_info", rec).
						Bytes("debug_stack", debug.Stack()).
						Str("request_id", request_id.GetReqID(r.Context())).
						Msg("log system error")

					http.Error(ww, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					return
				}

				var login *zerolog.Event
				status := ww.Status()
				if status >= 200 && status < 300 {
					login = log.Info()
				} else {
					login = log.Err(errors.New(""))
				}
				login.
					Int("status", ww.Status()).
					Int("bytes", ww.BytesWritten()).
					Str("elapsed_ms", time.Since(t1).String()).
					Str("request_id", request_id.GetReqID(r.Context())).
					Msg("Request processed")
			}()

			next.ServeHTTP(ww, r)
		})

		return http.HandlerFunc(fn)
	}
}
