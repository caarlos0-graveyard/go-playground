package main
import "fmt"

func main() {
  if true == false {
    fmt.Println("true is true")
  } else {
    fmt.Println("true is false")
  }

  for i := 0; i < 20; i++ {
    fmt.Println("Go!")
  }

  for {
    fmt.Println("This will go forever")
    break
  }

  j := true
  for j {
    fmt.Println("This should happen once only")
    j = false
  }
}
