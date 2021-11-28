package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println("Hello welcome to the Calculator")
	var operation string
	var operand1 int
	var operand2 int

	fmt.Println("Operation:")
	fmt.Scan(&operation)
	fmt.Println("Operand 1:")
	fmt.Scan(&operand1)
	fmt.Println("Operand 2:")
	fmt.Scan(&operand2)
	result, err := Calculate(operation, operand1, operand2)
	if err != nil {
		fmt.Println("Error:" + err.Error())
	} else {
		fmt.Printf("\nThe result is:%d", result)

	}
}

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
