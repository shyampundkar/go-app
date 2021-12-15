package swagger

import (
	"math"
	"testing"
)

func TestCalculateRepayment(t *testing.T) {

	t.Run("calculate interest repayment", func(t *testing.T) {
		InterestCalculator := InterestCalculator{}
		var want float64 = 2109
		got := math.Round(InterestCalculator.CalculateRepayment(7.23, 1, 350000, 12))

		if got != want {
			t.Errorf("got: %f want: %f", got, want)
		}
	})

}
