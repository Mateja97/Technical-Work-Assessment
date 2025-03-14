package server

import (
	"alluvial-task/config"
	"alluvial-task/handler"
	"context"
	"errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var srv *http.Server

func Init(handler *handler.HttpHandler, ethClientNum int) {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	mux.HandleFunc("/live", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	})

	mux.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		if ethClientNum == 0 {
			w.WriteHeader(http.StatusServiceUnavailable)
			_, _ = w.Write([]byte("Service not ready"))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	})
	mux.HandleFunc("/getBalance/", handler.GetBalance)
	srv = &http.Server{
		Addr:    config.ServerAddress(),
		Handler: mux,
	}
}

func Start() {
	log.Println("Starting server on: ", config.ServerAddress())
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("server error: %v", err)
	}
}

func Shutdown(ctx context.Context) error {
	if err := srv.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
