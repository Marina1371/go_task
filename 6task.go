package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	digits    string = "0123456789"
	lowercase string = "abcdefghijklmnopqrstuvwxyz"
	uppercase string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special   string = "_!@#$%^&"
)

func main() {
	var inPassword string
	var count int
	var isDigits, isLowercase, isUppercase, isSpecial, isLength, truePassword bool

	rec(count, isDigits, isLowercase, isUppercase, isSpecial, isLength, truePassword, inPassword)
}

func rec(count int, isDigits, isLowercase, isUppercase, isSpecial, isLength, truePassword bool, inPassword string) {
	if count < 5 && !truePassword {
		count++
		fmt.Println("попытка №:", count)
		fmt.Println("Введите пароль.\nПароль должен содржать цифру, прописные и заглавные ланиские буквы, спец символы \"_!@#$%^&\".\nДлина пароля должна быть от 8(мин) до 15(макс) символов.")
		inPassword = askPassword()
		truePassword = checkPassword(isDigits, isLowercase, isUppercase, isSpecial, isLength, truePassword, inPassword)
		rec(count, isDigits, isLowercase, isUppercase, isSpecial, isLength, truePassword, inPassword)
	} else if count < 5 && truePassword {
		fmt.Println("the password is correct")
	} else {
		fmt.Println("Слишком много попыток")
	}
}

func checkPassword(isDigits, isLowercase, isUppercase, isSpecial, isLength, truePassword bool, inPassword string) bool {
	if len(inPassword) >= 8 && len(inPassword) <= 15 {
		isLength = true
		for i := 0; i < len(inPassword); i++ {

			if !isDigits {
				isDigits = haveDigits(string(inPassword[i]))
			}
			if !isLowercase {
				isLowercase = haveLowercase(string(inPassword[i]))
			}
			if !isUppercase {
				isUppercase = haveUppercase(string(inPassword[i]))
			}
			if !isSpecial {
				isSpecial = haveSpecial(string(inPassword[i]))
			}

		}
		if !isDigits {
			fmt.Println("the password does not number")
		}
		if !isLowercase {
			fmt.Println("Пароль не содержит маленьких букв")
		}
		if !isUppercase {
			fmt.Println("Пароль не содержит больших букв")
		}
		if !isSpecial {
			fmt.Println("Пароль не содержит спец символов")
		}
	} else {
		isLength = false
		fmt.Println("password is incorrect")
	}

	if isDigits && isLowercase && isUppercase && isSpecial && isLength {
		return true
	} else {
		return false
	}
}

func askPassword() string {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	data := strings.TrimSpace(line)
	return data
}

func haveDigits(v string) bool {
	if strings.Count(digits, v) != 0 {
		return true
	} else {
		return false
	}
}

func haveLowercase(v string) bool {
	if strings.Count(lowercase, v) != 0 {
		return true
	} else {
		return false
	}
}

func haveUppercase(v string) bool {
	if strings.Count(uppercase, v) != 0 {
		return true
	} else {
		return false
	}
}

func haveSpecial(v string) bool {
	if strings.Count(special, v) != 0 {
		return true
	} else {
		return false
	}
}
