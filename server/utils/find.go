package utils

func Find(array []uint, value uint) bool {
	inArray := false
	for _, currValue := range array {
		if currValue == uint(value) {
			inArray = true
		}
	}

	return inArray
}
