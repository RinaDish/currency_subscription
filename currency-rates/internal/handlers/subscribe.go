package handlers

import (
	"context"
	"net/http"

	"github.com/RinaDish/currency-rates/internal/repo"
	"go.uber.org/zap"
)

type Db interface {
	SetEmail(ctx context.Context, email string) error
	GetEmails(ctx context.Context) ([]repo.Email, error)
}

type SubscribeHandler struct {
	l *zap.SugaredLogger
	r Db
}

func NewSubscribeHandler(l *zap.SugaredLogger, r Db) SubscribeHandler {

	return SubscribeHandler{
		l: l,
		r: r,
	}
}

func (h SubscribeHandler) CreateSubscription(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Form parse error", http.StatusConflict)
		return
	}

	formData := r.Form

	email := formData.Get("email")

	err = h.r.SetEmail(r.Context(), email)
	responseStatus := http.StatusOK
	if err != nil {
		responseStatus = http.StatusConflict
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
}

func (h SubscribeHandler) GetEmailsList(w http.ResponseWriter, r *http.Request) {
	emails, err := h.r.GetEmails(r.Context())
	if err != nil {
		h.l.Error(err)
	}
	for _, email := range emails {
		h.l.With("id", email.ID).Info(email.Email)
	}
}
