package main

import (
  "fmt"
  "math"
  "github.com/Manik5/go_crash_course/03_packages/strutil"
)

func main() {
  fmt.Println(math.Floor(2.7))
  fmt.Println(math.Ceil(2.7))
  fmt.Println(math.Sqrt(2.7))
  fmt.Println(strutil.Reverse("olleh"))
}
