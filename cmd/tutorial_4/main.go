package main

import "fmt"

func main() {
  intArr := [3]int32{2,3,4} // array
  fmt.Println(intArr[2])
  fmt.Println(len(intArr))

  intSlice := []int32{4, 5, 8} // slice
  fmt.Println(intSlice[2])
  fmt.Println(len(intSlice))

  /**
  Removing a number (3) from the array definition makes it a new data structure 😭
  */
}
