package main

import (
	"fmt"
)

// func reverseNumber(num int) int {

//    res := 0
//    for num > 0 {
//       remainder := num % 10
//       res = (res * 10) + remainder
//       num /= 10
//    }
//    return res
// }

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	var s string

	fmt.Println("Введите значение:")

	fmt.Scanf("%s ", &s)
	fmt.Println(Reverse(s))
}
