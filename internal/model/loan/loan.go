package loan

import (
	"time"

	"github.com/dannyhenry/billing/internal/model/constant"
)

type Loan struct {
	ID              int64               `gorm:"primaryKey;autoIncrement" json:"id"`
	CustomerID      int64               `gorm:"not null" json:"customer_id"`
	PrincipalAmount int64               `gorm:"not null" json:"principal_amount"`
	InterestRate    float64             `gorm:"not null" json:"interest_rate"`
	TotalAmount     int64               `gorm:"not null" json:"total_amount"`
	TenureWeeks     int                 `gorm:"not null" json:"tenure_weeks"`
	WeeklyAmount    int64               `gorm:"not null" json:"weekly_amount"`
	StartDate       time.Time           `gorm:"not null" json:"start_date"`
	Status          constant.LoanStatus `gorm:"type:varchar(20);not null" json:"status"`

	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
	IsDeleted bool      `json:"is_deleted"`
}
