package loan

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/dannyhenry/billing/internal/model/api_response"
	"github.com/labstack/echo/v4"
)

func (h *LoanHandler) GetOutstandingLoan(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 6*time.Second)
	defer cancel()

	loanID, err := strconv.ParseInt(c.Param("loan_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, api_response.ApplicationError("invalid loan id",
			http.StatusBadRequest, err, nil))
	}

	data, err := h.LoanScheduleService.GetOutstanding(ctx, loanID)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			api_response.ApplicationError("Data Not Found", http.StatusBadRequest,
				err, nil))
	}

	return c.JSON(http.StatusOK,
		api_response.BuildSuccessResponse("Data Found", http.StatusOK, data))
}
