package repository

import (
	"context"

	"github.com/dannyhenry/billing/internal/model/constant"
	entity "github.com/dannyhenry/billing/internal/model/loan_schedule"
)

func (r *LoanScheduleRepository) GetOutstandingLoanById(ctx context.Context, loanID int64) (int64, error) {
	var outstanding int64

	err := r.database.GetMaster().
		WithContext(ctx).
		Model(&entity.LoanSchedule{}).
		Select("COALESCE(SUM(amount), 0)").
		Where("loan_id = ? AND status = ?", loanID, constant.ScheduleStatusPending).
		Scan(&outstanding).
		Error

	return outstanding, err
}
