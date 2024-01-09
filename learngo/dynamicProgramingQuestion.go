package learngo

import (
	"fmt"
	"math"
)

/*
finding the length of the longest increasing subsequence in an array?

A subsequence is a sequence that can be derived from another sequence by deleting
some or no elements without changing the order of the remaining elements.
In this case, we want to find the length of the longest subsequence in an array
where the elements are in increasing order.
*/

// lengthOfLIS finds the length of the longest increasing subsequence using the Patience Sorting algorithm.
func lengthOfLIS(nums []int) int {
	// Create an array to keep track of potential candidates for the increasing subsequence
	tails := make([]int, 0, len(nums))

	// Iterate through each element in the input array
	for _, num := range nums {
		// Use binary search to find the insertion position for 'num' in 'tails'
		i := BinarySearch(tails, num)

		// If 'num' is greater than the last element in 'tails',
		// append it to 'tails' to extend the subsequence.
		if i == len(tails) {
			tails = append(tails, num)
		} else {
			// If 'num' can replace an element in 'tails',
			// update that element with 'num' to potentially form a longer subsequence.
			tails[i] = num
		}
		fmt.Println(tails)
	}

	// The length of 'tails' represents the length of the longest increasing subsequence.
	return len(tails)
}

// binarySearch finds the insertion position of 'target' in 'arr' using binary search.
func BinarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid // Found 'target' at index 'mid'
		} else if arr[mid] < target {
			low = mid + 1 // Search in the right half of 'arr'
		} else {
			high = mid - 1 // Search in the left half of 'arr'
		}
	}

	// 'target' was not found, return the insertion position
	return low
}

func callLengthOfLIS() {
	arr := []int{10, 22, 9, 33, 21, 50, 41, 60} // Change this array for different sequences

	// Calculate the length of the longest increasing subsequence
	result := lengthOfLIS(arr)

	// Display the result
	fmt.Printf("The length of the longest increasing subsequence is: %d\n", result)
}

/*
classic dynamic programming problem: the "0/1 Knapsack Problem."

The 0/1 Knapsack Problem is a combinatorial optimization problem
where you are given a set of items, each with a weight and a value,
and you want to determine the number of each item to include in a collection so that
the total weight is less than or equal to a given limit, and the total value is maximized.
*/
// knapsack1D solves the 0/1 Knapsack Problem using a 1D DP array.
func knapsack1D(weights, values []int, capacity int) int {
	cap := make([]int, capacity+1)
	for idx, w := range weights {
		// it its about adjusting capacity array for every new element in weights array.
		for c := capacity; c >= w; c-- {
			cap[c] = max(
				cap[c],               // if we take the previouly calculated value when the current element was not seen by our code.
				values[idx]+cap[c-w], // if we take the new value which is now seen by our code now and with remaining capacity value which is see earlier.
			)
		}
	}
	return cap[capacity]
}

// max returns the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func callKnapsack() {
	weights := []int{2, 3, 4, 5} // Change weights of items
	values := []int{3, 4, 5, 6}  // Change values of items
	capacity := 5                // Change knapsack capacity

	maxValue := knapsack1D(weights, values, capacity)
	fmt.Printf("The maximum value that can be obtained: %d\n", maxValue)
}

/*

"Longest Common Subsequence (LCS)" problem?

The Longest Common Subsequence problem is about
finding the longest subsequence present in given sequences (strings),
which are common to both the sequences and in the same relative order
but not necessarily contiguous.
It's a classic problem in computer science often used in bioinformatics,
text comparison, and more.
*/

