package payment

import "time"

type Payment struct {
	ID          int64     `gorm:"primaryKey;autoIncrement"`
	LoanID      int64     `gorm:"not null;index"`
	Amount      int64     `gorm:"not null"`
	PaymentDate time.Time `gorm:"not null"`
	CreatedAt   time.Time
}

func (Payment) TableName() string { return "payments" }

func (Payment) ModelName() string { return "Payment" }

type PaymentDetails struct {
	ID         int64 `gorm:"primaryKey;autoIncrement"`
	PaymentID  int64 `gorm:"not null;index"`
	ScheduleID int64 `gorm:"not null;index"`
	Amount     int64 `gorm:"not null"`
	CreatedAt  time.Time
}

func (PaymentDetails) TableName() string { return "payment_details" }
func (PaymentDetails) ModelName() string { return "PaymentDetails" }
