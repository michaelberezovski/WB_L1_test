package task1

import "fmt"

type Human struct { //структура
	name string
	age  int
	sex  string
}

type Action struct { // структцра со вложенной Human для использования методов "родительской"
	actionName string
	Human
}

func (h Human) PrintNameAndAge() { // методы Human
	fmt.Printf("Human name and age %s (%d)\n", h.name, h.age)
}

func (h Human) PrintNameAndSex() {
	fmt.Printf("Human name and sex %s (%s)\n", h.name, h.sex)
}

func (a Action) PrintActionName() { // метод второй структуры
	fmt.Println(a.actionName)
}

func main() { //вызываем методы из всех структур у дочернего объекта
	var act = Action{actionName: "Print Human fields",
		Human: Human{"Antonio Banderas", 50, "male"}}
	act.PrintActionName()
	act.PrintNameAndAge()
	act.PrintNameAndSex()
}
