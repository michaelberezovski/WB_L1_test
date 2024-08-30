package task10

import (
	"fmt"
	"math"
)

func main() {

	dataTemp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	temperatures := make(map[float64][]float64) //мапа со слайсом в поле value, чтобы добавлять новые значения по мере вычислений

	for _, temp := range dataTemp {
		key := math.Trunc(temp/10.0) * 10.0                 //делим без остатка на 10 для получения значения ключа
		temperatures[key] = append(temperatures[key], temp) //добавляем значения в мапу по ключу
	}

	for key, value := range temperatures {
		fmt.Println(key, value)
	}

}
