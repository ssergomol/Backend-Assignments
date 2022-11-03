package main

import "fmt"

// Родительская структура Human, которая имеет свои поля и методы
type Human struct {
	name   string
	age    int
	energy int
}

func (h Human) GetHumanName() string {
	return h.name
}

func (h Human) GetHumanAge() int {
	return h.age
}

func (h Human) GetEnergy() int {
	return h.energy
}

// Дочерняя структура Action, которая встраивает структуру Human.
// Таким образом, происходит наследование методов и полей структуры Human
type Action struct {
	Human
	actionName string
	energyCost int
}

func (a *Action) DoAction() {
	if a.energy < a.energyCost {
		fmt.Println("Too tired!")
		return
	}
	a.energy -= a.energyCost
	fmt.Println("Energy remained: ", a.energy)
}

func main() {
	human := Human{
		name:   "John Doe",
		age:    20,
		energy: 20,
	}

	action := Action{
		Human:      human,
		actionName: "run",
		energyCost: 10,
	}

	// Реализовав встраивание, можем вызывать методы структуры Human у Action
	fmt.Printf("name: %s, age: %d\n", action.GetHumanName(), action.GetHumanAge())

	action.DoAction()
	action.DoAction()
	action.DoAction()
}
