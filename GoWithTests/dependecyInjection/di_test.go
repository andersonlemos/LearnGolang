package dependecyInjection

import (
	"bytes"
	"testing"
)

func TestGreeting(t *testing.T) {
	buffer := bytes.Buffer{}
	Greeting(&buffer, "Greeting")

	result := buffer.String()
	expected := "Hi, Greeting"
	if result != expected {
		t.Errorf("Greeting() = %q; want %q", result, expected)
	}
}
