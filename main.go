package main

import (
	"os"
	"os/signal"

	"github.com/papijo/go-api-gateway/application"
)

func main() {
	// Start the application
	e := application.StartApiGatewayService()

	// Gracefully stop the application
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	application.StopApiGatewayService(e)

}
