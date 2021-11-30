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

	t.Run("saying hellow to people", func(t *testing.T) {
		got := Hello("Shyam")
		want := "Hello, Shyam!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {

		got := Hello("")
		want := "Hello, World!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})
}
