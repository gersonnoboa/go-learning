package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Gerson", "")
		want := "Hello, Gerson"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello World' when empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run(("in Spanish"), func(t *testing.T) {
		got := Hello("Relika", "Spanish")
		want := "Hola, Relika"
		assertCorrectMessage(t, got, want)
	})
	t.Run(("in French"), func(t *testing.T) {
		got := Hello("Relika", "French")
		want := "Bonjour, Relika"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
