package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to name supplied", func(t *testing.T) {
		got := HelloWorld("Nisarg", "English")
		want := "Hello, Nisarg"
		assertTest(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := HelloWorld("", "")
		want := "Hello, Unknown"
		assertTest(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := HelloWorld("Nisarg", "Spanish")
		want := "Hola, Nisarg"
		assertTest(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := HelloWorld("Nisarg", "french")
		want := "Bonjour, Nisarg"
		assertTest(t, got, want)
	})

	t.Run("in french with no name", func(t *testing.T) {
		got := HelloWorld("", "French")
		want := "Bonjour, Unknown"
		assertTest(t, got, want)
	})
}

func assertTest(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
