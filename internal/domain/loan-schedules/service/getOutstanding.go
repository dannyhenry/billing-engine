package service

import "context"

func (s *LoanScheduleService) GetOutstanding(ctx context.Context, loanId int64) (int64, error) {
	return s.repository.GetOutstandingLoanById(ctx, loanId)
}
