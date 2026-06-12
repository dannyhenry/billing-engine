package rest

import (
	"github.com/dannyhenry/billing/common/config"
	"github.com/labstack/echo/v4"

	customMiddleware "github.com/dannyhenry/billing/internal/infrastructure/middleware"

	"github.com/dannyhenry/billing/internal/handler/loan"
)

func RegisterRoutes(e *echo.Echo,
	loandHandler *loan.LoanHandler) {
	e.Use(customMiddleware.TokenAuth(config.GetConfig().APIToken))
	loan := e.Group("/loans")

	//loan.POST("", loandHandler.CreateLoan)
	//
	//loan.GET("/:id/schedule", loandHandler.GetSchedule)
	loan.GET("/:loan_id/outstanding", loandHandler.GetOutstandingLoan)
	loan.GET("/:loan_id/delinquent", loandHandler.IsDelinquent)
	loan.POST("/payments", loandHandler.PayLoan)
}
