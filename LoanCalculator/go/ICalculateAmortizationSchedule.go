package swagger

type ICalculateAmortizationSchedule interface {
	CalculateRepayment(interestRate float64, loanTerm int32, loanAmount float64, totalNumberOfPayments int32) (repayment float64)
	CalculateTotalInterestPayable(loanAmount float64, repayment float64, totalNumberOfPayments int32) float64
	CalculateAmountOwning(interestRate float64, loanTerm int32, loanAmount float64, monthlyRepayment float64, totalNumberOfPayments int32) (loanRepaymentsAmountOwing []LoanRepaymentsAmountOwing)
}
