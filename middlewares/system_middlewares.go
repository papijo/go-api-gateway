package middlewares

import (
	"crypto/subtle"
	"fmt"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/papijo/go-api-gateway/pkg/logger"
)

func CORSMiddleware() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowCredentials: true,
		AllowHeaders:     []string{"Authorization", "Content-Type"},
	})
}

func BasicAuthMiddleware() echo.MiddlewareFunc {
	auth_username := os.Getenv("BASIC_AUTH_USERNAME")
	auth_password := os.Getenv("BASIC_AUTH_PASSWORD")
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if subtle.ConstantTimeCompare([]byte(username), []byte(auth_username)) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(auth_password)) == 1 {
			return true, nil
		}
		return false, nil
	})
}

func LoggerConfigMiddleware() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogRemoteIP: true,
		LogLatency:  true,
		LogURI:      true,
		LogMethod:   true,
		LogError:    true,
		BeforeNextFunc: func(ctx echo.Context) {
			ctx.Set("startTime", time.Now())

		},
		LogValuesFunc: func(ctx echo.Context, values middleware.RequestLoggerValues) error {
			// startTime, _ := ctx.Get("startTime").(time.Time)

			logger.Logger.Info().
				Str("method", values.Method).
				Str("uri", values.URI).
				Int("status", values.Status).
				Str("latency", values.Latency.String()).
				Str("error", fmt.Sprintf("%v", values.Error)).
				Str("ip-address", values.RemoteIP).
				Msg("Request logged")

			return nil
		},
	})

}
