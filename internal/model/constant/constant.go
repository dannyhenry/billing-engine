package constant

type LoanStatus string

const (
	LoanStatusActive LoanStatus = "ACTIVE"
	LoanStatusClosed LoanStatus = "CLOSED"
)

type ScheduleStatus string

const (
	ScheduleStatusPending ScheduleStatus = "PENDING"
	ScheduleStatusPaid    ScheduleStatus = "PAID"
)
