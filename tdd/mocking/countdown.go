package mocking

import (
	"fmt"
	"io"
	"time"
)

const (
	finalWord  = "Go!"
	countStart = 3
	sleep      = "sleep"
	write      = "write"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func NewConfigurableSleeper(
	duration time.Duration,
	sleep func(time.Duration),
) *ConfigurableSleeper {
	return &ConfigurableSleeper{
		duration: duration,
		sleep:    sleep,
	}
}

func Countdown(buff io.Writer, s Sleeper) {
	for i := countStart; i > 0; i-- {
		fmt.Fprintln(buff, i)
		s.Sleep()
	}
	fmt.Fprint(buff, finalWord)
}
