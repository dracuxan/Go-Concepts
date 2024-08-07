package main

import (
	"fmt"
	"net/http"
	"net/url"
)

const Url = "http://localhost:8075/movies?tab=repo&movie=1"

func check(e error)  {
  if e != nil {
    panic(e)
  }
}

func main()  {
  resp, err := http.Get(Url)
  check(err)
  defer resp.Body.Close()

  data, err := url.Parse(Url)
  check(err)
  qparam := data.Query()
  fmt.Println(qparam["tab"])
  
  madeUrl := &url.URL{
    Scheme: "http",
    Host: "localhost:8075",
    Path: "/movies",
    RawPath: "movie=no",
  }
  fmt.Println(madeUrl.String())
}
