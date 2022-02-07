package main

import (
	"fmt"
	"os"
)

func main() {
	a := []int{3}
	b := []int{-2, -1}
	fmt.Println(findMedianSortedArrays(a, b))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if !isConstrained(nums1) || !isConstrained(nums2) {
		os.Exit(3)
	}
	// isEmpty and isEmpty
	if len(nums1) == 0 && len(nums2) == 0 {
		return float64(0)
	}
	// isEmpty and isAlone
	if len(nums1) == 0 {
		if len(nums2) == 1 {
			return float64(nums2[0])
		} else {
			return oneArrayMedian(nums2)
		}
	} else if len(nums2) == 0 {
		if len(nums1) == 1 {
			return float64(nums1[0])
		} else {
			return oneArrayMedian(nums1)
		}
	}
	// isAlone and isAlone
	if len(nums1) == 1 && len(nums2) == 1 {
		return float64(nums1[0]+nums2[0]) / float64(2)
	}

	// median is inside subset (inner array)
	if isSubset(nums1, nums2) {
		return oneArrayMedian(nums1)
	} else if isSubset(nums2, nums1) {
		return oneArrayMedian(nums2)
	} else { // Intersection nums1 nums2 isEmpty
		// median may be (i) the first element of one array
		// (ii) the last element of the other
		// (iii) average of first and last elements

		if isAtLeft(nums1, nums2) {
			if isOddArray(nums1) && isEvenArray(nums2) {
				if len(nums1) < len(nums2) {
					return float64(nums2[0])
				} else {
					return float64(nums1[len(nums1)-1])
				}

			}

			if isEvenArray(nums1) && isOddArray(nums2) {
				if len(nums1) > len(nums2) {
					return float64(nums1[len(nums1)-1])
				} else {
					return float64(nums2[0])
				}
			}
			// isEven and isEven
			return float64(nums1[len(nums1)-1]+nums2[0]) / float64(2)

		} else { // Num2 isAtLeft
			if isOddArray(nums1) && isEvenArray(nums2) {
				if len(nums1) < len(nums2) {
					return float64(nums2[len(nums2)-1])
				} else {
					return float64(nums2[0])
				}

			}

			if isEvenArray(nums1) && isOddArray(nums2) {
				if len(nums1) > len(nums2) {
					return float64(nums1[0])
				} else {
					return float64(nums2[len(nums2)-1])
				}
			}

			// isEven and isEven
			return float64(nums2[len(nums2)-1]+nums1[0]) / float64(2)

		}
	}
}

func oneArrayMedian(a []int) float64 {
	if isEvenArray(a) {
		return float64(a[(len(a)/2)-1]+a[len(a)/2]) / float64(2)
	} else {
		return float64(a[len(a)/2])
	}
}

func isConstrained(a []int) bool {
	return len(a) == 0 || len(a) > 0 && len(a) <= 1000 &&
		a[0] >= -1000000 && a[len(a)-1] <= 1000000
}

func isSubset(a []int, b []int) bool {
	return b[0] <= a[0] && b[len(b)-1] >= a[len(a)-1]
}

func isAtLeft(a []int, b []int) bool {
	return a[len(a)-1] <= b[0]
}

func isOddArray(a []int) bool {
	return (len(a) % 2) == 1
}

func isEvenArray(a []int) bool {
	return (len(a) % 2) == 0
}
