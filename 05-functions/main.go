package main

import (
	"errors"
	"fmt"
)

func add(a int, b int) int {
	return a + b
}


func divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

func main () {
	firstNum :=10
	secondNum :=2
	sum := add(firstNum,secondNum)
	fmt.Println("sum ",firstNum,"+",secondNum,"=", sum)

	result,err := divide(float64(firstNum),float64(secondNum))
	if err != nil {
		fmt.Print("error while divide ",firstNum,"/",secondNum,"=", err)
		return
	} 
	fmt.Println("devide " , firstNum , "/"  , secondNum , "=", result)
}