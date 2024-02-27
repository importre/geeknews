package geeknews

import "strconv"

func parseUint(s string) uint {
	result, error := strconv.Atoi(s)

	if error != nil {
		return 0
	}

	if result < 0 {
		return 0
	}

	return uint(result)
}
