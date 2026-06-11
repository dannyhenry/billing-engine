package service

import (
	"context"
	"time"
)

func (s *LoanScheduleService) IsDelinquent(ctx context.Context, loanID int64, date time.Time) (bool, error) {
	numOfMissing, err := s.repository.GetCountDueSchedules(ctx, loanID, date)
	if err != nil {
		return false, err
	}

	if numOfMissing >= 2 {
		return true, nil
	}
	return false, nil
}
