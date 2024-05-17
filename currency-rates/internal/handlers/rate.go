package handlers

import (
	"net/http"

	"go.uber.org/zap"
)

type RateHandler struct {
	l   *zap.SugaredLogger
}

func NewRateHandler(l *zap.SugaredLogger) RateHandler {

	return RateHandler{
		l: l,
	}
}

func (h RateHandler) GetCurrentRate (w http.ResponseWriter, r *http.Request) {
	h.l.Info("CHTO UGODNO")
}