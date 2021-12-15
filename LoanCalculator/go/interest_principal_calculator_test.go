package swagger

import (
	"math"
	"testing"
)

func TestCalculatePrincipalInterestRepayment(t *testing.T) {

	interestPrincipalCalculator := InterestPrincipalCalculator{}
	var want float64 = 30321
	got := math.Round(interestPrincipalCalculator.CalculateRepayment(7.23, 1, 350000, 12))

	if got != want {
		t.Errorf("got: %f want: %f", got, want)
	}
}
