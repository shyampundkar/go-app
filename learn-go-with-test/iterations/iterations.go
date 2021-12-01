package main

func RepeatBy(init string, counter int) (result string) {

	for i := 0; i < counter; i++ {
		result += init
	}
	return result
}
