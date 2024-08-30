package task23

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 10, 14, 16, 17, 35, 58, 29, 46, 12, 98, 89, 74, 37}

	fmt.Println("Original slice:", arr)

	var index int
	fmt.Printf("Enter the index of the element to delete 0 to %d: ", len(arr)-1)
	_, err := fmt.Scanf("%d", &index)
	if err != nil || index < 0 || index >= len(arr) {
		fmt.Println("Invalid index. Please enter a valid index.")
		return
	}

	arr = deleteElement(arr, index)

	fmt.Println("Modified slice:", arr)
}

// функция возвращает два склеенных слайса без удаляемого элемента
func deleteElement(slice []int, index int) []int {

	if index < 0 || index >= len(slice) {
		return slice
	}

	return append(slice[:index], slice[index+1:]...)
}
