package main

func Sum(nums [5]int) (result int) {
	
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return result
}
