package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("syamim", "English")

		want := "hello, syamim"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")

		want := "hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")

		want := "hola, Elodie"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("go %q want %q", got, want)
	}
}
