package chapter1

/*
  Given an image represented by an NxN matrix, where each pixel in the image
  is 4 bytes, write a method to rotate the image by 90 degrees. can you do this
  in place?

  Hint #51: Try thinking about it layer by layer. Can you rotate a specific
  layer?

  Hint #100: Rotating a specific layer would just mean swapping the values in
  four arrays. If you were asked to swap the values in two arrays, could you do
  this? Can you then extend it to four arrays?
*/

func rotateMatrix(m [][]byte, antiClockwise bool) [][]byte {
	k := len(m)

	for i := 0; i < k/2; i++ {
		for j := i; j < k-i-1; j++ {
			if antiClockwise {
				rotateAntiClockwise(m, i, j, k)
			} else {
				rotateClockwise(m, i, j, k)
			}
		}
	}

	return m
}

func rotateClockwise(m [][]byte, i, j, k int) {
	a := k - i - 1
	b := k - j - 1

	topRight := m[j][a]
	topLeft := m[i][j]
	btmLeft := m[b][i]
	btmRight := m[a][b]

	m[j][a], m[i][j] = topLeft, btmLeft
	m[b][i], m[a][b] = btmRight, topRight
}

func rotateAntiClockwise(m [][]byte, i, j, k int) {
	a := k - i - 1
	b := k - j - 1

	topRight := m[j][a]
	topLeft := m[i][j]
	btmLeft := m[b][i]
	btmRight := m[a][b]

	m[j][a], m[i][j] = btmRight, topRight
	m[b][i], m[a][b] = topLeft, btmLeft
}
