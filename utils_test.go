package algods

import (
	"testing"
)

func TestSwap(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	Swap[int](arr, 0, 1)

	if arr[0] != 2 || arr[1] != 1 {
		t.Errorf("expected arr[0]=2, arr[1]=1, actual arr[0]=%d, arr[1]=%d", arr[0], arr[1])
	}
}
