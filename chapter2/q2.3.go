package chapter2

/*
  Delete Middle Node: Implement an algorithm to delete a node in the middle
  (i.e., any node but the first and last node, not necessarily the exact middle)
  of a singly linked list, given only access to that node.

  Hint #72: Picture the list 1- > 5 - > 9 - > 12. Removing 9 would make it look
  like 1- > 5 - > 12. You only have access to the 9 node. Can you make it look
  like the correct answer?
*/

//Time Complexity: O(N)
func deleteNode(n *node) {
	for k := n; k != nil && k.next != nil; k = k.next {
		k.value = k.next.value
		if k.next.next == nil {
			k.next = nil
		}
	}
}
