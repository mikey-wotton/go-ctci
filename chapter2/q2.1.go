package chapter2

/*
  Remove Dups: Write code to remove duplicates from an unsorted linked list.
  How would you solve this problem if a temporary buffer is not allowed?

	Hint #9: Have you tried a hash table? You should be able to do this in a
	single pass of the linked list.

	Hint #40: Without extra space, you'll need a (N2) time. Try using two
	pointers, where the second one searches ahead of the first one.
*/

type node struct {
	value int
	next  *node
}

//space complexity: O(N)
//time complexity: O(N)
func dedupeHashMap(head *node) {
	seen := make(map[int]struct{})
	n := head

	seen[n.value] = struct{}{}

	for n.next != nil {
		if _, exists := seen[n.next.value]; !exists {
			seen[n.next.value] = struct{}{}
			n = n.next
		} else {
			n.next = n.next.next
		}
	}
}

//space complexity: O(1)
//time complexity: O(2N)
func dedupeRunner(head *node) {
	current := head
	for current != nil {
		runner := current
		for runner.next != nil {
			if current.value == runner.next.value {
				runner.next = runner.next.next
			} else {
				runner = runner.next
			}
		}
		current = current.next
	}
}
