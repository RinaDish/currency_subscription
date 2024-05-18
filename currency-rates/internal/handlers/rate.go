package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"go.uber.org/zap"
)

type RateClient interface {
	GetDollarRate(ctx context.Context) (float64, error) 
}

type RateHandler struct {
	l   *zap.SugaredLogger
	rateClients []RateClient
}

func NewRateHandler(l *zap.SugaredLogger, r []RateClient) RateHandler {

	return RateHandler{
		l: l,
		rateClients: r,
	}
}

func (h RateHandler) GetCurrentRate(w http.ResponseWriter, r *http.Request) {
	var rate float64
	var err error
	for _, c := range h.rateClients {
		ctx, _ := context.WithTimeout(context.Background(), 500 * time.Millisecond)
		rate, err = c.GetDollarRate(ctx)
		if err == nil { break }
	}

	w.Header().Set("Content-Type", "application/json")
	if err == nil { 
		w.WriteHeader(http.StatusOK)
		strRate := strconv.FormatFloat(rate, 'f', -1, 64)
		w.Write([]byte(strRate))
		return
	} 
	w.WriteHeader(http.StatusBadRequest)
	
}