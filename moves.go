package main

func up(numbers []int, n int) []int {
	lookup := lookup_point(n)

	for i := range numbers {
		if numbers[i] == 0 {
			p_current := lookup(i)
			if p_current.y == n-1 {
				return nil
			} else {
				copiedNumbers := make([]int, len(numbers))
				copy(copiedNumbers, numbers)
				copiedNumbers[i] = copiedNumbers[i+n]
				copiedNumbers[i+n] = 0
				return copiedNumbers
			}
		}
	}
	return nil
}

func down(numbers []int, n int) []int {
	lookup := lookup_point(n)

	for i := range numbers {
		if numbers[i] == 0 {
			p_current := lookup(i)
			if p_current.y == 0 {
				return nil
			} else {
				copiedNumbers := make([]int, len(numbers))
				copy(copiedNumbers, numbers)
				copiedNumbers[i] = copiedNumbers[i-n]
				copiedNumbers[i-n] = 0
				return copiedNumbers
			}
		}
	}
	return nil
}

func left(numbers []int, n int) []int {
	lookup := lookup_point(n)

	for i := range numbers {
		if numbers[i] == 0 {
			p_current := lookup(i)
			if p_current.x == n-1 {
				return nil
			} else {
				copiedNumbers := make([]int, len(numbers))
				copy(copiedNumbers, numbers)
				copiedNumbers[i] = copiedNumbers[i+1]
				copiedNumbers[i+1] = 0
				return copiedNumbers
			}
		}
	}
	return nil
}

func right(numbers []int, n int) []int {
	lookup := lookup_point(n)

	for i := range numbers {
		if numbers[i] == 0 {
			p_current := lookup(i)
			if p_current.x == 0 {
				return nil
			} else {
				copiedNumbers := make([]int, len(numbers))
				copy(copiedNumbers, numbers)
				copiedNumbers[i] = copiedNumbers[i-1]
				copiedNumbers[i-1] = 0
				return copiedNumbers
			}
		}
	}
	return nil
}
