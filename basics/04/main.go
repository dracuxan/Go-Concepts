package main

import "fmt"

func main()  {
  varr := 2
  ptr := &varr

  fmt.Println(*ptr)
  fmt.Println(&varr)
  fmt.Println(ptr)
  fmt.Println(&ptr)
}
