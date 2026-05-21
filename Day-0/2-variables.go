package main

import "fmt"

func main() {
	var name string = "Alice"
	var age int = 30
	var isEmployed bool = true

	var city string = "New York"

	var population int = 8000000
	var isdeveloped bool = true
	var temperature float64 = 22.5

	fmt.Println("----------------------------------")
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Employed:", isEmployed)
	fmt.Println("----------------------------------")
	fmt.Println("City:", city)
	fmt.Println("Population:", population)
	fmt.Println("Is Developed:", isdeveloped)
	fmt.Printf("Temperature: %.1f°C\n", temperature)

	team := "Team A"
	score := 85
	isWinner := true

	fmt.Println("----------------------------------")
	fmt.Println("Team:", team)
	fmt.Println("Score:", score)
	fmt.Println("Is Winner:", isWinner)
}
