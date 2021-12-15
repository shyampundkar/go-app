package swagger

import "errors"

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
