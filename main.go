package main

import (
	"errors"
	"fmt"
	"os"
)

func printpuzzle(numbers []int, n int) {
	for i := range numbers {
		fmt.Printf("%v[%v]\t", calculate_coordinate(i, n), numbers[i])
		if (i+1)%n == 0 {
			fmt.Printf("\n\n")
		}
	}
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) != 1 {
		if len(argsWithoutProg) == 0 {
			check(errors.New("Please give a filename as argument"))
		} else {
			check(errors.New("Please give only 1 filename as argument"))
		}
	}
	numbers, n := parser(os.Args[1])

	numbers = up(numbers, n)
	numbers = down(numbers, n)
	numbers = left(numbers, n)
	numbers = right(numbers, n)

	printpuzzle(numbers, n)
	fmt.Println("Manhattan: ", manhattan_distance(numbers, n))
	fmt.Println("Hamming: ", hamming_distance(numbers, n))
	fmt.Println("Manhattan + Linear Conflict: ", md_linear_conflict(numbers, n))
}
