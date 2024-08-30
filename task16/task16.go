package task16

import "fmt"

func main() {
	nums := []int{12, 94, 7, -8, -1, 0, 283, 22, 57, 16}
	quickSort(nums)
	fmt.Println(nums)
}

func quickSort(nums []int) {

	if len(nums) <= 1 {
		return
	}

	split := splitter(nums)

	//рекурсивный вызов ффункции
	quickSort(nums[:split])
	quickSort(nums[split:])
}

func splitter(nums []int) int {

	//разбиваем слайс, создаем точку относительно которой будет производиться сравнение и перещение
	pivot := nums[len(nums)/2]

	//устанавливаем точки начала и конца слайса
	left := 0
	right := len(nums) - 1
	//проходим циклом, элементы меньше pivot перемещаются влево, а больше - вправо
	for {
		for ; nums[left] < pivot; left++ {
		}
		for ; nums[right] > pivot; right-- {
		}

		if left >= right {
			return right
		}
		swap(nums, left, right)
	}

}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
