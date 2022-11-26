package main

func isNarc(input int) bool {
	digits := digitize(input)
	sum := 0

	for i := 0; i < len(digits); i++ {
		sum += pow(digits[i], len(digits))
	}

	return sum == input
}

func pow(i, k int) int {
	result := 1

	for ; k > 0; k-- {
		result *= i
	}

	return result
}

func digitize(input int) []int {
	if input <= 0 {
		return nil
	}

	result := make([]int, 0)

	for input > 0 {
		result = append(result, input%10)

		input /= 10
	}

	return result
}
