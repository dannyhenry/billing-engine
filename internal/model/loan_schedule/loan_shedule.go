package loan_schedule

import (
	"time"

	"github.com/dannyhenry/billing/internal/model/constant"
)

type LoanSchedule struct {
	ID         int64                   `gorm:"primaryKey;autoIncrement" json:"id"`
	LoanID     int64                   `gorm:"not null;index" json:"loan_id"`
	WeekNumber int                     `gorm:"not null" json:"week_number"`
	DueDate    time.Time               `gorm:"not null" json:"due_date"`
	Amount     int64                   `gorm:"not null" json:"amount"`
	Status     constant.ScheduleStatus `gorm:"type:varchar(20);not null" json:"status"`
	PaidAt     *time.Time              `json:"paid_at"`

	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
	IsDeleted bool      `json:"is_deleted"`
}

func (LoanSchedule) ModelName() string {
	return "LoanSchedule"
}

func (LoanSchedule) TableName() string {
	return "loan_schedules"
}
