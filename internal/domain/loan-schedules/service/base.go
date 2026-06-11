package service

import (
	loan_schedules "github.com/dannyhenry/billing/internal/domain/loan-schedules"
)

type LoanScheduleService struct {
	repository loan_schedules.Repository
}

func NewService(repository loan_schedules.Repository) loan_schedules.Service {
	return &LoanScheduleService{repository: repository}
}
