package service

import (
	"context"
	"errors"

	"github.com/dannyhenry/billing/internal/model/loan"
)

func (s *LoanService) PayLoans(ctx context.Context, request loan.PaymentLoanRequest) (loan.PaymentLoanResponse, error) {
	if request.Amount <= 0 {
		return loan.PaymentLoanResponse{}, errors.New("payment amount must be greater than zero")
	}

	return s.repository.CreatePayment(ctx, request.LoanID, request.Amount)
}
