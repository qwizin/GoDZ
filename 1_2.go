package main

import (
	"fmt"
)

type Employee struct {
	Name     string // имя
	Age      int    // возраст
	Position string // позиция
	Salary   int    // зарплата
}

var commands = `
1 - Добавить нового сотрудника
2 - Удалить сотрудника
3 - Вывести список сотрудников
4 - Выйти из программы
`

func main11() {
	const size = 512
	empls := [size]*Employee{}
	for {
		cmd := 0
		fmt.Print(commands)
		fmt.Scanf("%d", &cmd)

		switch cmd {
		case 1:
			// Добавляем нового сотрудника
			empl := new(Employee)
			fmt.Println("\nИмя:")
			fmt.Scanf("%s", &empl.Name)
			fmt.Println("Возраст:")
			fmt.Scanf("%d", &empl.Age)
			fmt.Println("Позиция:")
			fmt.Scanf("%s", &empl.Position)
			fmt.Println("Зарплата:")
			fmt.Scanf("%d", &empl.Salary)
			added := false
			for i := 0; i < size; i++ {
				if empls[i] == nil {
					empls[i] = empl
					added = true
					fmt.Println("Сотрудник добавлен.")
					break
				}
			}
			if !added {
				fmt.Println("Ошибка: достигнут лимит сотрудников.")
			}
		case 2:
			// Удаляем сотрудника
			fmt.Println("Введите индекс сотрудника для удаления:")
			var index int
			fmt.Scanf("%d", &index)
			if index >= 0 && index < size && empls[index] != nil {
				empls[index] = nil
				fmt.Println("Сотрудник удален.")
			} else {
				fmt.Println("Ошибка: некорректный индекс.")
			}
		case 3:
			// Вывод списка сотрудников
			fmt.Println("Список сотрудников:")
			for i, empl := range empls {
				if empl != nil {
					fmt.Printf("%d: %s, %d лет, %s, зарплата: %d\n", i, empl.Name, empl.Age, empl.Position, empl.Salary)
				}
			}
		case 4:
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Ошибка: неизвестная команда.")
		}
	}
}
