package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	var op string
	var result float64

	fmt.Println("Это калькулятор")
	fmt.Println("Пока он может выполнять только простейшие функции")
	time.Sleep(1500 * time.Millisecond)
	fmt.Println("Если захотите выйти - введите exit вместо числа")
	fmt.Println("Можем начинать!")
	fmt.Println("-----------------")
	time.Sleep(1500 * time.Millisecond)

	a, err := readNumber("Введите первое число:")
	if err != nil {
		if err.Error() == "exit" {
			return
		}
		fmt.Println(err)
		return
	}

	op, err = readOperator("Выберите действие (+, -, *, /):")
	if err != nil {
		if err.Error() == "exit" {
			return
		}
		fmt.Println(err)
		return
	}

	b, err := readNumber("Введите второе число:")
	if err != nil {
		if err.Error() == "exit" {

		}
		fmt.Println(err)
		return
	}

	result, err = calculate(a, b, op)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Printf("Ответ: %.2f\n", result)
	fmt.Println("------")

}

func calculate(a float64, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("На ноль делить нельзя")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("неизвестный оператор: %s", operator)
	}
}

func readNumber(prompt string) (float64, error) {
	var input string
	for {
		fmt.Println(prompt)
		fmt.Scanln(&input)

		if input == "exit" {
			return 0, fmt.Errorf("exit")
		}

		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Printf("некорректный ввод: %s\n", input)
			continue
		}

		return num, nil
	}
}

func readOperator(prompt string) (string, error) {
	var input string
	for {
		fmt.Println(prompt)
		fmt.Scanln(&input)

		if input == "exit" {
			return "", fmt.Errorf("exit")
		}

		if input == "+" || input == "-" || input == "*" || input == "/" {
			return input, nil
		}
		fmt.Println("Неверный знак. Выберите из +, -, *, /")
	}

}
