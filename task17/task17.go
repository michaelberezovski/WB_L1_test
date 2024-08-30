package task17

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 10, 14, 16, 17, 35, 58, 29, 46, 12, 98, 89, 74, 37}
	sort.Ints(array) // бинарный поиск работает только с отсортированным набором элементов
	target := 35
	result := binarySearch(array, target)
	if result < 0 {
		fmt.Println("Number not found", target)
	} else {
		fmt.Printf("Target %d has index %d \n", target, result)
	}
}

func binarySearch(array []int, n int) int {
	var result int
	middle := len(array) / 2
	switch {
	case len(array) == 0:
		result = -1
	case array[middle] > n:
		result = binarySearch(array[:middle], n)
	case array[middle] < n:
		result = binarySearch(array[middle+1:], n)
		if result >= 0 {
			result += middle + 1
		}
	default:
		result = middle
	}

	return result
}
