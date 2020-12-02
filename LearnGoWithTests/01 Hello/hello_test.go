package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Dude", "")
		want := "Hello Dude!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when no name is given", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("en Espanol", func(t *testing.T) {
		got := Hello("Marisol", "Spanish")
		want := "Hola Marisol!"
		assertCorrectMessage(t, got, want)
	})
}
