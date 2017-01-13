package main
import "fmt"

var x = []int{
  48,96,86,68,
  57,82,63,70,
  37,34,83,27,
  19,97, 9,17,
}
for var i int = 0; i < len(x); i += 1 {
  fmt.Println(x[i]);
}
