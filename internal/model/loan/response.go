package loan

import "time"

type PaymentLoanResponse struct {
	PaymentID            int64     `json:"payment_id"`
	LoanID               int64     `json:"loan_id"`
	Amount               int64     `json:"amount"`
	PaidWeeks            []int     `json:"paid_weeks"`
	RemainingOutstanding int64     `json:"remaining_outstanding"`
	LoanStatus           string    `json:"loan_status"`
	PaymentDate          time.Time `json:"payment_date"`
}
