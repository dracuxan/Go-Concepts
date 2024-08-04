package main

import (
	"fmt"
	"time"
)

func main() {
  pTime := time.Now()
  fmt.Println(pTime)

  fmt.Println(pTime.Format("02-01-2006"))
}
