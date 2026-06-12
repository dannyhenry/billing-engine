package loan

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/dannyhenry/billing/internal/model/api_response"
	"github.com/dannyhenry/billing/internal/model/loan"
	"github.com/labstack/echo/v4"
)

// MakePayment godoc
// @Summary Make loan payment
// @Description Pay one or multiple installments. Payment amount must be a multiple of weekly installment amount.
// @Tags Loans
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body loan.PaymentLoanRequest true "Payment request"
// @Success 200 {object} api_response.SuccessResponse{data=loan.PaymentLoanResponse}
// @Failure 500 {object} api_response.ErrorResponse{errors=[]string}
// @Failure 400 {object} api_response.ErrorResponse{errors=[]string}
// @Router /loans/payments [post]
func (h *LoanHandler) PayLoan(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	var req loan.PaymentLoanRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, api_response.BuildErrorResponse("invalid request : ",
			http.StatusBadRequest, err, nil))
	}

	fmt.Printf("%+v", req)
	result, err := h.LoanService.PayLoans(ctx, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			api_response.BuildErrorResponse("Data Not Found", http.StatusBadRequest,
				err, nil))
	}

	return c.JSON(http.StatusOK,
		api_response.BuildSuccessResponse("Data Found", http.StatusOK, result))

}
