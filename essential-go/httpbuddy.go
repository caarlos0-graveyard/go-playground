package main

import (
  "fmt"
  "flag"
  "net/http"
  "log"
)

var port int

func main() {
  flag.IntVar(&port, "port", 3000, "The port to run the server on")
  flag.Parse()

  fmt.Printf("Serving file on localhost:%v\n", port)
  err := ServeStatic(port)
  if err != nil {
    log.Fatalln(err)
  }
}


func ServeStatic(port int) error {
  host := fmt.Sprintf("localhost:%v", port)
  return http.ListenAndServe(host, http.FileServer(http.Dir(".")))
}
