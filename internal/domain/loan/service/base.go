package service

import (
	"github.com/dannyhenry/billing/internal/domain/loan"
)

type LoanService struct {
	repository loan.Repository
}

func NewService(repo loan.Repository) loan.Service {
	return &LoanService{
		repository: repo,
	}
}
