package main

func SumVariadic(variableArgs ...[]int) (result []int) {

	result = make([]int, len(variableArgs))

	for i, v := range variableArgs {
		result[i] = SumAll(v)

	}

	return
}

func SumAll(numbersToSum []int) (sum int) {

	for _, v := range numbersToSum {
		sum += v
	}

	return
}
