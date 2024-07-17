package mocks

import (
	"fmt"
	"io"
	"time"
)

const lastWord = "Go!"
const startCount = 3

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	pause    func(duration time.Duration)
}

func (cs *ConfigurableSleeper) Sleep() {
	cs.pause(cs.duration)
}

func Count(writer io.Writer, sleeper Sleeper) {
	for i := startCount; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(writer, i)
	}
	sleeper.Sleep()
	fmt.Fprint(writer, lastWord)
}
