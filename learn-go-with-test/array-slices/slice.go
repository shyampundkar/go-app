package main

func SumSlices(slices []int) (result int) {

	for _, val := range slices {
		result += val
	}
	return result
}
