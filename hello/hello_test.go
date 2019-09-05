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
		got := Hello("Chirs", "")
		want := "Hello, Chirs"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello world when a empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("French", "French")
		want := "Bonjour, French"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Chinese", func(t *testing.T) {
		got := Hello("世界", "Chinese")
		want := "你好, 世界"

		assertCorrectMessage(t, got, want)
	})
}
