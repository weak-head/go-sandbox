package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("User", "")
		want := "hwgo, User!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello when empty name is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "hwgo, world!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Steve", "French")
		want := "Bonjour, Steve!"
		assertCorrectMessage(t, got, want)
	})
}
