package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	for {

		var a int
		var b int
		var op string
		var result float64
		var input string

		fmt.Println("Это калькулятор")
		fmt.Println("Пока он может выполнять только простейшие функции")
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("Если захотите выйти - введите exit вместо числа")
		fmt.Println("Можем начинать!")
		fmt.Println("-----------------")
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("Введите первое число:")
		fmt.Scanln(&input)

		if input == "exit" {
			break
		}

		a, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Введите целое число!")
			break
		}

		for {
			fmt.Println("выберите действие(+, -, * или /)")
			fmt.Scanln(&op)

			if op == "exit" {
				return
			}

			if op != "+" && op != "-" && op != "*" && op != "/" {
				fmt.Println("Неверный знак. Выберите из +, -, *, /")
				continue
			}
			break
		}

		fmt.Println("Введите второе число:")
		fmt.Scanln(&input)

		if input == "exit" {
			continue
		}

		b, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Введите целое число!")
			continue
		}

		switch op {
		case "+":
			result = float64(a + b)
		case "-":
			result = float64(a - b)
		case "*":
			result = float64(a * b)
		case "/":
			if b == 0 {
				fmt.Println("На 0 делить нельзя")
				continue
			}
			result = float64(a) / float64(b)
		}
		fmt.Printf("Ответ: %.2f\n", result)
		fmt.Println("-----------------")

		continue
	}

}
