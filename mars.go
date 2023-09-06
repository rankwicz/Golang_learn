package main

import "fmt"

func main() {
	const speed = 100800
	distance := 96300000
	point := distance / speed
	days := 24 * point
	years := days / 365
	day := years % 365
	fmt.Printf("Астронавты долетят за %v дней или %v года и %v дней", days, years, day)
}
