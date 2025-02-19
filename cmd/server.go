package server

import (
	"os"
	"os/signal"
	"syscall"
	"zmeet/pkg/api/gin"
	"zmeet/pkg/logger"
	"zmeet/pkg/store"
)

func Start() {
	logger.NewLogger(logger.DEBUG)
	logger.Info("Starting server")

	newStore := store.NewStore()

	go gin.Initilize(newStore)

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTSTP, syscall.SIGTERM)

	// Block until a signal is received
	<-signalChannel
}
