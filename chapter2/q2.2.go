package chapter2

/*
  Return Kth to Last: Implement an algorithm to find the kth to last element of
  a singly linked list.

  Hint #8: What if you knew the linked list size? What is the difference between
  finding the Kth-to last element and finding the Xth element?

  Hint #25: If you don't know the linked list size, can you compute it? How does
  this impact the runtime?

  Hint #47: Look at this graph. Is there any node you can identify that will
  definitely be okay to build first?

  Hint #67: You might find it useful to return multiple values. Some languages
  don't directly support this, but there are workarounds in essentially any
  language. What are some of those workarounds?

  Hint #726: We're probably going to run this algorithm many times. If we did
  more preprocessing, is there a way we could optimize this?
*/

//Time complexity O(N)
func returnKth(head *node, k int) *node {
	node, depth := returnKthRecursive(head, k)
	if depth != k {
		return nil
	}

	return node
}

func returnKthRecursive(n *node, k int) (*node, int) {
	if n.next == nil {
		return n, 0
	}

	node, i := returnKthRecursive(n.next, k)
	if i == k {
		return node, k
	}

	return n, i + 1
}
