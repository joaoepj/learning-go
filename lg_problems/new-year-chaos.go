package lg_problems

import "fmt"

func MinimumBribes1(q []int32) {
    // Write your code here
    var bribes int32
    //original position is number minus one
    // for each test case
    // for each number
    // if actual position is greater then original position minus two
    // print Too chaotic
    // accumulate the difference and goto net number
    for actualPos := int32(len(q) - 1); actualPos >= 0; actualPos-- {
        if actualPos == q[actualPos] {
            continue
        }
        
        origPos := q[actualPos] - 1
        
        if  origPos - actualPos > 2 {
            fmt.Println("Too chaotic")
            return
        }
        
        bribes += origPos - actualPos
        
    }
    fmt.Println(bribes)
}

/* minimumBribes2 did not achieved the time requirements
Time limit exceeded
Your code did not execute within the time limits. 
Please optimize your code. For more information
on execution time limits,
refer to the environment page */

func minimumBribes2(q []int32) {
    // Write your code here
    bribes := make(map[int32]int32)
    for u := int32(0); u < int32(len(q)); u++ {
        for v := u; v < int32(len(q)); v++ {
            if _, ok := bribes[q[u]]; q[v] < q[u] && !ok {
                bribes[q[u]] = 1
                continue
            }
            
            if _, ok := bribes[q[u]]; q[v] < q[u] && ok {
                bribes[q[u]] = bribes[q[u]] + 1
                if bribes[q[u]] > 2 {
                    fmt.Println("Too chaotic")
                    return
                }
            }
            
            //fmt.Println("map: ", bribes)
        
        }
            
    }
    
    
    var sum int32
    for _,v := range bribes {
        sum += v
    }
    fmt.Println(sum)
}

/* test case 0
2
5
2 1 5 3 4
5
2 5 1 3 4 */

/* test case 1
2
8
5 1 2 3 7 8 6 4
8
1 2 5 3 7 8 6 4 */

/* test case 2
1
8
1 2 5 3 4 7 8 6 */