package main

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"os/signal"
	"syscall"
	"zmeet/pkg/api/gin"
	"zmeet/pkg/logger"
	"zmeet/pkg/store"
)

func main() {
	customLogger := logger.NewLogger(logger.DEBUG)
	customLogger.Info("Server started.")

	newStore := store.NewStore(customLogger)

	go gin.Initilize(newStore)

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTSTP, syscall.SIGTERM)

	// Block until a signal is received
	<-signalChannel
}
