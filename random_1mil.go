package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  for i := 0; i < 1000000; i++ {
    fmt.Println(rand.Intn(1000000))
  }
}
