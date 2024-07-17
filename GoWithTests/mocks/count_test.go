package mocks

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SleeperSpy struct {
	Calls int
}

func (s *SleeperSpy) Sleep() {
	s.Calls++
}

const pause = "stop"
const write = "write"

type CountOperationSpy struct {
	Calls []string
}

type TimeSpy struct {
	pauseDuration time.Duration
}

func (s *TimeSpy) Pause(duration time.Duration) {
	s.pauseDuration = duration
}
func (s *CountOperationSpy) Sleep() {
	s.Calls = append(s.Calls, pause)
}

func (s *CountOperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCount(t *testing.T) {
	t.Run("print 3,2,1 and Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Count(buffer, &CountOperationSpy{})

		result := buffer.String()
		expected := `3
2
1
Go!`

		if result != expected {
			t.Errorf("got %s; want %s", result, expected)
		}
	})

	t.Run("pause after each print", func(t *testing.T) {
		spyPrintSleeper := &CountOperationSpy{}
		Count(spyPrintSleeper, spyPrintSleeper)

		expected := []string{
			pause,
			write,
			pause,
			write,
			pause,
			write,
			pause,
			write,
		}

		if !reflect.DeepEqual(expected, spyPrintSleeper.Calls) {
			t.Errorf("esperado %v chamadas, resultado %v", expected, spyPrintSleeper.Calls)
		}
	})

}

func TestConfigurationSleeper(t *testing.T) {
	timePause := 5 * time.Second
	timeSpy := &TimeSpy{}

	sleeper := ConfigurableSleeper{timePause, timeSpy.Pause}
	sleeper.Sleep()

	if timeSpy.pauseDuration != timePause {
		t.Errorf("pause duration does not match: expected value %v during %v", timeSpy.pauseDuration, timePause)
	}
}
