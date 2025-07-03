package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	var op string
	var result float64

	fmt.Println("Введите первое число:")
	fmt.Scanln(&a)
	fmt.Println("выберите знак")
	fmt.Scanln(&op)
	if op != "+" && op != "-" && op != "*" && op != "/" {
		fmt.Println("Неверный знак. Выберите из +, -, *, /")
	}
	fmt.Println("Введите второе число:")
	fmt.Scanln(&b)

	switch op {
	case "+":
		result = float64(a + b)
	case "-":
		result = float64(a - b)
	case "*":
		result = float64(a * b)
	case "/":
		result = float64(a / b)
	}

	fmt.Printf("Ответ: %.2f\n", result)
}
