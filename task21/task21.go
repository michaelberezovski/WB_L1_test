package task21

import "fmt"

type Adapter interface { //интерфейс, для адаптации методов из сторонних структур
	Operation()
}

type WeUseMethodsFromHere struct { //класс
}

func (a *WeUseMethodsFromHere) MethodOperation() {
	fmt.Println("I am method from WeUseMethodsFromHere")
}

type AdapterUserShell struct { //реализация под конкретный случай с наследованием структуры с нужными нам методами
	*WeUseMethodsFromHere
}

func (adapter *AdapterUserShell) Operation() { //создаём метод с реализацией интерфейса для вызова нужного нам метода
	adapter.MethodOperation()
}

func NewAdapter(a *WeUseMethodsFromHere) Adapter { //конструктор адаптера
	return &AdapterUserShell{a}
}

func main() {
	adapter := NewAdapter(&WeUseMethodsFromHere{})
	adapter.Operation() //вызов нужной операции
}
