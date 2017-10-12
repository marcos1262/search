package main

func getInvCount(tab [][]uint8) (invCount int) {
	var n = len(tab)
	for i := 0; i < n*n-1; i++ {
		for j := i + 1; j < n*n; j++ {
			// count pairs(i, j) such that i appears
			// before j, but i > j.
			if tab[j/n][j%n] > 0 &&
				tab[i/n][i%n] > 0 &&
				tab[i/n][i%n] > tab[j/n][j%n] {
				invCount++
			}
		}
	}
	return
}

// find Position of blank from bottom
func findXPosition(tab [][]uint8) int {
	var n = len(tab)
	// start from bottom-right corner of matrix
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if tab[i][j] == 0 {
				return n - i
			}
		}
	}
	return 0
}

// This function returns true if given
// instance of N*N - 1 puzzle is solvable
func isSolvable(tab [][]uint8) bool {
	var n = len(tab)
	// Count inversions in given puzzle
	var invCount = getInvCount(tab)

	// If grid is odd, return true if inversion
	// count is even.
	if n%2 == 1 {
		return invCount%2 == 0
	} else { // grid is even
		var pos = findXPosition(tab)

		if pos%2 == 1 {
			return invCount%2 == 0
		} else {
			return invCount%2 == 1
		}
	}
}
