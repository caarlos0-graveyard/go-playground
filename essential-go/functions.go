package main

import (
  "fmt"
  "strings"
  "log"
)

func main() {
    fmt.Println(double(5))
    first, _ := parseName("Carlos Becker")
    fmt.Println(first)

    greet := func(name string) {
      fmt.Println("Hello", name)
    }

    greet(first)

    log.Println("treta")
    fmt.Println(adder()(1, 2))
}

func double(n int) int {
  return n * 2
}

func parseName(name string) (first, last string) {
  parsed := strings.Split(name, " ")
  return parsed[0], parsed[1]
}

func adder() func(a, b int) int {
  return func(a, b int) int {
    return a + b
  }
}
