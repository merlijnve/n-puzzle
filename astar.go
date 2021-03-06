package main

import (
	"fmt"
	"strconv"
	"container/heap"
)

type Node struct {
	heurValue  int
	toStart    int
	total      int
	state      []int
	move       string
	parent     *Node
	Index      int
}

type PriorityQueue []Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].total < pq[j].total
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(Node)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func move_lowest_to_closed(open map[string]bool, closed map[string]Node, root *Node, n int, priorityQueue PriorityQueue) (Node, PriorityQueue) {
	var key_of_lowest string

	lowest := heap.Pop(&priorityQueue).(Node)
	key_of_lowest = stateToString(lowest.state, n)
	closed[key_of_lowest] = lowest
	delete(open, key_of_lowest)
	return lowest, priorityQueue
}

func calc_to_start(current Node) int {
	i := 1
	for current.parent != nil {
		i++
		current = *current.parent
	}
	return i
}

func get_successors(current Node, n int, heur func([]int, int, []point) int, goal []point) []Node {
	successors := make([]Node, 0)

	upState := up(current.state, n)
	if upState != nil {
		heurVal := heur(upState, n, goal)
		toStart := calc_to_start(current)
		successors = append(successors, Node{heurVal, toStart, heurVal + toStart, upState, "UP", &current, 0})
	}

	downState := down(current.state, n)
	if downState != nil {
		heurVal := heur(downState, n, goal)
		toStart := calc_to_start(current)
		successors = append(successors, Node{heurVal, toStart, heurVal + toStart, downState, "DOWN", &current, 0})
	}

	leftState := left(current.state, n)
	if leftState != nil {
		heurVal := heur(leftState, n, goal)
		toStart := calc_to_start(current)
		successors = append(successors, Node{heurVal, toStart, heurVal + toStart, leftState, "LEFT", &current, 0})
	}

	rightState := right(current.state, n)
	if rightState != nil {
		heurVal := heur(rightState, n, goal)
		toStart := calc_to_start(current)
		successors = append(successors, Node{heurVal, toStart, heurVal + toStart, rightState, "RIGHT", &current, 0})
	}

	return successors
}

func stateToString(numbers []int, n int) string {
	var s string

	for i := 0; i < n*n; i++ {
		s = s + strconv.Itoa(numbers[i])
	}
	return s
}

func find_and_compare_states(list map[string]Node, current Node, n int) bool {
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

func recursive_print_moves(n Node, moves int) {
	if n.parent != nil {
		moves--
		recursive_print_moves(*n.parent, moves)
		fmt.Println(moves, n.move)
	}
}

func astar(numbers []int, n int, heur func([]int, int, []point) int, goal []point) {
	priorityQueue := make(PriorityQueue, 1)
	open := make(map[string]bool)
	closed := make(map[string]Node)
	node_current := Node{}
	time_complexity := 0
	size_complexity := 0
	solution_moves := 0

	node_goal := Node{0, 0, 0, goal_map_to_array(goal, n), "", nil, 0}
	heur_value := heur(numbers, n, goal) * 2
	node_start := Node{heur_value, 0, heur_value + 0, numbers, "", nil, 0}
	open[stateToString(node_start.state, n)] = true
	priorityQueue[0] = node_start
	heap.Init(&priorityQueue)
	for true {
		node_current, priorityQueue = move_lowest_to_closed(open, closed, &node_start, n, priorityQueue)
		time_complexity++
		if stateToString(node_current.state, n) == stateToString(node_goal.state, n) {
			solution_moves = node_current.toStart
			fmt.Println("Time complexity (nodes selected in open):", time_complexity)
			fmt.Println("Size complexity (max nodes saved in memory):", size_complexity)
			fmt.Println("Moves to solution:", solution_moves)
			recursive_print_moves(node_current, solution_moves+1)
			return
		} else {
			successors := get_successors(node_current, n, heur, goal)
			for s := range successors {
				if find_and_compare_states(closed, successors[s], n) == false &&
				!open[stateToString(successors[s].state, n)] {
					open[stateToString(successors[s].state, n)] = true
					heap.Push(&priorityQueue, successors[s])
				}
			}
		}
		if priorityQueue.Len() > size_complexity {
			size_complexity = priorityQueue.Len()
		}
	}
	fmt.Println("Could not solve")
	return
}
