package lg_problems

import "fmt"

var MinimumBribes_arr = [][]int32{
	{2, 1, 5, 3, 4},
	{2, 5, 1, 3, 4},
}

func MinimumBribes(q []int32) {
	var bribes, tmp int32
	/*
	   if len(q) == 1 {
	       fmt.Println(0)
	       return
	   }

	   if len(q) == 2 {
	       if q[0] > q[1] {
	           fmt.Println(1)
	           return
	       } else {
	           fmt.Println(0)
	           return
	       }
	   }

	   if len(q) == 3 {
	       if q[0] > q[2] {
	           fmt.Println(2)
	           return
	       }

	       if q[0] > q[1] || q[1] > q[2] {
	           fmt.Println(1)
	           return
	       }
	   }
	*/
	// this is a sorting problem
	// case 1: number is in it's place
	// case 2:
	//     a) have bribed 1
	//     b) have bribed 2
	//     c) too chaotic
	// q[0]=1 q[1]=2 q[2]=3 q[3]=5 q[4]=4
	for i := int32(len(q) - 1); i > int32(0); i-- {
		if q[i] != i+1 {
			if q[i-1] == i+1 {
				// case 2.a
				bribes += 1
				// swap i, i-1
				tmp = q[i]
				q[i] = q[i-1]
				q[i-1] = tmp
			} else if q[i-2] == i+1 {
				// case 2.b
				bribes += 2
				// swap i-2, i-1
				tmp = q[i-2]
				q[i-2] = q[i-1]
				q[i-1] = tmp
				// swap i-1, i
				tmp = q[i-1]
				q[i-1] = q[i]
				q[i] = tmp

			} else {
				fmt.Println("Too chaotic")
				return
			}

		}

	}
	fmt.Println(bribes)
}
