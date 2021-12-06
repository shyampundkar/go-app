package main

func SumVariadic(variableArgs ...[]int) (result []int) {

	//result = make([]int, len(variableArgs))

	for _, v := range variableArgs {
		result = append(result,SumAll(v))
	}

	return
}

func SumAll(numbersToSum []int) (sum int) {

	for _, v := range numbersToSum {
		sum += v
	}

	return
}
