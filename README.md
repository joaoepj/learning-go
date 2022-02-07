# learningo-go
 
 This isn't exactly start at the beggining but having some experience in programming I choosed a hard problem at leetcode.com known as Find Median Sorted Arrays so that are two levels of complexity, the problem and the language.
 
the findMedianSortedArrays function was written to catch some cases that once matched should return the of an integer sorted slice.
```go
if len(nums1) == 0 {
		if len(nums2) == 1 {
			return float64(nums2[0])
		} else {
			return oneArrayMedian(nums2)
		}```
