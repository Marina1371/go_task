package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Cipher(str string) []string {
	cipherMap := map[string]string{
		"a": "00",
		"b": "01",
		"c": "02",
		"d": "03",
		"e": "04",
		"f": "05",
		"g": "06",
		"h": "07",
		"i": "08",
		"j": "09",
		"k": "10",
		"l": "11",
		"m": "12",
		"n": "13",
		"o": "14",
		"p": "15",
		"q": "16",
		"r": "17",
		"s": "18",
		"t": "19",
		"u": "20",
		"v": "21",
		"w": "22",
		"x": "23",
		"y": "24",
		"z": "25",
		" ": "26",
	}
	word := make([]string, 0)
	slice := SplitStrings(str, 1)
	for i := 0; i < len(slice); i++ {
		for key, val := range cipherMap {
			if key == slice[i] {
				word = append(word, val)
			}
		}
	}
	return word
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func SplitStrings(s string, chunkLength int) []string {
	length := len(s)

	chunksCount := length / chunkLength
	if length%chunkLength != 0 {
		chunksCount += 1
	}

	chunks := make([]string, chunksCount)

	for i := range chunks {
		from := i * chunkLength
		to := min(length, from+chunkLength)
		chunks[i] = s[from:to]
	}

	return chunks
}

func main() {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	var s string
	fmt.Println(strings.Join(Cipher(line), s))
}