// lcsLength1D calculates the length of the Longest Common Subsequence (LCS) of two strings using a 1D DP array.
func lcsLength1D(text1, text2 string) int {
	m, n := len(text1), len(text2)

	// Create a 1D DP array to store the lengths of LCS for different substrings
	dp := make([]int, n+1)

	// Temporarily store the previous value (top-left) before it gets updated
	prev := 0

	// Fill the DP array using a 1D approach
	for i := 1; i <= m; i++ {

		prev = 0 // Reset the value before each row iteration
		for j := 1; j <= n; j++ {
			fmt.Println(dp)
			temp := dp[j] // Store the current value before it gets updated
			if text1[i-1] == text2[j-1] {
				dp[j] = prev + 1 // Update the DP array with the LCS length
			} else {
				dp[j] = max(dp[j], dp[j-1]) // Take the maximum LCS length till the previous characters
			}
			prev = temp // Update the previous value for the next iteration
			fmt.Println(dp)
		}

		fmt.Println()
	}

	return dp[n]
}

// max returns the maximum of two integers.
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

func callLcsLength1D() {
	text1 := "abcde" // Change the first string
	text2 := "ace"   // Change the second string

	length := lcsLength1D(text1, text2)
	fmt.Printf("The length of the Longest Common Subsequence is: %d\n", length)
}

/*
 "Travelling Salesman Problem (TSP)."
 This problem involves finding the shortest possible route that visits each city exactly once
 and returns to the origin city.
*/

const inf = math.MaxInt64

func tsp1D(graph [][]int) int {
	n := len(graph)
	allVisited := (1 << uint(n)) - 1 // Represents all cities visited

	// Create DP array for memoization
	dp := make([]int, 1<<uint(n))
	for i := range dp {
		dp[i] = inf
	}

	// Base case: Going from a city to itself costs 0
	dp[1] = 0

	// Iterate through all possible subsets of cities
	for mask := 1; mask < allVisited; mask++ {
		for u := 0; u < n; u++ {
			if mask&(1<<uint(u)) == 0 {
				continue // Skip if city 'u' is not included in the subset
			}
			for v := 0; v < n; v++ {
				if mask&(1<<uint(v)) != 0 || u == v {
					continue // Skip if city 'v' is already visited or same as 'u'
				}
				nextMask := mask | (1 << uint(v))
				dp[nextMask] = min(dp[nextMask], dp[mask]+graph[u][v])
			}
		}
	}

	// Find the minimum cost to return to the starting city
	minCost := inf
	for i := 0; i < n; i++ {
		minCost = min(minCost, dp[(1<<uint(n))-1-(1<<uint(i))]+graph[i][0])
	}

	return minCost
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func calltsp1D() {
	// Example graph representing the distances between cities
	graph := [][]int{
		{0, 10, 15, 20},
		{10, 0, 35, 25},
		{15, 35, 0, 30},
		{20, 25, 30, 0},
	}

	result := tsp1D(graph)
	fmt.Printf("Minimum cost to visit all cities in TSP using 1D DP: %d\n", result)
}

/*
Longest Increasing Subsequence (LIS): Find the length of the longest subsequence in an array that is strictly increasing.

Knapsack Problems:

0/1 Knapsack: Given weights and values of items, determine the maximum value that can be put into a knapsack of a given capacity.
Unbounded Knapsack: Similar to 0/1 Knapsack, but with unlimited quantities of items.
Fibonacci Sequence: Find the Nth number in the Fibonacci sequence using various approaches like recursive, memoization, or bottom-up dynamic programming.

Coin Change: Determine the number of ways to make a given amount using a given set of coins.

Matrix Chain Multiplication: Given a sequence of matrices, determine the most efficient way to multiply these matrices together.

Edit Distance: Find the minimum number of operations required to convert one string into another, including insertions, deletions, and substitutions.

Maximum Subarray Sum: Find the contiguous subarray within an array that has the largest sum.

Longest Common Subsequence (LCS): Find the longest subsequence present in given sequences (strings), which are common to both the sequences and in the same relative order but not necessarily contiguous.

Partition Equal Subset Sum: Determine if a given set can be partitioned into two subsets such that the sum of elements in both subsets is equal.

Counting Problems: Various counting problems like counting paths in grids, counting valid parenthesis combinations, or counting distinct subsequences.

Each of these problems presents its own challenges and can be solved using various dynamic programming techniques like memoization, tabulation, bitmasking, and more. Understanding these problems and their solutions can greatly improve problem-solving skills and algorithmic thinking.
*/
