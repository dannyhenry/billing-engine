package repository

import (
	"context"
	"errors"
	"time"

	"github.com/dannyhenry/billing/internal/model/constant"
	"github.com/dannyhenry/billing/internal/model/loan"
	"github.com/dannyhenry/billing/internal/model/loan_schedule"
	"github.com/dannyhenry/billing/internal/model/payment"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *LoanRepository) CreatePayment(ctx context.Context, loanID int64, amount int64) (loan.PaymentLoanResponse, error) {

	var result loan.PaymentLoanResponse

	err := r.Database.GetMaster().WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			var loanData loan.Loan

			if err := tx.
				Clauses(clause.Locking{Strength: "UPDATE"}).
				First(&loanData, "id = ?", loanID).
				Error; err != nil {
				return err
			}

			if loanData.Status == constant.LoanStatusClosed {
				return errors.New("loan already closed")
			}

			if amount%loanData.WeeklyAmount != 0 {
				return errors.New("payment amount must be multiple of weekly amount")
			}

			installmentCount := int(amount / loanData.WeeklyAmount)

			var schedules []loan_schedule.LoanSchedule

			if err := tx.
				Clauses(clause.Locking{Strength: "UPDATE"}).
				Debug().
				Where("loan_id = ? AND status = ? AND due_date<= ?", loanID, constant.ScheduleStatusPending, time.Now()).
				Order("week_number ASC").
				Limit(installmentCount).
				Find(&schedules).
				Error; err != nil {
				return err
			}

			if len(schedules) < installmentCount {
				return errors.New("no outstanding billing schedules")
			}

			paidWeeks := make([]int, 0, len(schedules))
			scheduleIDs := make([]int64, 0, len(schedules))

			for _, schedule := range schedules {
				paidWeeks = append(paidWeeks, schedule.WeekNumber)
				scheduleIDs = append(scheduleIDs, schedule.ID)
			}

			paidTime := time.Now()

			if err := tx.
				Model(&loan_schedule.LoanSchedule{}).
				Where("id IN ?", scheduleIDs).
				Updates(map[string]any{
					"status":  constant.ScheduleStatusPaid,
					"paid_at": paidTime,
				}).
				Error; err != nil {
				return err
			}

			pay := &payment.Payment{
				LoanID:      loanID,
				Amount:      amount,
				PaymentDate: paidTime,
			}

			if err := tx.Create(pay).Error; err != nil {
				return err
			}

			allocations := make([]payment.PaymentDetails, 0, len(schedules))

			for _, schedule := range schedules {
				allocations = append(allocations, payment.PaymentDetails{
					PaymentID:  pay.ID,
					ScheduleID: schedule.ID,
					Amount:     schedule.Amount,
				})
			}

			if err := tx.Create(&allocations).Error; err != nil {
				return err
			}

			var remainingOutstanding int64

			if err := tx.
				Model(&loan_schedule.LoanSchedule{}).
				Select("COALESCE(SUM(amount), 0)").
				Where("loan_id = ? AND status = ?", loanID, constant.ScheduleStatusPending).
				Scan(&remainingOutstanding).
				Error; err != nil {
				return err
			}

			if remainingOutstanding == 0 {
				loanData.Status = constant.LoanStatusClosed

				if err := tx.Save(&loan.Loan{}).Error; err != nil {
					return err
				}
			}

			result = loan.PaymentLoanResponse{
				PaymentID:            pay.ID,
				LoanID:               loanData.ID,
				Amount:               pay.Amount,
				PaidWeeks:            paidWeeks,
				RemainingOutstanding: remainingOutstanding,
				LoanStatus:           string(loanData.Status),
				PaymentDate:          pay.PaymentDate,
			}

			return nil

		})
	return result, err
}
