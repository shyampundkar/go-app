package swagger

import "math"

type InterestCalculator struct {
	ICalculateAmortizationSchedule
}

// Calculate loan repayment based on frequency
func (InterestCalculator) CalculateRepayment(InterestRate float64, LoanTerm int32, LoanAmount float64, totalNumberOfPayments int32) (repayment float64) {

	if InterestRate != 0 {
		rate := (InterestRate / float64(totalNumberOfPayments) / 100)
		repayment = float64(LoanAmount) * rate
	}
	return
}

// Calculate total interest payable for the whole loan term (e.g 1,2,3year...etc)
func (InterestCalculator) CalculateTotalInterestPayable(loanAmout, repayment float64, totalNumberOfPayments int32) float64 {

	return repayment * float64(totalNumberOfPayments)
}

// Calculate the reducing interest and principal for the loan term
func (InterestCalculator) CalculateAmountOwning(interestRate float64, loanTerm int32, loanAmount float64, monthlyRepayment float64, totalNumberOfPayments int32) (loanRepaymentsAmountOwing []LoanRepaymentsAmountOwing) {

	currentInterest := (monthlyRepayment * float64(totalNumberOfPayments))
	previousInterest := currentInterest
	initialPeriod := createInitialPeriod(loanAmount, currentInterest, monthlyRepayment, totalNumberOfPayments)
	loanRepaymentsAmountOwing = append(loanRepaymentsAmountOwing, initialPeriod)

	rate := calculateInterestRate(interestRate, totalNumberOfPayments)

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

func calculateInterestRate(interestRate float64, totalNumberOfPayments int32) float64 {
	return (interestRate / float64(totalNumberOfPayments) / 100)
}

func createInitialPeriod(loanAmount, currentInterest, monthlyRepayment float64, totalNumberOfPayments int32) LoanRepaymentsAmountOwing {
	var initialPeriod LoanRepaymentsAmountOwing = LoanRepaymentsAmountOwing{}
	initialPeriod.Year = 0
	initialPeriod.Principal = loanAmount
	initialPeriod.Interest = (math.Ceil(currentInterest))
	initialPeriod.Total = loanAmount + float64(initialPeriod.Interest)
	return initialPeriod
}
