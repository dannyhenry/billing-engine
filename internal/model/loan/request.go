package loan

type CreateLoanRequest struct {
	CustomerID      int64   `json:"customer_id"`
	PrincipalAmount int64   `json:"principal_amount"`
	InterestRate    float64 `json:"interest_rate"`
	TenureWeeks     int     `json:"tenure_weeks"`
}

type PaymentLoanRequest struct {
	LoanID int64 `json:"loan_id"`
	Amount int64 `json:"amount"`
}
