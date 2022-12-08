package main

import (
	"testing"
)

func TestStackEquals(t *testing.T) {
	stack1 := Stack{'A', 'B', 'C'}
	stack2 := Stack{'A', 'B', 'C'}
	stack3 := Stack{'A', 'B', 'D'}
	stack4 := Stack{'A', 'B'}

	if !stack1.Equals(stack2) {
		t.Errorf("%v and %v should be equal", stack1, stack2)
	}
	if stack1.Equals(stack3) {
		t.Errorf("%v and %v should not be equal", stack1, stack3)
	}
	if stack1.Equals(stack4) {
		t.Errorf("%v and %v should not be equal", stack1, stack4)
	}
}

func TestStackPush(t *testing.T) {
	stack := Stack{}

	stack.Push('A')
	if !stack.Equals(Stack{'A'}) {
		t.Fatalf("pushing [A] to the stack resulted in %v, expected [A]", stack)
	}

	stack.Push('B')
	if !stack.Equals(Stack{'A', 'B'}) {
		t.Fatalf("pushing [A] to the stack resulted in %v, expected [A][B]", stack)
	}
}

func TestStackPop(t *testing.T) {
	stack := Stack{'A', 'B', 'C'}
	top := stack.Pop()
	if top != 'C' {
		t.Errorf("pop resulted in %v, expected [C]", top)
	}

	expected := Stack{'A', 'B'}
	if !stack.Equals(expected) {
		t.Errorf("popped stack is: %v, expected: %v", stack, expected)
	}
}

func TestParseMoveOp(t *testing.T) {
	op, err := ParseMoveOp("move 2 from 1 to 3")
	if err != nil {
		t.Fatalf("couldn't parse MoveOp: %v", err)
	}
	if !op.Equals(MoveOp{2, 1, 3}) {
		t.Fatalf("parsing MoveOp resulted in '%v', expected 'move 2 from 1 to 3'", op)
	}
}
