package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	s := []int{3, 5, 6, 1, 3, 4, 6}
	res := quickSort(s)
	assert.Equal(t, s, []int{})
}
