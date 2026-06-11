package rest

import (
	"github.com/labstack/echo/v4"

	"github.com/dannyhenry/billing/internal/handler/loan"
)

func RegisterRoutes(e *echo.Echo,
	loandHandler *loan.LoanHandler) {
	loan := e.Group("/loans")

	//loan.POST("", loandHandler.CreateLoan)
	//
	//loan.GET("/:id/schedule", loandHandler.GetSchedule)
	loan.GET("/:loan_id/outstanding", loandHandler.GetOutstandingLoan)
	loan.GET("/:loan_id/delinquent", loandHandler.IsDelinquent)
	loan.POST("/payments", loandHandler.PayLoan)
}
