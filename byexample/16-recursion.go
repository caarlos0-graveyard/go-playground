package main

import "fmt"

func fact(n int) int {
  if n == 0 {
    return 1
  }
  return n * fact(n-1)
}

func fibb(n int, memo map[int]int) int {
  if _, ok := memo[n]; ok {
    return memo[n]
  }
  memo[n] = fibb(n - 1, memo) + fibb(n - 2, memo)
  return memo[n]
}

func main() {
  fmt.Println("fact(20): ", fact(20))
  memo := make(map[int]int)
  memo[0] = 0
  memo[1] = 1
  f := fibb(100, memo)
  fmt.Println("Fibb(100): ", f)
}
