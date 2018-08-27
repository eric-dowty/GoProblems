package test

import (
  "testing"
  "go_problems/lib"
)

func TestFizzBuzz(t *testing.T) {
  if lib.FizzBuzz(1) != "Fizz" {
    t.Error("It should return 'Fizz' for n = 1, got:", lib.FizzBuzz(1))
  }

  if lib.FizzBuzz(2) != "Fizz Fizz" {
    t.Error("It should return 'Fizz Fizz' for n = 2, got:", lib.FizzBuzz(2))
  }

  if lib.FizzBuzz(3) != "Fizz Fizz Buzz" {
    t.Error("It should return 'Fizz Fizz Buzz' for n = 2, got:", lib.FizzBuzz(3))
  }
}
