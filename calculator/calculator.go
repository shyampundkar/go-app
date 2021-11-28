package calculator

import (
	"errors"
)

func Calculate(operation string, arg1, arg2 int) (int, error) {
	switch operation {
	case "+":
		{
			return arg1 + arg2, nil
		}

	case "-":
		{
			return arg1 - arg2, nil
		}

	case "*":
		{
			return arg1 * arg2, nil
		}

	case "/":
		{
			return arg1 / arg2, nil
		}
	default:
		return 0, errors.New("invalid operator")
	}
}
