package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func check(e error)  {
  if e != nil {
    panic(e)
  }
}

type Course struct {
  Name     string `json:"coursename"`
  Price    float64 
  Password string `json:"-"`
  Tag      []string `json:"tags,omitempty"`
}

func encodeJson() {
  cources := []Course {
    {"Golang", 299.99, "Nothing0", []string{"web", "dev", "ops"}},
    {"JS", 199.99, "Nothing1", []string{"web", "dev", "ops"}},
    {"React", 199.99, "Nothing2", nil},
  }
  finalJson, err := json.MarshalIndent(cources, "", "\t")
  check(err)
  
  fmt.Printf("%s\n", finalJson)  
}

func getData() []byte{
  const myUrl = "http://localhost:8075/movies"
  
  resp, err := http.Get(myUrl)
  check(err)
  defer resp.Body.Close()

  content, err := io.ReadAll(resp.Body)
  check(err)
  
  return content
}

func decodeJson() {
  jsonData := []byte(`
  {
    "coursename" : "C++",
    "Price" : 399,
    "tags" : ["dev", "ai", "cloud"]
  }
  `)
  var course Course

  checkValid := json.Valid(jsonData)

  if checkValid {
    fmt.Println("Data is valid")
    json.Unmarshal(jsonData, &course)
    fmt.Printf("%#v\n", course)
  } else {
    fmt.Println("Json data invalid!!")
  }

  var dataMap map[string]interface{}
  json.Unmarshal(jsonData, &dataMap)
  fmt.Println(dataMap)
}

func main()  {
  decodeJson()
}
