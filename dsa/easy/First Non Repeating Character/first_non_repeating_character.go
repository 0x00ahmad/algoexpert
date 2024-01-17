package main

func FirstNonRepeatingCharacter(str string) int {
  if len(str) == 0 {
    return -1
  }
  m := make(map[rune]int)
  for _, c := range str {
    m[c]++
  }
  for i, c := range str {
    if m[c] == 1 {
      return i
    }
  }
  return -1
}

