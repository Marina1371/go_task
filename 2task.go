package main

import "fmt"

func main() {
	var number, first, second, third int

	fmt.Println("Введите трехзначное число:")
	fmt.Scanf("%d\n", &number)

	first = number / 100
	second = (number / 10) % 10
	third = number % 10

	result := third*100 + second*10 + first

	fmt.Printf("%d", result)
}