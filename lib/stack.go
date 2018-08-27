package lib

import (
  "errors"
)

type Stack struct {
  Elements []int
}

func StackPeek(stack Stack) (int, error) {
  if StackEmpty(stack) {
    return -1, errors.New("Stack is empty")
  }

  return stack.Elements[0], nil
}

func StackEmpty(stack Stack) bool {
  if len(stack.Elements) == 0 {
    return true
  }

  return false
}

func StackAdd(stack Stack, n int) Stack {
  length := len(stack.Elements) + 1
  new_elements := make([]int, length)
  new_elements[0] = n

  for i := 0; i < len(stack.Elements); i++ {
    new_elements[i+1] = stack.Elements[i]
  }

  return Stack{new_elements}
}

func StackRemove(stack Stack) (int, Stack) {
  length := len(stack.Elements) - 1
  new_elements := make([]int, length)
  removed_val := stack.Elements[0]

  for i := 1; i < len(stack.Elements); i++ {
    new_elements[i-1] = stack.Elements[i]
  }

  return removed_val, Stack{new_elements}
}

func StackPop(stack Stack) (int, Stack, error) {
  if StackEmpty(stack) {
    return -1, Stack{}, errors.New("Stack is empty")
  }
  popped_val, new_stack := StackRemove(stack)
  
  return popped_val, new_stack, nil
}

