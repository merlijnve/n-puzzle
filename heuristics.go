package main

type point struct {
	x int
	y int
}

// Closure used for looking up the coordinate of the point in the
// array with given index
//
// Returns a function that can be initialized in a function
func lookup_point(n int) func(int) point {
	lookup := make([]point, n*n+1)

	for i := range lookup {
		lookup[i] = calculate_coordinate(i, n)
	}
	return func(i int) point {
		return lookup[i]
	}
}

func create_goal_map(n int) []point {
	goal := make([]point, n*n+1)

	index := 1
	top := 0
	bottom := n - 1
	right := n - 1
	left := 0

	dir := 1
	for top <= bottom && left <= right {
		if dir == 1 {
			for i := left; i <= right; i += 1 {
				goal[index] = point{x: i, y: top}
				index++
			}
			top += 1
			dir = 2
		} else if dir == 2 {
			for i := top; i <= bottom; i += 1 {
				goal[index] = point{x: right, y: i}
				index++
			}
			right -= 1
			dir = 3
		} else if dir == 3 {
			for i := right; i >= left; i -= 1 {
				goal[index] = point{x: i, y: bottom}
				index++
			}
			bottom -= 1
			dir = 4
		} else if dir == 4 {
			for i := bottom; i >= top; i -= 1 {
				goal[index] = point{x: left, y: i}
				index++
			}
			left += 1
			dir = 1
		}
	}
	goal[0] = goal[n*n]
	goal = goal[:(n * n)]
	return goal
}

func calculate_coordinate(i int, n int) point {
	p := point{}

	p.x = i % n
	p.y = (i - (i % n)) / n
	return p
}

// Calculate sum of manhattan distances for puzzle with n^2 - 1 tiles
//
// (x2 - x1) + (y2 - y1)
func manhattan_distance(numbers []int, n int) int {
	sum := 0

	goal := create_goal_map(n)
	for i := range numbers {
		if numbers[i] != 0 {
			p_current := calculate_coordinate(i, n)
			p_goal := goal[numbers[i]]
			p_distance := (p_current.x - p_goal.x) + (p_current.y - p_goal.y)
			if p_distance < 0 {
				p_distance *= -1
			}
			sum += p_distance
		}
	}
	return sum

}

// Calculates how many tiles of puzzle are not in the correct position
func hamming_distance(numbers []int, n int) int {
	sum := 0

	goal := create_goal_map(n)
	for i := range numbers {
		p_current := calculate_coordinate(i, n)
		p_goal := goal[numbers[i]]
		if p_goal != p_current {
			sum++
		}
	}
	return sum
}

// Calculates manhattan distance heuristic,
// procedes to iterate snail-wise through the numbers,
// adds 2 to sum every time a linear conflict is found
func md_linear_conflict(numbers []int, n int) int {
	sum := 0

	top := 0
	bottom := n - 1
	right := n - 1
	left := 0

	tmp := 0
	dir := 1
	sum = manhattan_distance(numbers, n)
	for top <= bottom && left <= right {
		if dir == 1 {
			for i := left; i <= right; i += 1 { // left to right -->
				if n*top+i > 0 {
					if tmp > numbers[n*top+i] && numbers[n*top+i] != 0 {
						sum += 2
					}
				}
				tmp = numbers[n*top+i]
			}
			top += 1
			dir = 2
		} else if dir == 2 {
			for i := top; i <= bottom; i += 1 { // top to bottom ˅˅˅˅˅˅
				if n*i+right > 0 {
					if tmp > numbers[n*i+right] && numbers[n*i+right] != 0 {
						sum += 2
					}
				}
				tmp = numbers[n*i+right]
			}
			right -= 1
			dir = 3
		} else if dir == 3 {
			for i := right; i >= left; i -= 1 { // right to left <--
				if n*bottom+i > 0 {
					if tmp > numbers[n*bottom+i] && numbers[n*bottom+i] != 0 {
						sum += 2
					}
				}
				tmp = numbers[n*bottom+i]
			}
			bottom -= 1
			dir = 4
		} else if dir == 4 {
			for i := bottom; i >= top; i -= 1 { // bottom to top ᐱ
				if n*bottom+i > 0 {
					if tmp > numbers[n*i+left] && numbers[n*i+left] != 0 {
						sum += 2
					}
				}
				tmp = numbers[n*i+left]
			}
			left += 1
			dir = 1
		}
	}
	return sum
}
