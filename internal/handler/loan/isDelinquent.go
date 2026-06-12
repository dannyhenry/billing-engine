package loan

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/dannyhenry/billing/internal/model/api_response"
	"github.com/labstack/echo/v4"
)

// IsDelinquent godoc
// @Summary Check loan delinquency
// @Description Returns whether borrower is delinquent. Borrower is delinquent when due unpaid schedules are greater than or equal to 2.
// @Tags Loans
// @Produce json
// @Security BearerAuth
// @Param loan_id path int true "Loan ID"
// @Param mock_date query string false "Mock date for testing, format YYYY-MM-DD" example(2026-01-25)
// @Success 200 {object} api_response.SuccessResponse{data=bool}
// @Failure 500 {object} api_response.ErrorResponse{errors=[]string}
// @Failure 400 {object} api_response.ErrorResponse{errors=[]string}
// @Router /loans/{loan_id}/delinquent [get]
func (h *LoanHandler) IsDelinquent(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 6*time.Second)
	defer cancel()

	loanID, err := strconv.ParseInt(c.Param("loan_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, api_response.BuildErrorResponse("invalid loan id",
			http.StatusBadRequest, err, nil))
	}
	mockDateStr := c.QueryParam("mock_date")

	var mockDate time.Time
	if mockDateStr == "" {
		mockDate = time.Now()
	} else {
		mockDate, err = time.Parse(
			"2006-01-02",
			mockDateStr,
		)

		if err != nil {
			return c.JSON(http.StatusBadRequest,
				api_response.BuildErrorResponse("invalid as_of_date format, use YYYY-MM-DD", http.StatusBadRequest,
					err, nil))
		}
	}

	data, err := h.LoanScheduleService.IsDelinquent(ctx, loanID, mockDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			api_response.BuildErrorResponse("Data Not Found", http.StatusBadRequest,
				err, nil))
	}

	return c.JSON(http.StatusOK,
		api_response.BuildSuccessResponse("Data Found", http.StatusOK, data))
}
