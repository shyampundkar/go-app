package main

import "testing"

func TestArray (t *testing.T){
	//var nums [5]int= [5]int{1,2,3,4,5}
	nums := [...]int{1, 2, 3, 4, 5}
	got :=Sum(nums)
	want :=15

	if got!=want{
		t.Errorf("got %d want %d, %v", got, want, nums)
	}
}
