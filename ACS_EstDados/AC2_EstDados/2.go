package main

import "fmt"

func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	var input string
	fmt.Println("Digite uma string: ")
	fmt.Scan(&input)

	reversed := reverseString(input)
	fmt.Prinln("Número invertido: ", reversed)
}