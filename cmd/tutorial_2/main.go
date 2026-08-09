package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int
	fmt.Println(intNum)

	var float1 float32 = 12345678.9
	var float2 float64 = 12345678.9
	fmt.Println(float1)
	fmt.Println(float2)

	var float3 float32 = 122
	var int1 int32 = 10
	var res float32 = float3 + float32(int1)
	fmt.Println(res)

	var myString1 = "Hello World!"
	fmt.Println(myString1)

	var myString2 = "Hello \nWorld!"
	fmt.Println(myString2)
	var myString3 = `Hello
World!`
	fmt.Println(myString3)

	fmt.Println(len(myString2))
	fmt.Println(len("γ"))                    // 2
	fmt.Println(utf8.RuneCountInString("γ")) // 1

	var myRune rune = 'a'
	fmt.Println(myRune) // 97

	myVar := "Test"
	fmt.Println(myVar)

	var var1, var2 int = 1, 2
	fmt.Println(var1, var2)
	var3, var4 := 3, 4
	fmt.Println(var3, var4)

	const myConst int = 138493
	fmt.Println(myConst)
}
