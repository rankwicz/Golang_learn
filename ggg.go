package main

import (
	"fmt"
)

func main() {
	var relax string
	fmt.Println("Что ты хочешь сделать? \nОтдохнуть или поесть")
	fmt.Scan(&relax)
	if relax == "отдохнуть" {
		fmt.Println("Иди спать")
	} else if relax == "поесть" {
		fmt.Println("Иди поешь")
	} else {
		fmt.Println("Ничего не понятно")
	}

}
