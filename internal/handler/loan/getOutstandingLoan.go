package loan

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/dannyhenry/billing/internal/model/api_response"
	"github.com/labstack/echo/v4"
)

// GetOutstandingLoan godoc
// @Summary Get loan outstanding amount
// @Description Returns remaining outstanding amount for a loan.
// @Tags Loans
// @Produce json
// @Security BearerAuth
// @Param loan_id path int true "Loan ID"
// @Success 200 {object} api_response.SuccessResponse{data=int64}
// @Failure 500 {object} api_response.ErrorResponse{errors=[]string}
// @Failure 400 {object} api_response.ErrorResponse{errors=[]string}
// @Router /loans/{loan_id}/outstanding [get]
func (h *LoanHandler) GetOutstandingLoan(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 6*time.Second)
	defer cancel()

	loanID, err := strconv.ParseInt(c.Param("loan_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, api_response.BuildErrorResponse("invalid loan id",
			http.StatusBadRequest, err, nil))
	}

	data, err := h.LoanScheduleService.GetOutstanding(ctx, loanID)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			api_response.BuildErrorResponse("Data Not Found", http.StatusBadRequest,
				err, nil))
	}

	return c.JSON(http.StatusOK,
		api_response.BuildSuccessResponse("Data Found", http.StatusOK, data))
}
