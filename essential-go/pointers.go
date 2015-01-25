package main

import "fmt"

func main() {
  num := 5
  double(num)
  fmt.Println(num)

  triple(&num)
  fmt.Println(num)

  a := 1
  b := 10
  swap(&a, &b)
  fmt.Println(a, b)
}

func double(n int) {
  n = n * 2
}

func triple(n *int) {
  *n = *n * 3
}

func swap(a* int, b *int) {
  *b, *a = *a, *b
}
