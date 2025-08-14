package arithmetic

import "fmt"

func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// func MathOper() (float64, error) {
// 	var num1, num2 float64
// 	var op string

// 	fmt.Print("Enter the First Number :  ")
// 	fmt.Scan(&num1)
// 	fmt.Print("Enter the Operation bw( +,-,*,/):")
// 	fmt.Scan(&op)
// 	fmt.Print("Enter the Second Number :   ")
// 	fmt.Scan(&num2)

// 	switch op {
// 	case "+":
// 		return num1 + num2, nil
// 	case "-":
// 		return num1 - num2, nil
// 	case "*":
// 		return num1 * num2, nil
// 	case "/":
// 		if num2 != 0 {
// 			return num1 / num2, nil
// 		}
// 		return 0, fmt.Errorf("division by zero is not allowed")
// 	default:
// 		return 0, fmt.Errorf("invalid operator")
// 	}
// }

var Num1 int = 20
