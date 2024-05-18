package app

import (
	"context"
	"net/http"

	"github.com/RinaDish/currency-rates/internal/clients"
	"github.com/RinaDish/currency-rates/internal/handlers"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type App struct {
	cfg Config
	l   *zap.SugaredLogger
}

func NewApp(c Config, l *zap.SugaredLogger) *App {
	return &App{
		cfg: c,
		l:   l,
	}
}

func (app *App) Run(ctx context.Context) error {
	nbuClient := clients.NewNBUClient(app.l)
	privatClient := clients.NewPrivatClient(app.l)

	h := handlers.NewRateHandler(app.l, []handlers.RateClient{nbuClient, privatClient})

	r := chi.NewRouter()
	r.Get("/rate", h.GetCurrentRate)
	app.l.Info("app run")
	return http.ListenAndServe(app.cfg.Address, r)
}
