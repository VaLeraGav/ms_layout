package handlers

import (
	"errors"
	"net/http"
	"time"

	"gitlab.toledo24.ru/web/ms_layout/internal/entities"
	"gitlab.toledo24.ru/web/ms_layout/internal/ui/respond"
	"gitlab.toledo24.ru/web/ms_layout/pkg/validate"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog"
)

type User interface {
	Find(string) (*entities.User, error)
}

var (
	ErrEmailNotValid          = errors.New("the email is not valid")
	ErrNotFindUser            = errors.New("not find user")
	ErrNotStructureResponse1C = errors.New("not find the structure of the response")
)

func GetUser(log *zerolog.Logger, storeUser User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		time.Sleep(5 * time.Second)
		companyLog := log.With().Str("method", "DeleteUser").Logger()

		email := chi.URLParam(r, "email")

		if email == "" || !validate.IsValidEmail(email) {
			respond.ErrorHandle(w, r, http.StatusBadRequest, ErrEmailNotValid)
			return
		}

		user, err := storeUser.Find(email)
		if err != nil {
			companyLog.Error().Err(err).Msg("failed to delete company")
			respond.ErrorHandle(w, r, http.StatusInternalServerError, ErrNotFindUser)
			return
		}

		respond.Respond(w, r, http.StatusOK, user)
	}
}
