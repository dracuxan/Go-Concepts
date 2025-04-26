package main

import (
	"os"

	"Go-Concepts/tdd/mocking"
)

func main() {
	s := &mocking.DefaultSleeper{}
	mocking.Countdown(os.Stdout, s)
}
