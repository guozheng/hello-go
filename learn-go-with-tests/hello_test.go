package main

import "testing"

func TestHello(t *testing.T) {
	testOutput := func(t *testing.T, got, want string) {
		t.Helper() //correctly report test failure line number
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		testOutput(t, Hello("Alex"), "Hello, Alex")
	})

	t.Run("say hello to nobody", func(t *testing.T) {
		testOutput(t, Hello(""), "Hello, nobody")
	})
}
