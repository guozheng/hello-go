package array

import "testing"
import "reflect"

func TestSum(t *testing.T) {
	t.Run("sum of 5 nums", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		want := 15
		got := Sum(input)
		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, input)
		}
	})

	t.Run("sum of 3 nums", func(t *testing.T) {
		input := []int{1, 2, 3}
		want := 6
		got := Sum(input)
		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, input)
		}
	})
}

func TestSumEach(t *testing.T) {
	want := []int{2, 9}
	got := SumEach([]int{1, 1}, []int{0, 9})
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v want %v", got, want)
	}
}
