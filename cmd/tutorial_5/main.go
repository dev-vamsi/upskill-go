package main

import "fmt"

func main() {
    var myMap map[string]uint32 = make(map[string]uint32)
    fmt.Println(myMap)
    
    myMap1 := map[string]uint32{"Adam": 22, "Sarah": 28, "John": 32}
    fmt.Println(myMap1)
    
    var age, ok = myMap1["Adam"]
    fmt.Println(age, ok)
    
    var _age, _ok = myMap1["Vamsi"]
    fmt.Println(_age, _ok)
    
    delete(myMap1, "Adam")
    fmt.Println(myMap1)
    fmt.Println("---------------")
    for name, age := range myMap1 {
        fmt.Printf("Name: %v, Age: %v \n", name, age)
    }
    fmt.Println("---------------")
    intArr := []int32{10, 20, 30, 40}
    for i, v := range intArr {
        fmt.Printf("Index: %v, Value: %v \n", i, v)
    }
    fmt.Println("---------------")
    for i := 0; i < 10; i++ {
        fmt.Printf("%v - ", i)
    }
    fmt.Println(10)
    fmt.Println("---------------")
    i := 0
    for {
        if (i > 10) {
            break
        }
        fmt.Printf("%v - ", i)
        i++
    }
}
