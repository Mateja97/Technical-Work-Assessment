package main

import (
	_ "alluvial-task/metrics"

	"alluvial-task/client"
	"alluvial-task/config"
	"alluvial-task/handler"
	"alluvial-task/server"
	"alluvial-task/service"
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	ethClientPool := client.NewEthClientPool(config.EthClients())

	cacheMap := new(sync.Map)
	balanceService := service.NewBalanceService(ethClientPool, cacheMap)

	httpHandler := handler.NewHttpHandler(balanceService)

	server.Init(httpHandler, ethClientPool.Len())

	go func() {
		server.Start()
	}()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM)

	<-stopChan

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}

	log.Println("Server gracefully stopped")
}
