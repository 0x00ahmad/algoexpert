package main

var INT_MIN int = -2147483648

func FindThreeLargestNumbers(array []int) []int {
  var first = INT_MIN
  var second= INT_MIN
  var third = INT_MIN

  for _, value := range array {
    if value > first {
      third = second
      second = first
      first = value
    } else if value > second {
      third = second
      second = value
    } else if value > third {
      third = value
    }
  }

  return []int{third, second, first}

}

