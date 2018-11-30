package bubblesort

import "testing"

func TestBubbleSort1(t *testing.T) {
	values := []int{0, 2, 1, 3, 4, 5}
	BubbleSort(values)
	if values[0] != 0 || values[1] != 1 || values[2] != 2 || values[3] != 3 || values[4] != 4 || values[5] != 5 {
		t.Error("BubbleSort() failded. Got", values, "Expected 0 1 2 3 4 5")
	}
}
