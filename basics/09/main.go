package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func check(e error)  {
  if e != nil {
    panic(e)
  }
}

func performGet(myUrl string)  {
  resp, err := http.Get(myUrl)
  check(err)
  defer resp.Body.Close()

  content, err := io.ReadAll(resp.Body)
  check(err)

  var respStr strings.Builder
  byteCount, err := respStr.Write(content)
  fmt.Println("Byte Count:", byteCount)
  fmt.Println(respStr.String())
}

func main()  {
  myUrl := "http://localhost:8000"
  performGet(myUrl)
}
