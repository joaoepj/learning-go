func main() {
        x := []int{1,2}
    y := []int{3,4}
    fmt.Println(findMedianSortedArrays(x,y))
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
        return float64(a[(len(a)/2)-1] +a[len(a)/2]) / float64(2)
    } else {
        return float64(a[len(a)/2])
    }
}    

func jointSortedArray(a []int, b []int) []int {
    i, j := 0,0
    c := make([]int, 0)
    for len(c) < len(a) + len(b)  {
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
        fmt.Println(c,a[i],b[j])
    }
    return c
}

func isConstrained(a []int) bool {
	return len(a) == 0 || len(a) > 0 && len(a) <= 1000 &&
		a[0] >= -1000000 && a[len(a)-1] <= 1000000
}


func isOddArray(a []int) bool {
	return (len(a) % 2) == 1
}


func isEvenArray(a []int) bool {
	return (len(a) % 2) == 0
}
