package test

import (
	"testing"
	"go_problems/lib"
	"reflect"
)

func TestStack(t *testing.T) {
	// Stack spec
	elements := []int{1}
	stack := lib.Stack{elements}
	if !reflect.DeepEqual(stack.Elements,elements) {
		t.Error("It should maintiain elements assigned, returned:", stack.Elements)
	}

	// StackPeek spec
	value, err := lib.StackPeek(lib.Stack{[]int{1,2}})
	if value != 1 {
		t.Error("It should return element at top of stack, returned:", err)
	}

	value, err = lib.StackPeek(lib.Stack{[]int{}})
	if err.Error() != "Stack is empty" {
		t.Error("It should return error if stack is empty:", value)
	}

	// StackEmpty spec
	if lib.StackEmpty(lib.Stack{[]int{}}) != true {
		t.Error("It should return true if stack is empty, returned:", lib.StackEmpty(stack))
	}

	if lib.StackEmpty(lib.Stack{[]int{1}}) != false {
		t.Error("It should return false if stack is not empty, returned:", lib.StackEmpty(stack))
	}

	// StackAdd spec
	stack = lib.StackAdd(lib.Stack{[]int{1}}, 5)
	value, err = lib.StackPeek(stack)
	if value != 5 {
		t.Error("It should return new stack with n added to top, returned:", value)
	}

	// StackPop spec
	value, stack, err = lib.StackPop(lib.Stack{[]int{11,2}})
	if value != 11 {
		t.Error("It should return value from top of stack, returned:", value)
	}
	value, err = lib.StackPeek(stack)
	if value != 2 {
		t.Error("It should remove value from top of stack, returned:", value)
	}
}


