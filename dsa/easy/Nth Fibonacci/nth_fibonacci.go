package main

func GetNthFib(n int) int {
  fibies := make([]int, n+1)
  fibies[0] = 0
  fibies[1] = 1
  for i := 2; i <= n; i++ {
    fibies[i] = fibies[i-1] + fibies[i-2]
  }
  return fibies[n-1]
}

