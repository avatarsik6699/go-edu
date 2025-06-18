package summator

import (
	"slices"
	"testing"
)

func TestSummator(t *testing.T) {
	t.Run("should return sum of all passed slices", func(t *testing.T) {
		want := []int{5, 10, 15}
		got := SumAll([]int{1, 4}, []int{5, 2, 3}, []int{15, 0})

		if !slices.Equal(want, got) {
			t.Errorf("want -> %+v, got -> %+v", want, got)
		}
	})
}
