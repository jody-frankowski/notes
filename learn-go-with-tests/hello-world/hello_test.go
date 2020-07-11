package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Jody", "")
		want := "Hello, Jody"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hola, Jody' in Spanish", func(t *testing.T) {
		got := Hello("Jody", "Spanish")
		want := "Hola, Jody"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Bonjour, Jody' in French", func(t *testing.T) {
		got := Hello("Jody", "French")
		want := "Bonjour, Jody"
		assertCorrectMessage(t, got, want)
	})

}
