package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

// Объявление методов структуры Human. Методы привязаны в этой структуре посредством (h *Human) перед названием функции.
func (h *Human) GetName() string {
	return h.Name
}

// Метод GetAge для структуры Human
func (h *Human) GetAge() int {
	return h.Age
}

// Структура Action с встраиванием методов от структуры Human.
// Передача с * означает передачу ссылки на структуру, а не копию. Тогда изменения в human будут доступны в Action.
type Action struct {
	*Human // Встраиваем структуру Human - передает поля и методы в Action.
}

// Метод PrintInfo для структуры Action, который использует методы из структуры Human
func (a *Action) PrintInfo() {
	fmt.Printf("Name: %s, Age: %d\n", a.GetName(), a.GetAge())
}

func main() {
	// Создаем объект типа Action
	human := Human{Name: "Alice", Age: 30}
	action := Action{Human: &human} // передаем указатель на структуру Human.

	// Вызываем метод PrintInfo для объекта action, который использует методы из структуры Human
	action.PrintInfo()
}
