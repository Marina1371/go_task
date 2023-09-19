package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const price = 48

	fmt.Print("Distance:")
	readder := bufio.NewReader(os.Stdin)
	text_d, err := readder.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	distance, err := strconv.Atoi(strings.Replace(text_d, "\n", "", -1))
	// fmt.Println(distance)

	fmt.Print("Rashod:")
	text_c, err := readder.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	consumption, err := strconv.Atoi(strings.Replace(text_c, "\n", "", -1))
	// fmt.Println(consumption)

	fmt.Println("расстояние км", distance)
	fmt.Println("расход:", consumption)
	fmt.Println("стоимость поездки:", (distance * price))
}
