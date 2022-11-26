package main

func needleInAHayStack(s []string, str string) int {
	for _, v := range s {
		if v == str {
			return 1
		}
	}

	return 0
}
