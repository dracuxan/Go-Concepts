package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
  fmt.Println("Enter a number from 2-200:")

  reader := bufio.NewReader(os.Stdin)
  inpt, _ := reader.ReadString('\n')

  fmt.Println(inpt)

  numInp, err := strconv.ParseFloat(strings.TrimSpace(inpt), 64)
  if err != nil {
    panic(err)
  }

  fmt.Printf("Type of input: %T \n", numInp)
}
