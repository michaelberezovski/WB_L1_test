package task8

import (
	"fmt"
)

func changeBit(num int64, position int, bit int) int64 { //функция для изменения бита в числе на заданной позиции

	bitMask := int64(1) << position //создаём битовую маску для дальейших манипуляций

	if bit == 1 {
		num |= bitMask // используем bitwise OR оператор, чтобы результатом операции был включенный бит
	} else {
		num &= ^bitMask // тут используем bitwise AND вместе с NOT, чтобы очистить бит
	}

	return num
}

func main() {

	var num int64 = 256
	position := 3
	bit := 1

	fmt.Printf("Original number: %d\n", num)

	num = changeBit(num, position, bit)

	fmt.Printf("Result of position #%d bit change to %d is: %d\n", position, bit, num)
}
