package main

import (
	"fmt"
	"time"
)

//	func main() {
//		myChan := make(chan string, 3)
//
//		chars := []string{"a", "b", "c"}
//
//		for _, c := range chars {
//			myChan <- c
//		}
//
//		close(myChan)
//
//		for res := range myChan {
//			fmt.Println(res)
//		}
//	}
func work(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("dddd")
		}
	}
}

func main() {
	done := make(chan bool)

	go work(done)
	time.Sleep(time.Second * 3)

	close(done)
}
