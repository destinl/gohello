package algorithm

import (
	"fmt"
	"testing"
)

func TestGcdValues(t *testing.T) {
	tests := []struct {
		nums     []int
		queries  []int64
		expected []int
	}{
		{
			nums:     []int{2, 3, 4},
			queries:  []int64{0, 2, 2},
			expected: []int{0, 1, 3},
		},
		// {
		// 	nums:     []int{2, 4, 6},
		// 	queries:  []int64{1, 5, 10},
		// 	expected: []int{0, 3, 3},
		// },
		// {
		// 	nums:     []int{5, 10, 15},
		// 	queries:  []int64{0, 1, 5, 6, 10},
		// 	expected: []int{0, 0, 1, 1, 2},
		// },
		// {
		// 	nums:     []int{8, 12, 16},
		// 	queries:  []int64{10, 15, 20},
		// 	expected: []int{2, 2, 2},
		// },
		// {
		// 	nums:     []int{1},
		// 	queries:  []int64{1},
		// 	expected: []int{0},
		// },
		// {
		// 	nums:     []int{},
		// 	queries:  []int64{1},
		// 	expected: []int{0},
		// },
	}

	for _, tt := range tests {
		t.Run("Testing gcdValues function", func(t *testing.T) {
			got := gcdValues(tt.nums, tt.queries)
			t.Log(got)
			// for i, v := range got {
			// 	if v != tt.expected[i] {
			// 		t.Errorf("For nums: %v, queries: %v; expected %v, got %v", tt.nums, tt.queries, tt.expected, got)
			// 	}
			// }
		})
	}
}
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

func TestConstructGridLayout(t *testing.T) {
	tests := []struct {
		n      int
		edges  [][]int
		result [][]int
	}{
		{
			n:      4,
			edges:  [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}},
			result: [][]int{{0, 1}, {2, 3}},
		},
		{
			n:      5,
			edges:  [][]int{{0, 1}, {1, 3}, {2, 3}, {2, 4}},
			result: [][]int{{0, 1, 2, 3, 4}},
		},
		{
			n:      9,
			edges:  [][]int{{0, 1}, {0, 4}, {0, 5}, {1, 7}, {2, 3}, {2, 4}, {2, 5}, {3, 6}, {4, 6}, {4, 7}, {6, 8}, {7, 8}},
			result: [][]int{{0, 1, 2}},
		},
		// {
		// 	n:      2,
		// 	edges:  [][]int{{0, 1}},
		// 	result: [][]int{{0, 1}},
		// },
		// {
		// 	n:      0,
		// 	edges:  [][]int{},
		// 	result: [][]int{},
		// },
	}

	for _, test := range tests {
		t.Run("Test case", func(t *testing.T) {
			got := constructGridLayout(test.n, test.edges)
			// if !reflect.DeepEqual(got, test.result) {
			// 	t.Errorf("constructGridLayout(%d, %v) = %v; want %v", test.n, test.edges, got, test.result)
			// }
			// got := constructGridLayout(test.n, test.edges)
			t.Log(got)
		})
	}
}
