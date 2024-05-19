package app

import (
	"context"
	"net/http"

	"github.com/RinaDish/currency-rates/internal/clients"
	"github.com/RinaDish/currency-rates/internal/handlers"
	"github.com/RinaDish/currency-rates/internal/repo"
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
	
	adminRepository, err := repo.NewAdminRepository(app.cfg.DBUrl, app.l)
	if err != nil { return err }
	
	ratesHandler := handlers.NewRateHandler(app.l, []handlers.RateClient{nbuClient, privatClient})
	subscriptionHandler := handlers.NewSubscribeHandler(app.l, adminRepository) 

	r := chi.NewRouter()
	r.Get("/rate", ratesHandler.GetCurrentRate)
	r.Post("/subscribe", subscriptionHandler.CreateSubscription)
	r.Get("/list", subscriptionHandler.GetEmailsList)
	app.l.Info("app run")
	return http.ListenAndServe(app.cfg.Address, r)
}
