/*
 * Loan Calculator
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
)

type ICalculateAmortizationSchedule interface{}

func CalculateLoan(w http.ResponseWriter, r *http.Request) {

	requestBodyBytes, _ := ioutil.ReadAll(r.Body)
	var calculateloanBody CalculateloanBody
	json.Unmarshal(requestBodyBytes, &calculateloanBody)

	loanRepayment := CalculateLoanRepaymentsAmountOwing(calculateloanBody)

	j, _ := json.Marshal(&loanRepayment)
	w.Write(j)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.WriteHeader(http.StatusOK)
}

func CalculateLoanRepaymentsAmountOwing(calculateloanBody CalculateloanBody) (loanRepayment LoanRepayments) {

	switch calculateloanBody.LoanType {
	case "PrincipalAndInterest":
	case "Interest":
	default:
	}

	var totalNumberOfPayments int32

	switch calculateloanBody.PaymentFrequency {
	case "Monthly":
		totalNumberOfPayments = 12 * calculateloanBody.LoanTerm
	case "Fortnightly":
		totalNumberOfPayments = 26 * calculateloanBody.LoanTerm
	case "Weekly":
		totalNumberOfPayments = 52 * calculateloanBody.LoanTerm

	}
	monthlyRepayment := CalculateMonthlyRepayment(calculateloanBody.InterestRate, calculateloanBody.LoanTerm, calculateloanBody.LoanAmount, totalNumberOfPayments)
	loanRepayment.MonthlyRepayments = math.Round(monthlyRepayment)
	loanRepayment.TotalInterestPayable = math.Round(CalculateTotalInterestPayable(calculateloanBody.LoanAmount, monthlyRepayment, totalNumberOfPayments))
	loanRepayment.AmountOwning = CalculateAmountOwning(calculateloanBody, monthlyRepayment, totalNumberOfPayments)

	return
}

func CalculateMonthlyRepayment(InterestRate float64, LoanTerm int32, LoanAmount float64, totalNumberOfPayments int32) (monthlyRepayments float64) {

	if InterestRate != 0 {
		rate := (InterestRate / float64(totalNumberOfPayments) / 100)
		loanTermInMonths := LoanTerm * totalNumberOfPayments
		monthlyRepayments = float64(LoanAmount) * (rate + (rate / (math.Pow(float64(rate+1), float64(loanTermInMonths)) - 1)))
	}
	return
}

func CalculateTotalInterestPayable(loanAmout, repayment float64, totalNumberOfPayments int32) float64 {
	return (repayment * float64(totalNumberOfPayments)) - loanAmout
}

func CalculateAmountOwning(calculateloanBody CalculateloanBody, monthlyRepayment float64, totalNumberOfPayments int32) (loanRepaymentsAmountOwing []LoanRepaymentsAmountOwing) {

	interestrate := calculateloanBody.InterestRate

	loanTermInMonths := calculateloanBody.LoanTerm * totalNumberOfPayments
	var monthlyInterest = ((interestrate / 100) / float64(loanTermInMonths))
	loanAmount := calculateloanBody.LoanAmount

	var initialPeriod LoanRepaymentsAmountOwing = LoanRepaymentsAmountOwing{}
	initialPeriod.Year = 0
	initialPeriod.Principal = loanAmount
	currentInterest := (monthlyRepayment*float64(loanTermInMonths) - float64(loanAmount))
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
		reducingPrincipal := (loanAmount) - (monthlyRepayment - currentInterest)
		period.Principal = math.Round(reducingPrincipal)

		period.Total = math.Round(reducingPrincipal + reducingInterest)
		loanAmount = reducingPrincipal

		loanRepaymentsAmountOwing = append(loanRepaymentsAmountOwing, period)

	}
	return
}
