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

func findKthLargest(mas []int, k int) int {

	var temp int
	var lengh = (len(mas))
	if k <= lengh {
		var i = k - 1
		temp = mas[k-1]
		for {
			if temp < mas[i] {
				temp = mas[i]
			}
			i = i + k
			if i > lengh-1 {
				break
			}
		}
	} else {
		temp = 0
	}

	return temp

}
