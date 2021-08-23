package main

import (
	"fmt"
	"strconv"
)

type node struct {
	heurValue  int
	toStart    int
	total      int
	state      []int
	move       string
	parent     *node
	leftChild  *node
	rightChild *node
	tried      bool
}

func move_lowest_to_closed(open map[string]node, closed map[string]node, root *node) node {
	var key_of_lowest string

	for root.leftChild != nil {
		root = root.leftChild
		if ()
	}
	// for key, value := range open {
	// 	if key_of_lowest == "" {
	// 		key_of_lowest = key
	// 	}
	// 	if value.total < open[key_of_lowest].total {
	// 		key_of_lowest = key
	// 	}
	// }
	closed[key_of_lowest] = open[key_of_lowest]
	tmp := closed[key_of_lowest]
	delete(open, key_of_lowest)
	return tmp
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
		heurVal := heur(upState, n, goal) * 2
		toStart := calc_to_start(current)
		successors = append(successors, node{heurVal, toStart, heurVal + toStart, upState, "UP", &current, nil, nil, false})
	}

	downState := down(current.state, n)
	if downState != nil {
		heurVal := heur(downState, n, goal) * 2
		toStart := calc_to_start(current)
		successors = append(successors, node{heurVal, toStart, heurVal + toStart, downState, "DOWN", &current, nil, nil, false})
	}

	leftState := left(current.state, n)
	if leftState != nil {
		heurVal := heur(leftState, n, goal) * 2
		toStart := calc_to_start(current)
		successors = append(successors, node{heurVal, toStart, heurVal + toStart, leftState, "LEFT", &current, nil, nil, false})
	}

	rightState := right(current.state, n)
	if rightState != nil {
		heurVal := heur(rightState, n, goal) * 2
		toStart := calc_to_start(current)
		successors = append(successors, node{heurVal, toStart, heurVal + toStart, rightState, "RIGHT", &current, nil, nil, false})
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

func stateToString(numbers []int, n int) string {
	var s string

	for i := 0; i < n*n; i++ {
		s += strconv.Itoa(numbers[i])
	}
	return s
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

func find_and_compare_states(list map[string]node, current node, n int) bool {
	if len(list) == 0 {
		return false
	}
	key := stateToString(current.state, n)
	_, ok := list[key]
	if ok == true && list[key].total < current.total {
		return true
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

func insertIntoTree(root *node, current *node) *node {

	/* If the tree is empty, return a new node */
	if root == nil {
		return current
	}
	/* Otherwise, recur down the tree */
	if current.total < root.total {
		root.leftChild = insertIntoTree(root.leftChild, current)
	} else if current.total > root.total {
		root.rightChild = insertIntoTree(root.rightChild, current)
	}
	return root
}

func print_tree(root *node) {
	if root != nil {
		print_tree(root.leftChild)
		fmt.Printf("%d \n", root.total)
		print_tree(root.rightChild)
	}
}

func astar(numbers []int, n int, heur func([]int, int, []point) int, goal []point) {
	open := make(map[string]node)
	closed := make(map[string]node)
	// open := []node{}
	// closed := []node{}
	node_current := node{}
	time_complexity := 0
	size_complexity := 0
	solution_moves := 0

	node_goal := node{0, 0, 0, get_goal_state(goal, n), "", nil, nil, nil, false}
	heur_value := heur(numbers, n, goal) * 2
	node_start := node{heur_value, 0, heur_value + 0, numbers, "", nil, nil, nil, false}
	open[stateToString(node_start.state, n)] = node_start
	for true {
		node_current = move_lowest_to_closed(open, closed, &node_start)
		time_complexity++
		if stateToString(node_current.state, n) == stateToString(node_goal.state, n) {
			fmt.Println("Reached solved state")
			solution_moves = node_current.toStart
			fmt.Println("Time complexity (nodes selected in open):", time_complexity)
			fmt.Println("Size complexity (nodes saved in memory):", size_complexity)
			fmt.Println("Moves to solution:", solution_moves)
			print_tree(&node_start)
			recursive_print_moves(node_current, solution_moves+1)
			return
		} else {
			successors := get_successors(node_current, n, heur, goal)
			for s := range successors {
				if find_and_compare_states(closed, successors[s], n) == true ||
					find_and_compare_states(open, successors[s], n) == true {
				} else {
					open[stateToString(successors[s].state, n)] = successors[s]
					insertIntoTree(&node_start, &successors[s])
				}
			}
		}
		// fmt.Println(time_complexity, size_complexity)
		if len(open) > size_complexity {
			size_complexity = len(open)
		}
	}
	fmt.Println("Could not solve")
	return
}
