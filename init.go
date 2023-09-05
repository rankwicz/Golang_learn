// Объявление переменной

package main

import (
	"fmt"
)

func main() {
	var age int = 23
	var age uint = 23 // u означает что число не может быть меньше нуля
	var number float64 = 254.215
	var name string = "Denis"

	// Упрощённый способ объявления переменной
	age := 20
	name := "Denis"

	// Узнать длину строки
	fmt.Print(len(name))
	fmt.Print(age)
}
