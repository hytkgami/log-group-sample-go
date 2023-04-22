package controllers

import (
	"net/http"

	"github.com/hytkgami/log-group-sample-go/internal"
)

type HealthCheckController struct{}

func NewHealthCheckController() *HealthCheckController {
	return &HealthCheckController{}
}

func (c *HealthCheckController) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	internal.LogInfof(r.Context(), "pong")
	w.Write([]byte("pong"))
}
