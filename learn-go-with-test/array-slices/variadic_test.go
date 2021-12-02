package main

import (
	"reflect"
	"testing"
)

func TestVariadic(t *testing.T) {

	got := SumVariadic([]int{1, 2}, []int{1, 2, 3, 4})
	want := []int{3, 10}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v, got %v", want, got)
	}
}

// func TestSumAll(t *testing.T) {

//     got := SumAll([]int{1, 2}, []int{0, 9})
//     want := []int{3, 9}

//     if !reflect.DeepEqual(got, want) {
//         t.Errorf("got %v want %v", got, want)
//     }
// }
