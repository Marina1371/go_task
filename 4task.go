package main

import "fmt"

func main() {
	var number, first_digit, second_digit, third_digit, fourth_digit int
	fmt.Println("Enter a four-digit number : ")
	fmt.Scanf("%d\n", &number)
	if 1000 > number || number > 9999 {
		fmt.Println("The date is incorrect. Opetation failed.")
	} else {
		fourth_digit = number % 10
		third_digit = number / 10 % 10
		second_digit = number / 100 % 10
		first_digit = number / 1000
		if fourth_digit == first_digit && third_digit == second_digit {
			fmt.Printf("number %d - palindrome.\n", number)
		} else {
			fmt.Printf("number %d - not a palidrome.\n", number)
		}
	}
}
