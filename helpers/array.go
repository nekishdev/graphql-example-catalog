package helpers

func Contains(str string, array []string) bool {
	for _, val := range array {
		if val == str {
			return true
		}
	}
	return false
}
