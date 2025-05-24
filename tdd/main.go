package main

import (
	"os"
	"time"

	"Go-Concepts/tdd/mocking"
)

func main() {
	sleeper := mocking.NewConfigurableSleeper(250*time.Millisecond, time.Sleep)
	mocking.Countdown(os.Stdout, sleeper)
}
