package main

import (
	"fmt"
)

func editDistanceDP(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)

	// Create a 2D slice to store pre-computed edit distances
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	// Initialize the base cases:
	// If one of the strings is empty, the edit distance is the length of the other string
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}

	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	// Calculate the edit distance for each character
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = min(
				dp[i-1][j]+1, // Insertion
				dp[i][j-1]+1, // Deletion
				dp[i-1][j-1]+(text1[i-1] != text2[j-1]), // Substitution
			)
		}
	}

	return dp[m][n]
}
func main() {

	// Start the session expiration goroutine
	res := editDistanceDP("AGGTAB", "GXTXAYB")

	fmt.Println("Result :", res)

}
