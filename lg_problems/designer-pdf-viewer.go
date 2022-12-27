package lg_problems

var Alphabet_height = []int32{1,2,3,7,6,4,5,9,8,7,4,7,12,4,23,5,9,12,14,3,2,1,4,6,5,3}

func char2index(s string) []int32 {
	var res []int32
	for _, ch := range s {
		res = append(res, int32(ch-97))
	}
	return res
}

func higher(h []int32, i []int32) int32 {
	var res int32
	for _, v := range i {
		if h[v] > res {
			res = h[v]
		}
	}
	return res
}

func DesignerPdfViewer(h []int32, word string) int32 {
	// Write your code here
	return higher(h, char2index(word)) * int32(len(word))

}
