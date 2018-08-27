package lib

import "strings"

func FizzBuzz(n int) string {
  var output = make([]string, n+1)
  for i:=1; i <= n; i++ {
    if i % 3 == 0 {
      output[i] = "Buzz"
    } else {
      output[i] = "Fizz"
    }
  }

  return strings.TrimSpace(strings.Join(output, " "))
}
