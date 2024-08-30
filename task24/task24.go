package task24

import (
	"fmt"
	"math"
)

type Point struct { //доступ к (x y) только внутри пакета
	x, y float64
}

func createPoint(x, y float64) *Point { // конструктор
	return &Point{x: x, y: y}
}

func calculateDistance(x1, y1, x2, y2 float64) float64 { //функция подсчёта расстояния
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func main() {
	point1 := createPoint(10.0, 10.0)
	point2 := createPoint(100.0, 100.0)
	fmt.Printf("Distance is %f\n", calculateDistance(point1.x, point1.y, point2.x, point2.y))
}
