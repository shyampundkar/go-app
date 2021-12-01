package main

import "testing"

func TestIterations (t *testing.T){

	got:= RepeatBy("a", 5)
	want:="aaaaa"

	if got!= want{
		t.Errorf("got %q want %q", got, want);
	}
}


func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        RepeatBy("a",5)
    }
}
