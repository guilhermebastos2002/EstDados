package main

import "fmt"

func main() {
	vetor := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Elementos do vetor: ")
	for _, num := range vetor {
		fmt.Println(num)
	}
}