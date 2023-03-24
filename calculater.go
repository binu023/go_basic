package main

import "fmt"

func add(x float64, y float64) float64 {
	return x + y
}
func subtract(x float64, y float64) float64 {
	return x - y
}
func multiply(x float64, y float64) float64 {
	return x * y
}
func divide(x float64, y float64) float64 {
	return x / y

}

func calculator(x float64, y float64, operator string) float64 {
	if operator == string('+') {
		return add(x, y)
	} else if operator == string('*') {
		return multiply(x, y)
	} else if operator == string('-') {
		return subtract(x, y)
	} else {
		return divide(x, y)
	}
}

func main() {
	fmt.Println(calculator(2, 2, string('-')))
}
