package loan

import (
	"github.com/dannyhenry/billing/internal/domain/loan"
	"github.com/dannyhenry/billing/internal/domain/loan-schedules"
)

type LoanHandler struct {
	LoanService         loan.Service
	LoanScheduleService loan_schedules.Service
}

func NewHandler(ls loan.Service, lss loan_schedules.Service) *LoanHandler {
	return &LoanHandler{LoanService: ls, LoanScheduleService: lss}
}
