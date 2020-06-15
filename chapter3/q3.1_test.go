package chapter3

import (
	"testing"
)

func Test_push(t *testing.T) {
	stack := NewThreeStackArray()

	for i := 0; i <= 32; i++ {
		if err := stack.push(0, i); err != nil {
			t.Fatalf("pushing onto stack 1 encountered error - %s", err.Error())
		}
	}

	if err := stack.push(0, 33); err == nil {
		t.Fatal("expected stack full error but received no error")
	}
}

func Test_peekAndPop(t *testing.T) {
	//setup
	stack := NewThreeStackArray()
	for i := 0; i <= 32; i++ {
		if err := stack.push(0, i); err != nil {
			t.Fatalf("pushing onto stack 1 encountered error - %s", err.Error())
		}
	}

	//test
	for i := 32; i >= 0; i-- {
		//test peek
		j, err := stack.peek(0)
		if err != nil {
			t.Fatalf("encountered error peeking on stack but did not expect one - %s", err.Error())
		}
		if i != j {
			t.Fatalf("i{%d} does not equal j{%d} but should", i, j)
		}

		//test pop
		k, err := stack.pop(0)
		if i != k {
			t.Fatalf("i{%d} does not equal k{%d} but should", i, k)
		}
		if err != nil {
			t.Fatalf("received error from stack but did not expect one - %s", err.Error())
		}
	}

	//test empty stack peek causes error
	if _, err := stack.peek(0); err == nil {
		t.Fatal("expected empty stack error but received none")
	}

	//test empty stack pop causes error
	if _, err := stack.pop(0); err == nil {
		t.Fatal("expected empty stack error but received none")
	}
}
