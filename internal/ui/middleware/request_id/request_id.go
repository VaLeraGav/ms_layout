package request_id

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type contextIdKey string

const reqIDKey contextIdKey = "request_id"

// Middleware для генерации и установки request_id
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Генерируем новый request_id
		requestID := uuid.New().String()

		// Устанавливаем request_id в контекст

		ctx := context.WithValue(r.Context(), reqIDKey, requestID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

// Функция для получения request_id из контекста
func GetReqID(ctx context.Context) string {
	if reqID, ok := ctx.Value(reqIDKey).(string); ok {
		return reqID
	}
	return ""
}
