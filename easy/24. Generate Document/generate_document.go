package main

func GenerateDocument(characters string, document string) bool {
	if len(document) == 0 {
		return true
	}
	var charMap = make(map[rune]int)
	for _, char := range characters {
		if _, ok := charMap[char]; ok {
			charMap[char]++
		} else {
			charMap[char] = 1
		}
	}
	for _, char := range document {
		if count, ok := charMap[char]; !ok || count == 0 {
			return false
		}
		charMap[char]--
	}

	return true
}
