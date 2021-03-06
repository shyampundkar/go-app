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
	"log"
	"math"
	"net/http"
	"strconv"
)

func CalculateLoan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	requestBodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(&w, http.StatusBadRequest, err)
		return
	}

	if len(requestBodyBytes) == 0 {
		handleError(&w, http.StatusBadRequest, errors.New("RequestBodyEmpty : Please enter request body"))
		return
	}

	var calculateloanBody CalculateloanBody
	err = json.Unmarshal(requestBodyBytes, &calculateloanBody)

	if err != nil {
		handleError(&w, http.StatusBadRequest, err)
		return
	}

	loanRepayment := CalculateLoanRepayments(calculateloanBody)

	j, err := json.Marshal(&loanRepayment)
	if err != nil {
		handleError(&w, http.StatusInternalServerError, err)
		return
	}
	_, err = w.Write(j)
	if err != nil {
		handleError(&w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleError(w *http.ResponseWriter, errorCode int, er error) ModelError {
	var modelError = ModelError{Code: strconv.Itoa(errorCode), Message: er.Error()}
	modelErrorBytes, err := json.Marshal(&modelError)
	if err == nil {
		_, e := (*w).Write(modelErrorBytes)

		if e != nil {
			log.Fatal(e.Error())
		}

		(*w).WriteHeader(errorCode)

	}
	return modelError
}

func CalculateLoanRepayments(calculateloanBody CalculateloanBody) (loanRepayment LoanRepayments) {

	calculator, err := CreateCalculator(calculateloanBody.LoanType)

	if err != nil {
		return LoanRepayments{}
	}

	totalNumberOfPayments := CalculateTotalNumberOfPayments(calculateloanBody)
	repayment := calculator.CalculateRepayment(calculateloanBody.InterestRate, calculateloanBody.LoanTerm, calculateloanBody.LoanAmount, totalNumberOfPayments)
	loanRepayment.MonthlyRepayments = math.Round(repayment)
	loanRepayment.TotalInterestPayable = math.Round(calculator.CalculateTotalInterestPayable(calculateloanBody.LoanAmount, repayment, totalNumberOfPayments))
	loanRepayment.AmountOwning = calculator.CalculateAmountOwning(calculateloanBody.InterestRate, calculateloanBody.LoanTerm, calculateloanBody.LoanAmount, repayment, totalNumberOfPayments)

	return
}

func CalculateTotalNumberOfPayments(calculateloanBody CalculateloanBody) (totalNumberOfPayments int32) {
	switch calculateloanBody.PaymentFrequency {
	case "Monthly":
		totalNumberOfPayments = 12 * calculateloanBody.LoanTerm
	case "Fortnightly":
		totalNumberOfPayments = 26 * calculateloanBody.LoanTerm
	case "Weekly":
		totalNumberOfPayments = 52 * calculateloanBody.LoanTerm

	}
	return totalNumberOfPayments
}
