package handler

import (
	"alluvial-task/metrics"
	"alluvial-task/service"
	"log"
	"net/http"
	"strings"
	"time"
)

type HttpHandler struct {
	balanceService *service.BalanceService
}

func NewHttpHandler(balanceService *service.BalanceService) *HttpHandler {
	return &HttpHandler{balanceService: balanceService}
}

func (h *HttpHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	defer func() {
		duration := time.Since(startTime).Seconds()
		metrics.BalanceRequestDuration.Observe(duration)
	}()

	metrics.BalanceRequestsCounter.Inc()
	address := strings.TrimPrefix(r.URL.Path, "/getBalance/")
	if address == "" {
		http.Error(w, "Ethereum address is required", http.StatusBadRequest)
		return
	}

	if len(address) != 42 || !strings.HasPrefix(address, "0x") {
		http.Error(w, "Invalid Ethereum address", http.StatusBadRequest)
		return
	}

	balance, err := h.balanceService.GetBalance(address)
	if err != nil {
		metrics.BalanceRequestErrors.Inc()
		log.Println("error getBalance: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte(balance))
	if err != nil {
		metrics.BalanceRequestErrors.Inc()
		log.Println("error w.Write: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
