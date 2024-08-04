package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error)  {
  if e != nil {
    panic(e)
  }
}

func main() {
  // Taking user inputs
  welcome := "Welcome to this program?"
  fmt.Println(welcome)
  
  reader := bufio.NewReader(os.Stdin)
  fmt.Println("Enter a number from 1-10:")

  inp, e := reader.ReadString('\n')
  check(e)
  fmt.Println("The bumber you entered:", inp)
}
