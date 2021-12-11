package swagger

import "math"

type InterestCalculator struct {
	CalculatorBase
	ICalculateAmortizationSchedule
}

func (InterestCalculator) CalculateRepayment(InterestRate float64, LoanTerm int32, LoanAmount float64, totalNumberOfPayments int32) (repayment float64) {

	if InterestRate != 0 {
		rate := (InterestRate / float64(totalNumberOfPayments) / 100)
		repayment = float64(LoanAmount) * rate
	}
	return
}

func (InterestCalculator) CalculateTotalInterestPayable(loanAmout, repayment float64, totalNumberOfPayments int32) float64 {

	return repayment * float64(totalNumberOfPayments)
}

func (InterestCalculator) CalculateAmountOwning(interestRate float64, loanTerm int32, loanAmount float64, monthlyRepayment float64, totalNumberOfPayments int32) (loanRepaymentsAmountOwing []LoanRepaymentsAmountOwing) {
	var initialPeriod LoanRepaymentsAmountOwing = LoanRepaymentsAmountOwing{}
	initialPeriod.Year = 0
	initialPeriod.Principal = loanAmount
	currentInterest := (monthlyRepayment * float64(totalNumberOfPayments))
	initialPeriod.Interest = (math.Ceil(currentInterest))
	initialPeriod.Total = loanAmount + float64(initialPeriod.Interest)
	loanRepaymentsAmountOwing = append(loanRepaymentsAmountOwing, initialPeriod)
	previousInterest := currentInterest
	rate := (interestRate / float64(totalNumberOfPayments) / 100)
	currentInterest = loanAmount * rate

	for i := int32(1); i <= totalNumberOfPayments; i++ {

		period := LoanRepaymentsAmountOwing{}
		period.Year = i

		reducingInterest := previousInterest - currentInterest
		if reducingInterest > 0 {
			period.Interest = math.Round(reducingInterest)
		}

		previousInterest = reducingInterest
		period.Principal = math.Round(loanAmount)
		period.Total = math.Round(loanAmount + reducingInterest)

		loanRepaymentsAmountOwing = append(loanRepaymentsAmountOwing, period)

	}
	return
}
