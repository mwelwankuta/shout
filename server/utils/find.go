package utils

import "strconv"

func Find(array []string, value uint) bool {
	inArray := false
	for _, currValue := range array {
		if currValue == strconv.Itoa(int(value)) {
			inArray = true
		}
	}

	return inArray
}
