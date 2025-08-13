package main

import "fmt"

func calculator() {
	var num1, num2 float64
	var op string

	fmt.Print("Enter the First Number :  ")
	fmt.Scan(&num1)
	fmt.Print("Enter the Operation bw( +,-,*,/):")
	fmt.Scan(&op)
	fmt.Print("Enterr the Second Number :   ")
	fmt.Scan(&num2)

	switch op {
	case "+":
		fmt.Println("Result is:", num1+num2)
	case "-":
		fmt.Println("Result is:", num1-num2)
	case "*":
		fmt.Println("Result is:", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Println("Result is:", num1/num2)
		} else {
			fmt.Println("Error: never divide by zero")
		}

	default:
		fmt.Println("Enter the correct operation")
	}
}
