// Приём ввода от пользователя

package main

import (
	"fmt"
)

func main() {
	var name string
	var age int
	fmt.Println("What is your name?")
	fmt.Scan(&name)
	fmt.Println("Hello " + name + "!")
	fmt.Println("How old are you? ")
	fmt.Scan(&age)
	fmt.Println("You are " + fmt.Sprint(age) + " years old")

	// Функция Println автоматически перемещает текст на следующую строку, в то время как Printf и Print этого не делают.
	// При желании переместить что-то на новую строку в тексте нужно добавить \n.
	//Если уточняется несколько специальных символов, функция Printf подставит значения в следующем порядке:
	fmt.Printf("Меня возьмут на %v в %v", "стажировку", "WB")

}
