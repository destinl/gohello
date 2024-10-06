package algorithm

import (
	"fmt"
	"testing"
)

func TestBinarySort(t *testing.T) {
	TempIntSlice := []int{1, 18, 27}
	// binaryStrings := []string{"110", "10", "1010", "1"}
	// sort.Sort(BinarySlice(binaryStrings))
	fmt.Println(maxGoodNumber(TempIntSlice))
}

func TestRemainingMethods(t *testing.T) {
	n, k := 4, 1
	invocations := [][]int{{1, 2}, {0, 1}, {3, 2}}
	fmt.Println(remainingMethods(n, k, invocations))
}
