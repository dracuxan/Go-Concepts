package mocking

import (
	"fmt"
	"io"
	"time"
)

const (
	finalWord  = "Go!"
	countStart = 3
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(buff io.Writer, s Sleeper) {
	for i := countStart; i > 0; i-- {
		fmt.Fprintln(buff, i)
		s.Sleep()
	}
	fmt.Fprint(buff, finalWord)
}
