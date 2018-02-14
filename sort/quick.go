package sort

import "errors"

func quickSort(s []int) ([]int, error) {
	ss := make([]int, len(s))
	copy(ss, s)

	var counter int
	// one loop
	for {
		if counter > 1000 {
			return nil, errors.New("infinite loop break")
		}
		pivot := ss[len(ss)/2]
		var largeIdx int
		var large int
		for i := 0; i < len(ss)/2; i++ {
			if ss[i] >= pivot {
				largeIdx = i
				large = ss[i]
				break
			}
		}

		var smallIdx int
		var small int
		for i := len(ss) / 2; i < len(ss); i++ {
			if ss[i] < pivot {
				smallIdx = i
				small = ss[i]
			}
		}

		if largeIdx > smallIdx {
			break
		}
		// swap
		s[smallIdx] = large
		s[largeIdx] = small

		counter++
	}

	return ss, nil
}
