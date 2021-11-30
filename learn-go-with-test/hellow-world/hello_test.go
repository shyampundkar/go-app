package main

import "testing"

// func TestHello(t *testing.T) {
// 	var got string = Hello()
// 	var want string = "Hello, world"

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

func TestHellow(t *testing.T) {
	got := Hello("Shyam")
	want := "Hello, Shyam!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
