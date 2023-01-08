// This code is from the very begging. Don't run this at home
package lg_misc

import (
	"fmt"
	"os"
)

//v1

func FindMedianSortedArraysV1(nums1 []int, nums2 []int) float64 {
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

//v2

func FindMedianSortedArraysV2(nums1 []int, nums2 []int) float64 {
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
			return singleArrayMedian(nums2)
		}
	} else if len(nums2) == 0 {
		if len(nums1) == 1 {
			return float64(nums1[0])
		} else {
			return singleArrayMedian(nums1)
		}
	}
	// isAlone and isAlone
	if len(nums1) == 1 && len(nums2) == 1 {
		return float64(nums1[0]+nums2[0]) / float64(2)
	}
	// default case
	return singleArrayMedian(jointSortedArray(nums1, nums2))
}

func singleArrayMedian(a []int) float64 {
	if isEvenArray(a) {
		return float64(a[(len(a)/2)-1]+a[len(a)/2]) / float64(2)
	} else {
		return float64(a[len(a)/2])
	}
}

func jointSortedArray(a []int, b []int) []int {
	i, j := 0, 0
	c := make([]int, 0)
	for len(c) < len(a)+len(b) {
		if a[i] < b[j] {
			c = append(c, a[i])
			if i < len(a)-1 {
				i += 1
			} else {
				c = append(c, b[j])
			}
		} else {
			c = append(c, b[j])
			if j < len(b)-1 {
				j += 1
			} else {
				c = append(c, a[i])
			}
		}
		fmt.Println(c, a[i], b[j])
	}
	return c
}

//v3

func FindMedianSortedArraysV3(nums1 []int, nums2 []int) float64 {
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
			return singleArrayMedian(nums2)
		}
	} else if len(nums2) == 0 {
		if len(nums1) == 1 {
			return float64(nums1[0])
		} else {
			return singleArrayMedian(nums1)
		}
	}
	// isAlone and isAlone
	if len(nums1) == 1 && len(nums2) == 1 {
		return float64(nums1[0]+nums2[0]) / float64(2)
	}
	// default case
	return singleArrayMedian(jointSortedArray(nums1, nums2))
}

/* The other seems to be more elegant

func jointSortedArray(a []int, b []int) []int {
	result := []int{}
	tmp := 0
	first := 0
	last := len(a) - 1

	// empty
	if (len(a) == 0) && (len(b) == 0) {
		goto RETURN
	}

	if len(a) == 0 {
		result = b
		goto RETURN
	}
	if len(b) == 0 {
		result = a
		goto RETURN
	}
	// both alone
	if len(a) == 1 && len(b) == 1 {
		if a[0] < b[0] {
			result = append(a, b...)
		} else {
			result = append(b, a...)
		}
		goto RETURN
	}
    //
    if len(a) == 1 {
        result = insert(a, b)
        goto RETURN
    }
    //
    if len(b) == 1 {
        result = insert(b, a)
        goto RETURN
    }


	//intersection equals null
	if a[len(a)-1] < b[0] {
		result = append(a, b...)
		goto RETURN
	}
	//b is at left and a is at right
	if a[0] > b[len(b)-1] {
		result = append(b, a...)
		goto RETURN
	}
	// there is intersection

	// while last of a is greater than first of b, then swap it
	for a[last] > b[first] {
		tmp = a[last]
		a[last] = b[first]
		b[first] = tmp
		last -= 1
		first += 1

	}
	result = append(a, b...)
	goto RETURN

RETURN:
	return result
}*/

/* unused function

func insert(a []int, b []int) []int {
	var result []int
	f := false
	for _, v := range b {
		if a[0] > v {
			result = append(result, v)
		}
		if a[0] < v {
			if !f {
				result = append(result, a[0])
				result = append(result, v)
				f = true
			} else {
				result = append(result, v)
			}
		}
	}
	if !f {
		result = append(result, a[0])
	}
	return result
} */
