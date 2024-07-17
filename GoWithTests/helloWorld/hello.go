package main

import (
	"fmt"
	"strings"
)

const portuguese = "portuguese"
const spanish = "spanish"
const french = "french"

const portuguesePrefix = "Olá, "
const spanishPrefix = "Holla, "
const frenchPrefix = "Bonjour, "

func greettingsPrefix(language string) (prefix string) {
	switch strings.ToLower(language) {
	case portuguese:
		prefix = portuguesePrefix
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = portuguesePrefix
	}
	return
}

func Ola(name string, language string) string {
	if name == "" {
		name = "Mundo"
	}
	return greettingsPrefix(language) + name
}
func main() {
	fmt.Println(Ola("Cris", ""))
}
