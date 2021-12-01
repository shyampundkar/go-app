package main

func Sum(nums [5]int) (result int) {

	for _, num := range nums {
		result += num
	}
	return result
}
