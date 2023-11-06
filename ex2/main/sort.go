package main

//go env -w GO111MODULE=on
import (
	"fmt"
)

func main() {

	var n int
	var k int

	fmt.Print("Введите k: ")
	fmt.Scan(&k)

	fmt.Print("Введите количество элементов массива: ")

	fmt.Scan(&n)

	var mas []int = make([]int, n)

	fmt.Println("Введите элементы массива:")

	for i := 0; i < n; i++ {

		fmt.Print("Введите элемент ", i+1, ": ")

		fmt.Scan(&mas[i])

	}

	result := findKthLargest(mas, k)
	fmt.Print("Результат: \n", result)
}

func bubbleSort(mas []int) {
	len := len(mas)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if mas[j] > mas[j+1] {
				mas[j], mas[j+1] = mas[j+1], mas[j]
			}
		}
	}
}

func findKthLargest(mas []int, k int) int {

	bubbleSort(mas)
	var lengh = (len(mas))
	if k <= lengh {
		return mas[lengh-k]
	} else {
		return 0
	}
}
