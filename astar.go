package main

import (
	"fmt"
)

type node struct {
	heurValue int
	toStart   int
	total     int
	state     []int
	move      string
	parent    *node
}

func remove(s []node, i int) []node {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func move_lowest_to_closed(open []node, closed []node) ([]node, []node, node) {
	index_of_lowest := 0

	for i := range open {
		if open[i].total < open[index_of_lowest].total {
			index_of_lowest = i
		}
	}
	closed = append(closed, open[index_of_lowest])
	tmp := open[index_of_lowest]
	open = remove(open, index_of_lowest)
	return open, closed, tmp
}

func calc_to_start(current node) int {
	i := 1
	for current.parent != nil {
		i++
		current = *current.parent
	}
	return i
}

func get_successors(current node, n int, heur func([]int, int, []point) int, goal []point) []node {
	successors := make([]node, 0)

	upState := up(current.state, n)
	if upState != nil {
		heurVal := heur(upState, n, goal)
		toStart := calc_to_start(current)
		successors = append(successors, node{heurVal, toStart, heurVal + toStart, upState, "UP", &current})
	}

	downState := down(current.state, n)
	if downState != nil {
		heurVal := heur(downState, n, goal)
		toStart := calc_to_start(current)
		successors = append(successors, node{heurVal, toStart, heurVal + toStart, downState, "DOWN", &current})
	}

	leftState := left(current.state, n)
	if leftState != nil {
		heurVal := heur(leftState, n, goal)
		toStart := calc_to_start(current)
		successors = append(successors, node{heurVal, toStart, heurVal + toStart, leftState, "LEFT", &current})
	}

	rightState := right(current.state, n)
	if rightState != nil {
		heurVal := heur(rightState, n, goal)
		toStart := calc_to_start(current)
		successors = append(successors, node{heurVal, toStart, heurVal + toStart, rightState, "RIGHT", &current})
	}

	return successors
}

func get_goal_state(goal []point, n int) []int {
	goalState := make([]int, n*n)

	for i := range goal {
		goalState[goal[i].x+goal[i].y*n] = i
	}
	return goalState
}

func compare_states(current_numbers []int, goal_numbers []int) bool {
	if len(current_numbers) == 0 {
		return false
	}
	for i := range goal_numbers {
		if goal_numbers[i] != current_numbers[i] {
			return false
		}
	}
	return true
}

func find_and_compare_states(list []node, current node) bool {
	// fmt.Println("\nSTART FINDING STATES", len(list))
	if len(list) == 0 {
		return false
	}
	for i := range list {
		if compare_states(list[i].state, current.state) == true && list[i].total < current.total {
			return true
		}
	}
	return false
}

func printNode(current node) {
	fmt.Println("Heur:", current.heurValue)
	fmt.Println("ToStart:", current.toStart)
	fmt.Println("Total:", current.total)
	fmt.Println("--------------")
}

func recursive_print_moves(n node, moves int) {
	if n.parent != nil {
		moves--
		recursive_print_moves(*n.parent, moves)
		fmt.Println(moves, n.move)
	}
}

func astar(numbers []int, n int, heur func([]int, int, []point) int, goal []point) {
	open := []node{}
	closed := []node{}
	node_current := node{}
	time_complexity := 0
	size_complexity := 0
	solution_moves := 0

	node_goal := node{0, 0, 0, get_goal_state(goal, n), "", nil}
	heur_value := heur(numbers, n, goal)
	node_start := node{heur_value, 0, heur_value + 0, numbers, "", nil}
	open = append(open, node_start)
	for i := 0; len(open) > 0; i++ {
		// fmt.Println("FINDING LOWEST OF OPEN: ", len(open))
		open, closed, node_current = move_lowest_to_closed(open, closed)
		time_complexity++
		if compare_states(node_current.state, node_goal.state) == true {
			fmt.Println("Reached solved state")
			solution_moves = node_current.toStart
			fmt.Println("Time complexity (nodes selected in open):", time_complexity)
			fmt.Println("Size complexity (nodes saved in memory):", size_complexity)
			fmt.Println("Moves to solution:", solution_moves)
			recursive_print_moves(node_current, solution_moves+1)
			return
		} else {
			successors := get_successors(node_current, n, heur, goal)
			for s := range successors {
				if find_and_compare_states(open, successors[s]) == true ||
					find_and_compare_states(closed, successors[s]) == true {
				} else {
					open = append(open, successors[s])
				}
			}
		}
		if len(open) > size_complexity {
			size_complexity = len(open)
		}
	}
	fmt.Println("Could not solve")
	return
}
