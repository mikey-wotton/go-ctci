package chapter3

import (
	"fmt"
)

/*
  Three in One: Describe how you could use a single array to implement three
  stacks.

  Hint #2: A stack is simply a data structure in which the most recently added
  elements are removed first. Can you simulate a single stack using an array?
  Remember that there are many possible solutions, and there are tradeoffs of
  each.

  Hint #12: We could simulate three stacks in an array by just allocating the
  first third of the array to the first stack, the second third to the second
  stack, and the final third to the third stack. One might actually be much
  bigger than the others, though. Can we be more flexible with the divisions?

  Hint #38: If you want to allow for flexible divisions, you can shift stacks
  around. Can you ensure that all available capacity is used?

  Hint #58: Try thinking about the array as circular, such that the end of the
  array "wraps around" to the start of the array.
*/

//NewThreeStackArray Initialises a three array stack with the individual stack capacity provided.
func NewThreeStackArray() *ThreeStackArray {
	return &ThreeStackArray{
		stackCapacity: 33,
		array:         [99]int{},
		sizes:         [3]int{},
	}
}

//ThreeStackArray implements three separate stacks using a single underlying array.
type ThreeStackArray struct {
	stackCapacity int
	array         [99]int
	sizes         [3]int
}

func (t *ThreeStackArray) pop(stack int) (int, error) {
	if !t.isValidStack(stack) {
		return 0, fmt.Errorf("stack %d is not a valid stack", stack)
	}

	if t.isEmpty(stack) {
		return 0, fmt.Errorf("stack %d is empty", stack)
	}

	target := t.array[t.stackIndex(stack)]
	t.array[t.stackIndex(stack)] = 0
	t.sizes[stack]--

	return target, nil
}

func (t *ThreeStackArray) push(stack, target int) error {
	if !t.isValidStack(stack) {
		return fmt.Errorf("stack %d is not a valid stack", stack)
	}

	if t.isFull(stack) {
		return fmt.Errorf("stack %d is full", stack)
	}

	t.sizes[stack]++
	t.array[t.stackIndex(stack)] = target

	return nil
}

func (t *ThreeStackArray) peek(stack int) (int, error) {
	if t.isEmpty(stack) {
		return 0, fmt.Errorf("stack %d is empty", stack)
	}

	return t.array[t.stackIndex(stack)], nil
}

//Aux/helper functions

func (t *ThreeStackArray) stackIndex(stackNum int) int {
	offset := t.stackCapacity * stackNum
	size := t.sizes[stackNum]

	return offset + size - 1
}

func (t *ThreeStackArray) isFull(stack int) bool {
	return t.sizes[stack] == t.stackCapacity
}

func (t *ThreeStackArray) isEmpty(stack int) bool {
	return t.sizes[stack] == 0
}

func (t *ThreeStackArray) isValidStack(stack int) bool {
	return stack >= 0 && stack < 3
}
