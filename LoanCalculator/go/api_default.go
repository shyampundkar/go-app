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
	"errors"
	"io/ioutil"
	"math"
	"net/http"
)

func CalculateLoan(w http.ResponseWriter, r *http.Request) {

	requestBodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var calculateloanBody CalculateloanBody
	err = json.Unmarshal(requestBodyBytes, &calculateloanBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	loanRepayment := CalculateLoanRepaymentsAmountOwing(calculateloanBody)

	j, err := json.Marshal(&loanRepayment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(j)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.WriteHeader(http.StatusOK)
}

func CalculateLoanRepaymentsAmountOwing(calculateloanBody CalculateloanBody) (loanRepayment LoanRepayments) {

	var totalNumberOfPayments int32

	calculator, err := CreateCalculator(calculateloanBody.LoanType)

	if err != nil {
		return LoanRepayments{}
	}

	switch calculateloanBody.PaymentFrequency {
	case "Monthly":
		totalNumberOfPayments = 12 * calculateloanBody.LoanTerm
	case "Fortnightly":
		totalNumberOfPayments = 26 * calculateloanBody.LoanTerm
	case "Weekly":
		totalNumberOfPayments = 52 * calculateloanBody.LoanTerm

	}
	repayment := calculator.CalculateRepayment(calculateloanBody.InterestRate, calculateloanBody.LoanTerm, calculateloanBody.LoanAmount, totalNumberOfPayments)
	loanRepayment.MonthlyRepayments = math.Round(repayment)
	loanRepayment.TotalInterestPayable = math.Round(calculator.CalculateTotalInterestPayable(calculateloanBody.LoanAmount, repayment, totalNumberOfPayments))
	loanRepayment.AmountOwning = calculator.CalculateAmountOwning(calculateloanBody.InterestRate, calculateloanBody.LoanTerm, calculateloanBody.LoanAmount, repayment, totalNumberOfPayments)

	return
}

func CreateCalculator(calculatorType string) (ICalculateAmortizationSchedule, error) {

	switch calculatorType {
	case "PrincipalAndInterest":
		return InterestPrincipalCalculator{}, nil
	case "InterestOnly":
		return InterestCalculator{}, nil
	default:
		return nil, errors.New("Calculator is " + calculatorType + " not supported yet")

	}

}
