package sort

import (
	"testing"
)

//var s []int = []int{3, 5, 6, 1, 3, 4, 6, 10, 1000}

var data = []struct {
	seed     []int
	expected []int
}{
	{
		[]int{3, 5, 7, 1, 3, 4, 6, 10, 1000},
		[]int{1, 3, 3, 4, 5, 6, 7, 10, 1000},
	},
}

func TestQuickSort(t *testing.T) {
	for _, v := range data {
		a := quickSort(v.seed)
		assertSlice(t, v.expected, a)
	}
}

func assertSlice(t *testing.T, e, a []int) {
	if len(e) != len(a) {
		t.Errorf(
			"slice size not matcha:\n expected:%d \n actual:%d \n",
			len(e),
			len(a),
		)
	}
	for i, v := range e {
		if v != a[i] {
			t.Errorf(
				"slice value not match(i: %d):\n expected:%v \n actual:%v \n",
				i,
				e,
				a,
			)
		}
	}
}
