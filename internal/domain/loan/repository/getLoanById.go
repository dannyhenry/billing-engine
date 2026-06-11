package repository

import (
	"context"
	"errors"

	entity "github.com/dannyhenry/billing/internal/model/loan"
	"gorm.io/gorm"
)

func (r *LoanRepository) GetLoanByID(ctx context.Context, loanID int64) (*entity.Loan, error) {
	var loan entity.Loan

	err := r.Database.GetMaster().WithContext(ctx).
		First(&loan, "id = ?", loanID).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("loan not found")
	}

	return &loan, err
}
