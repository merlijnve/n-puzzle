package main

import (
	"bufio"
	"errors"
	"os"
	"sort"
	"strconv"
	"strings"
	"fmt"
	"math/rand"
	"time"
)

// Check if line is empty,
// if not, remove comments and return line
func clean_line(line string) string {
	if line == "" {
		check(errors.New("Empty line in file"))
	}
	if line[0] == '#' {
		return ""
	}
	withoutComment := strings.Split(line, "#")[0]
	return withoutComment
}

// Checks if puzzle size is correct,
//
// parses all elements in puzzle,
//
// returns array of numbers
func create_number_slice(lines []string, n int) []int {
	numbers := []int{}

	for i := range lines { // for all lines of puzzle
		elements := strings.Fields(lines[i]) // split on space
		if len(elements) != n {
			check(errors.New("Puzzle size not correct"))
		}
		for i := range elements { // for all elements in line
			number, err := strconv.Atoi(elements[i]) // convert to int
			check(err)
			numbers = append(numbers, number) // add to numbers splice
		}
	}
	return numbers
}

func goal_map_to_array(goal []point, n int) []int {
	var numbers = make([]int, n*n)

	for i := range goal {
		numbers[goal[i].x+(n*goal[i].y)] = i
	}
	return numbers
}

func calc_inversion(numbers []int, n int) int {
	inversions := 0

	for i := range numbers {
		for j := range numbers {
			if i+j < len(numbers) &&
				numbers[i] != 0 &&
				numbers[i+j] != 0 &&
				numbers[i] > numbers[j+i] {
				inversions++
			}
		}
	}
	return inversions
}

func check_inversion(numbers []int, n int, goal []point) {
	goal_inversions := calc_inversion(goal_map_to_array(goal, n), n)
	start_inversions := calc_inversion(numbers, n)
	pos_of_blank := goal[0].x + (n * goal[0].y)

	if n%2 == 0 {
		goal_inversions += pos_of_blank
		for i := range numbers {
			if numbers[i] == 0 {
				start_inversions += i
			}
		}
	}
	if start_inversions%2 != goal_inversions%2 {
		check(errors.New("Unsolvable"))
	}
}

func check_numbers(numbers []int, n int) {
	copiedNumbers := make([]int, len(numbers))
	copy(copiedNumbers, numbers)

	sort.Slice(copiedNumbers, func(i, j int) bool {
		return copiedNumbers[i] < copiedNumbers[j]
	})
	for i := range copiedNumbers {
		if copiedNumbers[i] != i {
			check(errors.New("Incorrect numbers (duplicate or skipping numbers)"))
		}
	}

}

func create_random_puzzle(filename string) ([]int, int, []point) {
	n, err := strconv.Atoi(filename[1:])
	
	if err != nil {
		check(errors.New("Incorrect size for randomized puzzle"))
	}
	tmp := make([]int, n)

	goal := create_goal_map(n)
	numbers := goal_map_to_array(goal, n)
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Generated puzzle with n =", n)
	for i:=0 ; i < 1000; i++ {
		randnum := rand.Int() % 4 
		switch randnum {
		case 0:
			tmp = up(numbers, n)
		case 1:
			tmp = down(numbers, n)
		case 2:
			tmp = left(numbers, n)
		case 3:
			tmp = right(numbers, n)
		}
		if tmp != nil {
			numbers = tmp
		}
	}
	print_puzzle(numbers, n)
	return numbers, n, goal
}

func parser(filename string) ([]int, int, []point) {
	lines := []string{}
	n := 0
	numbers := make([]int, 1)

	if filename[0] != '+' {
		file, err := os.Open(filename) // open file
		check(err)
		scanner := bufio.NewScanner(file) // create new scanner
		for scanner.Scan() {              // scan line until done
			str := clean_line(scanner.Text())
			if str != "" {
				lines = append(lines, str) // add line to 'lines' slice
			}
		}
		n, err := strconv.Atoi(lines[0]) // get n (width/height of puzzle)
		check(err)
		numbers = create_number_slice(lines[1:], n) // create slice with all numbers
		check_numbers(numbers, n)
		goal := create_goal_map(n)
		check_inversion(numbers, n, goal)
		return numbers, n, goal
	}
	numbers, n, goal := create_random_puzzle(filename)
	return numbers, n, goal
}
