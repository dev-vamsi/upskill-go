package main

import (
    "fmt"
    "strings"
)

func main() {
    // WTF 😭 // UTF-8 Encoding 🙌
    
    bullet := "ßµllet"
    indexed := bullet[0]
    fmt.Printf("%v, %T \n", indexed, indexed)
    
    for i, v := range bullet {
        fmt.Printf("%v, %v \n", i, v)
    }
    
    fmt.Println(bullet[1])
    
    bulletRune := []rune("ßµllet")
    for i, v := range bulletRune {
        fmt.Printf("%v, %v \n", i, v)
    }
    
    // string builder
    strSlice := []string{"h", "e", "l", "l", "o"}
    
    catStr := ""
    for _, v := range strSlice {
        catStr += v // everytime a new string is built
    }
    fmt.Println(catStr)
    
    var sb strings.Builder
    for _, v := range strSlice {
        sb.WriteString(v)
    }
    fmt.Println(sb.String(), "using sb")
}
