package task22

import (
	"fmt"
	"math/big"
)

func main() {
	//поскольку числа большие пользуемся стандарным пакетом "math/big"
	num1 := big.NewInt(1 << 20)
	num2 := big.NewInt(1 << 21)

	sum := new(big.Int).Add(num1, num2)
	fmt.Printf("Sum: %s\n", sum.String())

	difference := new(big.Int).Sub(num2, num1)
	fmt.Printf("Difference: %s\n", difference.String())

	multiply := new(big.Int).Mul(num1, num2)
	fmt.Printf("Multiply: %s\n", multiply.String())

	division := new(big.Int).Div(num2, num1)
	fmt.Printf("Division: %s\n", division.String())
}
