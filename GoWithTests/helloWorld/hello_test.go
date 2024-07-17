package main

import "testing"

func Test_Ola(t *testing.T) {
	checkCorrectMessage := func(t *testing.T, result, expected string) {
		t.Helper()

		if result != expected {
			t.Errorf("result '%s' , expected '%s'", result, expected)
		}
	}

	t.Run("diz olá Para as pessoas", func(t *testing.T) {
		result := Ola("mundo", "Portuguese")
		expected := "Olá, mundo"

		checkCorrectMessage(t, result, expected)
	})

	t.Run("'Mundo' como padrão para 'string' vazia", func(t *testing.T) {
		result := Ola("", "portuguese")
		expected := "Olá, Mundo"

		checkCorrectMessage(t, result, expected)
	})

	t.Run("em espanhol", func(t *testing.T) {
		result := Ola("Elodie", "Spanish")
		expected := "Holla, Elodie"

		checkCorrectMessage(t, result, expected)
	})

	t.Run("em frances", func(t *testing.T) {
		result := Ola("Elodie", "French")
		expected := "Bonjour, Elodie"

		checkCorrectMessage(t, result, expected)
	})
}
