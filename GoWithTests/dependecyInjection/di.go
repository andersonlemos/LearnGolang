package dependecyInjection

import (
	"fmt"
	"io"
)

func Greeting(writer io.Writer, msg string) {
	_, err := fmt.Fprintf(writer, "Hi, %s", msg)
	if err != nil {
		return
	}
}
