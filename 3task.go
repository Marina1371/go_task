package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var num1, num2, num3 int

	fmt.Println("Введите первое число:")
	fmt.Scanf("%d\n", &num1)

	fmt.Println("Введите второе число:")
	fmt.Scanf("%d\n", &num2)

	fmt.Println("Введите третье число:")
	fmt.Scanf("%d\n", &num3)

	fmt.Print("number:")
	readder := bufio.NewReader(os.Stdin)
	text_d, err := readder.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	text_d = ""
	if num1 < num2 {
		fmt.Println(num1, text_d)
		if num2 < num3 {
			fmt.Println(num2)
			fmt.Println(num3)
		} else {
			fmt.Println(num3)
			fmt.Println(num2)
		}
	} else if num2 < num3 {
		fmt.Println(num2)
		if num1 < num3 {
			fmt.Println(num1)
			fmt.Println(num3)
		} else {
			fmt.Println(num3)
			fmt.Println(num1)
		}
	} else {
		fmt.Println(num3)
		if num2 < num1 {
			fmt.Println(num2)
			fmt.Println(num1)
		} else {
			fmt.Println(num1)
			fmt.Println(num2)
		}
	}
}
