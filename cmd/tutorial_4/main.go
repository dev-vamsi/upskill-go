package main

import "fmt"

func main() {
  intArr := [3]int32{2,3,4} // array
  fmt.Println(intArr[2])
  fmt.Println(len(intArr))

  /**
  - Removing a number (3) from the array definition makes it a new data structure 😭.
  - Capacity will be doubled when size of slice and cap are equal and tries to add a new element.
  */
  intSlice := []int32{4, 5, 8} // slice
  fmt.Println(intSlice[2])
  fmt.Println(len(intSlice))
  fmt.Println(len(intSlice), cap(intSlice))
  intSlice = append(intSlice, 10)
  fmt.Println(len(intSlice), cap(intSlice))
  
  intNewArr := []int32{8,9}
  intSlice = append(intSlice, intNewArr...)
  fmt.Println(intSlice)
  fmt.Println(len(intSlice), cap(intSlice))
  
    
  // creating slice using make command
  intSlice1 := make([]int32, 2, 8)
  fmt.Println(intSlice1)
  fmt.Println(len(intSlice1), cap(intSlice1))
}
