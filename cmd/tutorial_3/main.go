package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe("Vamsi Krishna")

	int1, int2 := 22, 3
	fmt.Println(intDivision(int1, int2))
	fmt.Println(intDivisionPair(int1, int2))

	res, _err := intDivision(10, 0)
	// TODO: rewrite below logic using switch case
	if _err != nil {
		fmt.Println(_err.Error())
	} else {
		fmt.Println(res)
	}
}

func printMe(value string) {
	fmt.Println(value)
}

func intDivision(int1 int, int2 int) (int, error) {
	var err error
	if int2 == 0 {
		err := errors.New("Denominator cannot be zero")
		return 0, err
	}
	return int1 / int2, err
}

func intDivisionPair(int1 int, int2 int) (int, int) {
	return int1 / int2, int1 % int2
}
