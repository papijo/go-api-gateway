package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	errorservice "github.com/papijo/go-api-gateway/pkg/error_service"
	"github.com/papijo/go-api-gateway/pkg/logger"
	"github.com/papijo/go-api-gateway/pkg/response"
)

// ErrorHandler handles all errors, including 404.
func ErrorHandlerMiddleware() echo.MiddlewareFunc {
	timer := time.Now()
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			err := next(ctx)

			if err != nil {
				if appErr, ok := err.(*errorservice.AppError); ok {
					logger.ErrorLogger.Error().
						Str("method", ctx.Request().Method).
						Str("uri", ctx.Request().RequestURI).
						Int("status", appErr.Code).
						Str("latency", time.Since(timer).String()).
						Str("error", fmt.Sprintf("%v", appErr.Error())).
						Str("ip-address", ctx.RealIP()).
						Msg(appErr.Message)

					return ctx.JSON(appErr.Code, response.ErrorResponse(appErr.Message, appErr.Details))
				}

				if httpError, ok := err.(*echo.HTTPError); ok {
					if httpError.Code == http.StatusNotFound {
						// Log the 404 error
						logger.ErrorLogger.Error().
							Str("method", ctx.Request().Method).
							Str("uri", ctx.Request().RequestURI).
							Int("status", http.StatusNotFound).
							Str("latency", time.Since(timer).String()).
							Str("error", "API not found or still in construction").
							Str("ip-address", ctx.RealIP()).
							Msg("Resource Not Found")

						// Custom response for 404
						return ctx.JSON(http.StatusNotFound, response.ErrorResponse("Resource not found", nil))
					}

				}

				// Generic Logging
				logger.ErrorLogger.Error().
					Err(err).
					Msg("Error handled in custom middleware")

				// Optionally return a custom response or propagate the error
				return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse("We are working on a fix. Please try again shortly", nil))
			}

			return nil
		}
	}
}
