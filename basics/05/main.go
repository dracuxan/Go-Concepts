package main

import "fmt"

func main()  {
  U1 := User{"Nisarg", "n@gmail.com", 1, true}
  U1.GetStatus()
}

type User struct {
  Name string
  Email string
  Number int
  Status bool
}

func (u User) GetStatus() {
  setStatus := "Offline"
  if u.Status == true {
    setStatus = "Online"
  }
  fmt.Printf("Staus of user is %v\n", setStatus)
}
