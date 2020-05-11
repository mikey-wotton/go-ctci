package chapter1

/*
  Zero Matrix: Write an algorithm such that if an element in an MxN matrix is 0,
  its entire row and column are set to O.
  Hints: # 17, #74, #102

  Hint #17: If you just cleared the rows and columns as you found Os, you'd
  likely wind up clearing the whole matrix. Try finding the cells with zeros
  first before making any changes to the matrix.

  Hint #74: Can you use O(N) additional space instead of O(N2)? What
  information do you really need from the list of cells that are zero?

  Hint #102: You probably need some data storage to maintain a list of the rows
  and columns that need to be zeroed. Can you reduce the additional space usage
  to a (1) by using the matrix itself for data storage?
*/

type zeroLocation struct {
	row int
	col int
}

func zeroMatrix(m [][]int) [][]int {
	//loop over entire matrices, denoting 0s at rowXcolumn
	zeros := make([]zeroLocation, 0, len(m))
	for i := range m {
		for k := range m[i] {
			if m[i][k] == 0 {
				loc := zeroLocation{i, k}
				zeros = append(zeros, loc)
			}
		}
	}

	//for each zero found
	for _, loc := range zeros {
		//replace column with 0s
		for i := 0; i < len(m); i++ {
			m[i][loc.col] = 0
		}
		//replace row with 0s
		for k := 0; k < len(m[loc.row]); k++ {
			m[loc.row][k] = 0
		}
	}

	return m
}
