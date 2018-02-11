package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var s []int = []int{3, 5, 6, 1, 3, 4, 6}

func TestQuickSort(t *testing.T) {
	res := quickSort(s)
	assert.Equal(t, s, []int{})
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
