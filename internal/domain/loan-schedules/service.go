package loan_schedules

import (
	"context"
	"time"
)

type Service interface {
	GetOutstanding(ctx context.Context, loanId int64) (int64, error)
	IsDelinquent(ctx context.Context, loanID int64, date time.Time) (bool, error)
}
