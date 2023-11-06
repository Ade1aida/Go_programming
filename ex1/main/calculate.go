package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Print("Введите первое число: ")

	var temp string
	expression := []string{}

	fmt.Scanln(&temp)
	expression = append(expression, temp)

	fmt.Print("Выберите операцию (+, -, *, /): ")
	fmt.Scanln(&temp)
	expression = append(expression, temp)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&temp)
	expression = append(expression, temp)

	result := calculate(expression)
	fmt.Printf("Результат: %f\n", result)
}

func calculate(expression []string) float64 {

	if len(expression) != 3 {
		fmt.Println("Некорректное выражение.")
		os.Exit(1)
	}

	operand1, err := strconv.ParseFloat(expression[0], 64)
	if err != nil {
		fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.", err)
		os.Exit(1)
	}

	operator := expression[1]

	operand2, err := strconv.ParseFloat(expression[2], 64)
	if err != nil {
		fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.", err)
		os.Exit(1)
	}

	var result float64
	switch operator {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		if operand2 == 0 {
			fmt.Println("Деление на ноль невозможно.")
			os.Exit(1)
		}
		result = operand1 / operand2
	default:
		fmt.Println("Некорректная операция ", operator, " Пожалуйста, используйте символы +, -, * или /.")
		os.Exit(1)
	}

	return result
}
