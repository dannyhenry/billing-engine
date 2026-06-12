package rest

import (
	"github.com/dannyhenry/billing/common/config"
	customMiddleware "github.com/dannyhenry/billing/internal/infrastructure/middleware"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/dannyhenry/billing/internal/handler/loan"
)

func RegisterRoutes(e *echo.Echo, loandHandler *loan.LoanHandler) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	loan := e.Group("/loans")
	loan.Use(customMiddleware.TokenAuth(config.GetConfig().APIToken))
	loan.GET("/:loan_id/outstanding", loandHandler.GetOutstandingLoan)
	loan.GET("/:loan_id/delinquent", loandHandler.IsDelinquent)
	loan.POST("/payments", loandHandler.PayLoan)
}
