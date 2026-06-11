package repository

import (
	"github.com/dannyhenry/billing/common/database"
	"github.com/dannyhenry/billing/internal/domain/loan"
)

type LoanRepository struct {
	Database database.DatabaseManager
}

func NewRepository(db database.DatabaseManager) loan.Repository {
	return &LoanRepository{Database: db}
}
