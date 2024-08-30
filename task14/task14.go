package task14

import "fmt"

/*
Разработать программу, которая в рантайме способна определить
тип переменной: int, string, bool, channel из переменной типа interface{}.
*/

type TypeCheck interface{} //пустой интерфейс может содержать данные любого типа

func checkType(t TypeCheck) string { //со входным аргументом TypeCheck функция так же может принимть данные любого типа
	return fmt.Sprintf("%T", t) //в данном случае мы печатаем тип данных
}

func main() {
	var t TypeCheck //создаем переменную тпа тайпчек

	//далее вводим произвольный тип данных и пользуемся созданной выше функцией для отображения типа этих данных
	t = 1
	fmt.Printf("Type: %v\n", checkType(t))

	t = 1.1
	fmt.Printf("Type: %v\n", checkType(t))

	t = "some txt"
	fmt.Printf("Type: %v\n", checkType(t))

	t = false
	fmt.Printf("Type: %v\n", checkType(t))

	t = make(chan bool)
	fmt.Printf("Type: %v\n", checkType(t))

}
