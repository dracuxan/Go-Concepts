package main

import (
	"fmt"
	"io"
	"os"
)

func check(e error)  {
  if e != nil {
    panic(e)
  }
}

func writeFile(content, fileName string) int  {
  file, err := os.Create(fileName)
  check(err)
  l, err := io.WriteString(file, content)
  check(err)
  defer file.Close()
  return l
}

func readFile(filePath string) {
  data, err := os.ReadFile(filePath)
  check(err)
  fmt.Printf("Data inside the file is:\n%v", string(data))
}

func main()  {
  fileName := "./test.txt"
  content := "New content.\n"
  length := writeFile(content, fileName)
  fmt.Printf("File created successfully of length: %v", length)
  readFile(fileName)
}
