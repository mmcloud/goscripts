package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	want := 45
	got := Sum(numbers)

	if got != want {
		t.Errorf("wanted %d but got %d", want, got)
	}

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{4, 5, 6}, []int{6, 3, 5})
	want := []int{15, 14}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

}
