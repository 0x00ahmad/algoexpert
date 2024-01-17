package main

var Keypad = map[byte][]byte{
	'0': {'0'},
	'1': {'1'},
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func PhoneNumberMnemonics(phoneNumber string) []string {
	mnemonics := []string{}
	currMnemonic := make([]byte, len(phoneNumber))
	for i := range currMnemonic {
		currMnemonic[i] = '0'
	}
	phoneNumberMnemonicsRecursive(0, phoneNumber, currMnemonic, &mnemonics)

	return mnemonics
}

func phoneNumberMnemonicsRecursive(idx int, phoneNumber string, currentMnemonic []byte, mnemonics *[]string) {
	if idx == len(phoneNumber) {
		mnemonic := string(currentMnemonic)
		*mnemonics = append(*mnemonics, mnemonic)
		return
	}
	digit := phoneNumber[idx]
	letters := Keypad[digit]
	for _, letter := range letters {
		currentMnemonic[idx] = letter
		phoneNumberMnemonicsRecursive(idx+1, phoneNumber, currentMnemonic, mnemonics)
	}
}
