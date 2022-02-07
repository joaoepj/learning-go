# learningo-go
 
## Median of Two Sorted Arrays 
This isn't exactly start at the beggining but having some experience in ~programming~ not following advices I choosed a hard problem known as [Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/) so that are two levels of complexity, the problem and the language.
 
the findMedianSortedArrays function was written to catch some cases that once matched should calculate and return the median of an integer sorted slice. The function starts treating in crescent slices size. First case considers both empty slices, second considers when one of them is empty, third when both has only one element, and so far. THis is done with a chain of ifs and elses.

https://github.com/joaoepj/learningo-go/blob/4503ef3bde8fe0d7b66a80dda903f5229d13bce9/findMedianSortedArrays.go#L14-L35

It worked well when there is no intersection between slices but didn't solved the problem so I wrote a second version. Here I try to consolidate much of small functions in greater ones and apply better slice handling techniques in order to join the slices an deliver a unique sorted slice. This reduces the number of cases to be tested in findMedianSortedArrays. But the sorting function is still quite problematic.

https://github.com/joaoepj/learningo-go/blob/97a734fc53a4c3bc07d199c39b32cd97f2f058c3/findMedianSortedArraysV2.go#L48-L56
