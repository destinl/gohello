package algorithm

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"sort"
	"testing"
)

type Student1 struct {
	Age int
}

func TestMap(t *testing.T) {
	kv := map[string]Student1{"menglu": {Age: 21}}
	// kv["menglu"].Age = 22 // 不能直接修改map的值，需要先取出再赋值

	s := []Student1{{Age: 21}} //切片可以，切片存储的是对结构体实例的引用
	s[0].Age = 22
	fmt.Println(kv, s)
}

type Student struct {
	Name string
}

func TestEqual(t *testing.T) {
	fmt.Println(&Student{Name: "menglu"} == &Student{Name: "menglu"})
	fmt.Println(Student{Name: "menglu"} == Student{Name: "menglu"})
	fmt.Println([...]string{"1"} == [...]string{"1"})
	// 数组只能与相同纬度长度以及类型的其他数组比较，切片之间不能直接比较
	// fmt.Println([]string{"1"} == []string{"1"})
}

func TestSlice(t *testing.T) {
	str1 := []string{"a", "b", "c"}
	str2 := str1[1:]
	str2[1] = "new"
	fmt.Println(str1)
	str2 = append(str2, "z", "x", "y")
	fmt.Println(str1)
}

const (
	a = iota
	b = iota
)
const (
	name = "menglu"
	c    = iota
	d    = iota
)

func TestConstant(t *testing.T) {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

// 输出是4，不管有没有resp.Body.Close()，这个内存泄漏问题解决了？（同一域名）
func TestMemoryLeak(t *testing.T) {
	num := 12
	for index := 0; index < num; index++ {
		resp, _ := http.Get("https://www.baidu.com")
		_, _ = ioutil.ReadAll(resp.Body)
		resp.Body.Close()
	}
	fmt.Printf("此时goroutine个数= %d\n", runtime.NumGoroutine())
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j] // 修改为最大堆的比较逻辑
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func TestHeap(t *testing.T) {
	hp := hp{}
	heap.Push(&hp, 3)
	heap.Push(&hp, 2)
	heap.Push(&hp, 4)
	heap.Push(&hp, 1)
	fmt.Println(hp.IntSlice)
	fmt.Println(heap.Pop(&hp).(int))
	fmt.Println(hp.IntSlice)
	// h := &hp{}
	// h.Push(3)
	// h.Push(2)
	// h.Push(4)
	// h.Push(1)

	// fmt.Println(h.IntSlice)
	// fmt.Println(h.Pop())
	// fmt.Println(h.IntSlice)
}

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
