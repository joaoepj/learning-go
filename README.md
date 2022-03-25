# learningo-go


## Project Setup
$ mkdir learning-go
$ cd learning-go
$ git init
$ go mod init github.com/joaoepj/learning-go
... write code
How to output the executable to build/bin when issuing `go build`? 


## Median of Two Sorted Arrays 
This isn't exactly start at the beggining but having some experience in ~programming~ not taking advices I choosed a hard problem known as [Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/) so that are two levels of complexity, the problem and the language.
 
the findMedianSortedArrays function was written to catch some cases that once matched should calculate and return the median of an integer sorted slice. The function starts treating in crescent slices size. First case considers both empty slices, second considers when one of them is empty, third when both has only one element, and so far. THis is done with a chain of ifs and elses.

https://github.com/joaoepj/learningo-go/blob/f46c405d423fa202177eef187ac0c104b34b3138/findMedianSortedArrays.go#L18-L28

It worked well when there is no intersection between slices but didn't solved the problem so I wrote a second version. Here I try to consolidate much of small functions in greater ones and apply better slice handling techniques in order to join the slices an deliver a unique sorted slice. This reduces the number of cases to be tested in findMedianSortedArrays. But the sorting function is still quite problematic.

https://github.com/joaoepj/learningo-go/blob/97a734fc53a4c3bc07d199c39b32cd97f2f058c3/findMedianSortedArraysV2.go#L48-L56

In the third version I added some gotos (I'm not proud of it) to the function jointSortedArray and am still fighting to sort the slice

https://github.com/joaoepj/learningo-go/blob/1a9a6355a29ccc081d540d424efa4af2360ba3cf/findMedianSortedArraysV3.go#L50-L57
