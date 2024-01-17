package main

import "strconv"

func numerificate(result string, c string, count int) string {
	if count > 9 {
		for count > 9 {
			result += "9" + c
			count -= 9
		}
	}
	result += strconv.Itoa(count) + c
	return result
}

func RunLengthEncoding(str string) string {
	result := ""
	count := 1
	for i := 1; i < len(str); i++ {
		if str[i] == str[i-1] {
			count += 1
		} else {
			result = numerificate(result, string(str[i-1]), count)
		count = 1
		}
	}
	result = numerificate(result, string(str[len(str)-1]), count)
	return result
}
