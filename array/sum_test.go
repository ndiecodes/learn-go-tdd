package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("sum collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		sum := Sum(numbers)
		expected := 10

		if expected != sum {
			t.Errorf("expected %d got %d, given %v", expected, sum, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v got %v", want, got)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSum := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v got %v", want, got)
		}
	}
	t.Run("make the sum of sum slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		checkSum(t, got, want)
	})

	t.Run("make the sum of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSum(t, got, want)
	})

}
