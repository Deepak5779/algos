package binarySearchBased

import "math"

/*
Input_1
n = 8, root = 3
Output = 2

Input_2
N = 1000, root = 6
Output - 3.162 // Accurate up to to 3 decimal points
*/
func findNthRoot(n int, r int) float64 {
	return findNthRootBSUtil(float64(n), float64(r), float64(0), float64(n))
}

func findNthRootBSUtil(n float64, r float64, left float64, right float64) float64 {
	pivot := left + (right-left)/2
	exp := math.Pow(pivot, r)
	diff := n - exp
	if diff >= 0 && diff <= 0.001 {
		return float64(int(pivot*1000) / 1000)
	} else if exp > n {
		return findNthRootBSUtil(n, r, left, pivot)
	} else {
		return findNthRootBSUtil(n, r, pivot, right)
	}
}
