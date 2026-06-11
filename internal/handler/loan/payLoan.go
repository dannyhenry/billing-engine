package loan

import (
	"context"
	"net/http"
	"time"

	"github.com/dannyhenry/billing/internal/model/api_response"
	"github.com/dannyhenry/billing/internal/model/loan"
	"github.com/labstack/echo/v4"
)

func (h *LoanHandler) PayLoan(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	var req loan.PaymentLoanRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, api_response.ApplicationError("invalid request : ",
			http.StatusBadRequest, err, nil))
	}

	result, err := h.LoanService.PayLoans(ctx, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			api_response.ApplicationError("Data Not Found", http.StatusBadRequest,
				err, nil))
	}

	return c.JSON(http.StatusOK,
		api_response.BuildSuccessResponse("Data Found", http.StatusOK, result))

}
