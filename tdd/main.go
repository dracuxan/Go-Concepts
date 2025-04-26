package main

import (
	"os"
	"time"

	"Go-Concepts/tdd/mocking"
)

func main() {
	sleeper := mocking.NewConfigurableSleeper(1*time.Second, time.Sleep)
	mocking.Countdown(os.Stdout, sleeper)
}
