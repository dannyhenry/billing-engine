package loan_schedules

import (
	"context"
	"time"
)

type Repository interface {
	GetOutstandingLoanById(ctx context.Context, loanID int64) (int64, error)
	GetCountDueSchedules(ctx context.Context, loanID int64, date time.Time) (int64, error)
}
