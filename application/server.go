package application

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/papijo/go-api-gateway/config"
	"github.com/papijo/go-api-gateway/middlewares"
	errorservice "github.com/papijo/go-api-gateway/pkg/error_service"
	"github.com/papijo/go-api-gateway/pkg/response"
)

func StartApiGatewayService() *echo.Echo {

	err := config.LoadEnvironmentVariables()
	if err != nil {
		log.Fatal("Error loading environment variables", err)
	}

	e := echo.New()

	//Middlewares
	e.Use(middlewares.ErrorHandlerMiddleware())
	e.Use(middleware.Recover())
	e.Use(middlewares.CORSMiddleware())
	e.Use(middlewares.LoggerConfigMiddleware())

	e.GET("/", func(ctx echo.Context) error {
		app_identity := echo.Map{
			"Application Name":        "Go API Gateway",
			"Application Description": "API Gateway Server built with Golang Echo and PostgreSQL",
			"Application Owner":       "Ebhota Jonathan",
			"Application Version":     "1.0.0",
			"Application Engineer":    "Ebhota Jonathan",
		}

		if true {
			return errorservice.NewError(http.StatusBadRequest, "Testing error handling from a route", nil)
		}

		return ctx.JSON(http.StatusOK, response.SuccessResponse("Application Home Page", app_identity))
	})

	server := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		ReadTimeout:  5 * time.Minute,
		WriteTimeout: 5 * time.Minute,
	}

	go func() {
		err := e.StartServer(server)
		if err != nil {
			log.Fatalln("❌ Server start error: ", err, "Shutting down the server...")
		}
	}()

	log.Println("⚡️🚀 API Gateway Engine - Started")

	return e

}

func StopApiGatewayService(e *echo.Echo) error {
	time.Sleep(1 * time.Second)
	log.Println("⚡️🚀 API Gateway Engine - Stopping")

	time.Sleep(5 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err := e.Shutdown(ctx)

	if err != nil {
		log.Fatalf("❌ Error during shutdown: %v", err)
		return err
	}

	log.Println("✅ API Gateway Engine stopped gracefully.")
	return nil
}
