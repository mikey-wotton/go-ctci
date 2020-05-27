package chapter4

/*
	Route Between Nodes: Given a directed graph, design an algorithm to find out
	whether there is a	route between two nodes.

	Hint #127: Two well-known algorithms can do this. What are the tradeoffs
	between them?

	Hint #127 Answer: Both algorithms will eventually return the correct answer
	but if the route is nearby the starting node then BFS will find the answer
	sooner. Conversely,	if the route is deep down within a tree than DFS could
	return the answer	sooner.
*/

type node struct {
	val   int
	edges []*node
}

//DFS using recurision
func isRouteDFS(root *node, target int) bool {
	if root.val == target {
		return true
	}
	seen := make(map[int]bool)

	seen[root.val] = true

	for _, adj := range root.edges {
		if !seen[adj.val] {
			if isRouteDFS(adj, target) {
				return true
			}
		}
	}

	return false
}

//BFS using channels as a queue
func isRouteBFS(root *node, target int) bool {
	queue := make(chan *node, 50)
	defer close(queue)

	seen := make(map[int]bool)
	seen[root.val] = true
	queue <- root

	for len(queue) > 0 {
		n := <-queue
		if n.val == target {
			return true
		}

		for _, adj := range n.edges {
			if !seen[adj.val] {
				seen[adj.val] = true
				queue <- adj
			}
		}

	}

	return false
}
