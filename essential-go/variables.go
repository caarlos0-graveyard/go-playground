package main

import "fmt"

var (
  author = "@caarlos0"
  // uppercase methods/vars are public!
  Version = "0.0.1"
)

const (
  CCVisa = "Visa"
  CCMasterCard = "MasterCard"
  CCAmericanExpress = "AmericanExpress"
)

func main() {
  var greetings string = "Hello World!"
  var a, b, c int = 1, 2, 3
  fmt.Println(greetings, a, b, c)

  var name = "Carlos A Becker"
  var d, e, f = 1, 2.0, "Three"
  fmt.Println(name, d, e, f)

  course := "Essential Go!"
  x, y, z := 1, 2, 3
  fmt.Println(course, x, y, z)


}
