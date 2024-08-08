package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func check(e error)  {
  if e != nil {
    panic(e)
  }
}

func performGetRequest()  {
  const myUrl = "http://localhost:8075/movies"

  resp, err := http.Get(myUrl)
  
  check(err)
  defer resp.Body.Close()
  fmt.Println("Status Code:", resp.StatusCode)
  fmt.Println("Content Length:", resp.ContentLength)

  content, err := io.ReadAll(resp.Body)
  check(err)

  var respString strings.Builder
  byteCount, err := respString.Write(content)
  check(err)
  fmt.Println("Byte Count:",byteCount)
  fmt.Println(respString.String())

}


func performPostRequest()  {
  const myUrl = "http://localhost:8075/movie"

  //Json Data

  reqBody := strings.NewReader(`
    {
      "id": "99",
        "isbn": "12345",
        "title": "Star",
        "lead": {
            "fname": "Me",
            "lname": "llll"
        }
    }
  `)
  resp, err := http.Post(myUrl, "application/json", reqBody)
  check(err)
  defer resp.Body.Close()

  content, err := io.ReadAll(resp.Body)
  check(err)

  var respString strings.Builder
  byteCount, err := respString.Write(content)
  check(err)

  fmt.Println("Byte Count:", byteCount)
  fmt.Println(respString.String())
}


func performPostForm()  {
  const myUrl = "http://localhost:8000/postform"

  data := url.Values{}
  data.Add("fname", "Nisarg")
  data.Add("lname", "Khodke")
  data.Add("email", "nisarg@tmp.in")

  resp, err := http.PostForm(myUrl, data)
  check(err)  
  defer resp.Body.Close()

  content, err := io.ReadAll(resp.Body)
  check(err)

  var respString strings.Builder
  byteCount, err :=  respString.Write(content)
  check(err)

  fmt.Println("Byte Count:", byteCount)
  fmt.Println(respString.String())
}

func main()  {
  // performGetRequest()
  // performPostRequest()
  performPostForm()
}
