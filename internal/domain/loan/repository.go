package loan

import (
	"context"

	"github.com/dannyhenry/billing/internal/model/loan"
)

type Repository interface {
	GetLoanByID(ctx context.Context, loanID int64) (*loan.Loan, error)
	CreatePayment(ctx context.Context, loanID int64, amount int64) (loan.PaymentLoanResponse, error)
}
