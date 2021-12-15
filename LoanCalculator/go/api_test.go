package swagger

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestCalculateLoan(t *testing.T) {

	t.Run("Calculate interestOnly repayments - Monthly", func(t *testing.T) {
		var wantTotalInterestPayable float64 = 25305
		var wantMontlyRepayment float64 = 2109

		var calculateloanBody = CalculateloanBody{LoanAmount: 350000, LoanType: "InterestOnly", PaymentFrequency: "Monthly", InterestRate: 7.23, LoanTerm: 1}

		verify(calculateloanBody, wantMontlyRepayment, t, wantTotalInterestPayable)

	})

	t.Run("Calculate interestOnly repayments - Fortnightly", func(t *testing.T) {
		var wantTotalInterestPayable float64 = 25305
		var wantMontlyRepayment float64 = 973

		var calculateloanBody = CalculateloanBody{LoanAmount: 350000, LoanType: "InterestOnly", PaymentFrequency: "Fortnightly", InterestRate: 7.23, LoanTerm: 1}

		verify(calculateloanBody, wantMontlyRepayment, t, wantTotalInterestPayable)

	})

	t.Run("Calculate interestOnly repayments - Weekly", func(t *testing.T) {
		var wantTotalInterestPayable float64 = 25305
		var wantMontlyRepayment float64 = 487

		var calculateloanBody = CalculateloanBody{LoanAmount: 350000, LoanType: "InterestOnly", PaymentFrequency: "Weekly", InterestRate: 7.23, LoanTerm: 1}

		verify(calculateloanBody, wantMontlyRepayment, t, wantTotalInterestPayable)

	})

	t.Run("Calculate interest & principal repayments - Monthly", func(t *testing.T) {
		var wantTotalInterestPayable float64 = 13858
		var wantMontlyRepayment float64 = 30321

		var calculateloanBody = CalculateloanBody{LoanAmount: 350000, LoanType: "PrincipalAndInterest", PaymentFrequency: "Monthly", InterestRate: 7.23, LoanTerm: 1}

		verify(calculateloanBody, wantMontlyRepayment, t, wantTotalInterestPayable)

	})

	t.Run("Calculate interest & principal repayments - Fortnightly", func(t *testing.T) {
		var wantTotalInterestPayable float64 = 13291
		var wantMontlyRepayment float64 = 13973

		var calculateloanBody = CalculateloanBody{LoanAmount: 350000, LoanType: "PrincipalAndInterest", PaymentFrequency: "Fortnightly", InterestRate: 7.23, LoanTerm: 1}

		verify(calculateloanBody, wantMontlyRepayment, t, wantTotalInterestPayable)

	})

	t.Run("Calculate interest & principal repayments - Weekly", func(t *testing.T) {
		var wantTotalInterestPayable float64 = 13048
		var wantMontlyRepayment float64 = 6982

		var calculateloanBody = CalculateloanBody{LoanAmount: 350000, LoanType: "PrincipalAndInterest", PaymentFrequency: "Weekly", InterestRate: 7.23, LoanTerm: 1}

		verify(calculateloanBody, wantMontlyRepayment, t, wantTotalInterestPayable)

	})

	t.Run("Calculate interest & principal repayments - Weekly - Negative", func(t *testing.T) {
		var wantTotalInterestPayable float64 = 0
		var wantMontlyRepayment float64 = 0

		var calculateloanBody = CalculateloanBody{LoanAmount: 350000, LoanType: "PrincipalAndInterest1", PaymentFrequency: "Weekly", InterestRate: 7.23, LoanTerm: 1}

		verify(calculateloanBody, wantMontlyRepayment, t, wantTotalInterestPayable)

	})

	t.Run("Verify validation for empty body - Negative", func(t *testing.T) {

		req := httptest.NewRequest("POST", "/calculate-loan", nil)
		w := httptest.NewRecorder()
		CalculateLoan(w, req)
		requestBodyBytes, err := ioutil.ReadAll(w.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var modelError ModelError
		json.Unmarshal(requestBodyBytes, &modelError)

		if result, _ := strconv.Atoi(modelError.Code); (result != http.StatusBadRequest) || modelError.Message != "RequestBodyEmpty : Please enter request body" {
			t.Errorf("got: %s want: %d", modelError.Code, http.StatusBadRequest)
		}

	})

	t.Run("Verify validation for invalid body - Negative", func(t *testing.T) {

		postBody := []byte("test")

		responseBody := bytes.NewBuffer(postBody)

		req := httptest.NewRequest("POST", "/calculate-loan", responseBody)
		w := httptest.NewRecorder()
		CalculateLoan(w, req)
		requestBodyBytes, err := ioutil.ReadAll(w.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var modelError ModelError
		json.Unmarshal(requestBodyBytes, &modelError)

		if result, _ := strconv.Atoi(modelError.Code); result != http.StatusBadRequest {
			t.Errorf("got: %s want: %d", modelError.Code, http.StatusBadRequest)
		}

	})

	t.Run("Verify validation for Method Not Allowed - Negative", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/calculate-loan", nil)
		w := httptest.NewRecorder()
		CalculateLoan(w, req)
		requestBodyBytes, err := ioutil.ReadAll(w.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var modelError ModelError
		json.Unmarshal(requestBodyBytes, &modelError)

		if result, _ := strconv.Atoi(modelError.Code); result != http.StatusBadRequest {
			t.Errorf("got: %s want: %d", modelError.Code, http.StatusBadRequest)
		}

	})

}

func verify(calculateloanBody CalculateloanBody, wantMonthlyRepayment float64, t *testing.T, wantTotalInterestPayable float64) {
	t.Helper()

	loanRepayments := CalculateLoanRepayments(calculateloanBody)

	if loanRepayments.MonthlyRepayments != wantMonthlyRepayment {
		t.Errorf("Monthly repayment got: %f want: %f", loanRepayments.MonthlyRepayments, wantMonthlyRepayment)
	}

	if loanRepayments.TotalInterestPayable != wantTotalInterestPayable {
		t.Errorf("Monthly repayment got: %f want: %f", loanRepayments.TotalInterestPayable, wantTotalInterestPayable)
	}
}
