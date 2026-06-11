package loan

import (
	"context"

	"github.com/dannyhenry/billing/internal/model/loan"
)

type Service interface {
	PayLoans(ctx context.Context, request loan.PaymentLoanRequest) (loan.PaymentLoanResponse, error)
}
