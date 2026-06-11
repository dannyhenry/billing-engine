package repository

import (
	"context"
	"time"

	"github.com/dannyhenry/billing/internal/model/constant"
	entity "github.com/dannyhenry/billing/internal/model/loan_schedule"
)

func (r *LoanScheduleRepository) GetCountDueSchedules(ctx context.Context, loanID int64, date time.Time) (int64, error) {
	var numOfMissing int64

	err := r.database.GetMaster().WithContext(ctx).
		Model(&entity.LoanSchedule{}).
		Where("loan_id = ? AND due_date <= ? AND status = ?", loanID, date, constant.ScheduleStatusPending).
		Count(&numOfMissing).
		Error

	return numOfMissing, err
}
