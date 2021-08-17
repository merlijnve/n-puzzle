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
	heur := 0
	var heurFunc func([]int, int, []point) int

	if len(argsWithoutProg) != 1 {
		if len(argsWithoutProg) == 0 {
			check(errors.New("Please give a filename as argument"))
		} else {
			check(errors.New("Please give only 1 filename as argument"))
		}
	} else {
		for heur != 1 && heur != 2 && heur != 3 && heur != 4 {
			fmt.Println("Select a heuristic function & press enter:")
			fmt.Println("[1] Manhattan Distance")
			fmt.Println("[2] Hamming Distance")
			fmt.Println("[3] Manhattan Distance + Linear Confict")
			fmt.Scanf("%v", &heur)
			if heur != 1 && heur != 2 && heur != 3 {
				fmt.Println("That's not an option")
			}
		}
	}
	switch heur {
	case 1:
		heurFunc = manhattan_distance
	case 2:
		heurFunc = hamming_distance
	case 3:
		heurFunc = md_linear_conflict
	case 4:
		heurFunc = all_combined
	}
	numbers, n, goal := parser(os.Args[1])

	astar(numbers, n, heurFunc, goal)
	// printpuzzle(numbers, n)
	// fmt.Println("Manhattan: ", manhattan_distance(numbers, n, goal))
	// fmt.Println("Hamming: ", hamming_distance(numbers, n, goal))
	// fmt.Println("Manhattan + Linear Conflict: ", md_linear_conflict(numbers, n, goal))
}
