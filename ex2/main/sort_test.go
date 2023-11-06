package main

import "testing"

func TestSort(t *testing.T) {

	testCases := []struct {
		mas    []int
		k      int
		result int
	}{
		{[]int{100, 22, 43, 14, 15}, 2, 22},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, 6},
		{[]int{1, 1, 1, 1, 51, 1, 1, 1, 1, 10}, 5, 51},
		{[]int{11, 22, 12, 31, 22, 2, 31, 24, 34, 23, 44, 1, 65, 45, 67}, 3, 67},
		{[]int{1, 2, 3, 4, 5}, 10, 0},
	}

	for _, testCase := range testCases {
		results := findKthLargest(testCase.mas, testCase.k)
		if results != testCase.result {
			t.Errorf("Данные: массив = %v , k =%d . Результат: %d", testCase.mas, testCase.k, results)
		}
	}
}
