package repository

import (
	"github.com/dannyhenry/billing/common/database"
	loan_schedules "github.com/dannyhenry/billing/internal/domain/loan-schedules"
)

type LoanScheduleRepository struct {
	database database.DatabaseManager
}

func NewRepository(db database.DatabaseManager) loan_schedules.Repository {
	return &LoanScheduleRepository{database: db}
}
