package swagger

import "math"

type InterestPrincipalCalculator struct {
	ICalculateAmortizationSchedule
}

// Calculate how much needs to be paid back for each month/fortnightly/weekly
func (InterestPrincipalCalculator) CalculateRepayment(InterestRate float64, LoanTerm int32, LoanAmount float64, totalNumberOfPayments int32) (repayment float64) {

	if InterestRate != 0 {
		rate := (InterestRate / float64(totalNumberOfPayments) / 100)
		loanTermInMonths := LoanTerm * totalNumberOfPayments
		repayment = float64(LoanAmount) * (rate + (rate / (math.Pow(float64(rate+1), float64(loanTermInMonths)) - 1)))
	}
	return
}

// Calculate total interest payable for the whole loan term (e.g 1,2,3year...etc)
func (InterestPrincipalCalculator) CalculateTotalInterestPayable(loanAmout, repayment float64, totalNumberOfPayments int32) float64 {
	return (repayment * float64(totalNumberOfPayments)) - loanAmout
}

// Calculate the reducing interest and principal for the loan term
func (InterestPrincipalCalculator) CalculateAmountOwning(interestRate float64, loanTerm int32, loanAmount float64, repayment float64, totalNumberOfPayments int32) (loanRepaymentsAmountOwing []LoanRepaymentsAmountOwing) {

	interestrate := interestRate

	loanTermInMonths := loanTerm * totalNumberOfPayments
	var monthlyInterest = ((interestrate / 100) / float64(loanTermInMonths))

	var initialPeriod LoanRepaymentsAmountOwing = LoanRepaymentsAmountOwing{}
	initialPeriod.Year = 0
	initialPeriod.Principal = loanAmount
	currentInterest := (repayment*float64(loanTermInMonths) - float64(loanAmount))
	initialPeriod.Interest = (math.Ceil(currentInterest))
	initialPeriod.Total = loanAmount + float64(initialPeriod.Interest)
	loanRepaymentsAmountOwing = append(loanRepaymentsAmountOwing, initialPeriod)
	previousInterest := currentInterest

	for i := int32(1); i <= loanTermInMonths; i++ {

		currentInterest = loanAmount * monthlyInterest

		period := LoanRepaymentsAmountOwing{}
		period.Year = i

		reducingInterest := previousInterest - currentInterest
		period.Interest = math.Round(reducingInterest)
		previousInterest = reducingInterest
		reducingPrincipal := loanAmount - (repayment - currentInterest)
		period.Principal = math.Round(reducingPrincipal)

		period.Total = math.Round(reducingPrincipal + reducingInterest)
		loanAmount = reducingPrincipal

		loanRepaymentsAmountOwing = append(loanRepaymentsAmountOwing, period)

	}
	return
}
