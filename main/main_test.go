package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to name supplied", func(t *testing.T) {
		got := helloWorld("Nisarg")
		want := "Hello, Nisarg"
		assertTest(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := helloWorld("")
		want := "Hello, World"
		assertTest(t, got, want)
	})
}

func assertTest(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
