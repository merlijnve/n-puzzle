package main

import (
	"errors"
	"fmt"
	"os"
)

func print_puzzle(numbers []int, n int) {
	for i := range numbers {
		fmt.Printf("[%v]\t", numbers[i])
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

func usage() {
	check(errors.New("Usage:\n./npuzzle [filename or +3 for random with n=3] [heuristic]\n\nHeuristics to choose from:\n[md] Manhattan Distance\n[hd] Hamming Distance\n[mdlc] Manhattan Distance + Linear Conflict\n[all] All combined (extra but inadmissible)"))
}

func main() {
	argsWithoutProg := os.Args[1:]
	var heurFunc func([]int, int, []point) int

	if len(argsWithoutProg) != 2 {
		usage()
	} else {
		switch argsWithoutProg[1] {
		case "md":
			heurFunc = manhattan_distance
			fmt.Println("Solving with Manhattan Distance heuristic")
		case "hd":
			fmt.Println("Solving with Hamming Distance heuristic")
			heurFunc = hamming_distance
		case "mdlc":
			heurFunc = md_linear_conflict
			fmt.Println("Solving with Manhattan Distance + Linear Conflict heuristic")
		case "all":
			heurFunc = all_combined
			fmt.Println("Solving with all heuristics combined (inadmissible)")
		}
		if heurFunc == nil {
			usage()
		}
	}
	numbers, n, goal := parser(os.Args[1])
	astar(numbers, n, heurFunc, goal)
}
