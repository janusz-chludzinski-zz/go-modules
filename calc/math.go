package calc

import "errors"

// return sum of integers
func Add(numbers ...int) (error, int) {
	sum := 0

	if len(numbers) < 2 {
		return errors.New("Provide at least 2 arguments"), sum
	}

	for _, num := range numbers {
		sum += num
	}

	return nil, sum
}
