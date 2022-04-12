package main

func CaesarCipherEncryptor(str string, key int) string {
	result := ""
	for _, char := range str {
		if char >= 'a' && char <= 'z' {
			result += string(CaesarCipherEncryptorChar(char, key))
		} else if char >= 'A' && char <= 'Z' {
			result += string(CaesarCipherEncryptorChar(char, key))
		} else {
			result += string(char)
		}
	}

	return result
}

func CaesarCipherEncryptorChar(char rune, key int) rune {
	return rune((int(char)+key-97)%26 + 97)
}
