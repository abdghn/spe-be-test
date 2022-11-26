package main

func parityOutlier(input int) interface{} {
	var evens []int
	var odds []int

	for i := 0; i < input; i++ {
		if i%2 > 0 {
			odds = append(odds, i)
		}

		if i%2 == 0 {
			evens = append(evens, i)
		}

		if len(evens) > len(odds) {
			return evens[0]
		} else if len(evens) < len(odds) {
			return odds[0]
		}
	}

	return false
}
