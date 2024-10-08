package algorithm

import (
	"container/heap"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"sort"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

type W struct {
	b int32
	c int64
}

// uintptr和unsafe.Pointer的区别:
// unsafe.Pointer只是单纯的通用指针类型，用于转换不同类型指针，它不可以参与指针运算；
// 而uintptr是用于指针运算的，GC 不把 uintptr 当指针，也就是说 uintptr 无法持有对象， uintptr 类型的目标会被回收；
// unsafe.Pointer 可以和 普通指针 进行相互转换；
// unsafe.Pointer 可以和 uintptr 进行相互转换。
func TestPointer(t *testing.T) {
	var w *W = new(W)
	//这时w的变量打印出来都是默认值0，0
	fmt.Println(w.b, w.c)

	//现在我们通过指针运算给b变量赋值为10
	b := unsafe.Pointer(uintptr(unsafe.Pointer(w)) + unsafe.Offsetof(w.b))
	*((*int)(b)) = 10
	//此时结果就变成了10，0
	fmt.Println(w.b, w.c)
}

// 在 Go 语言中，nil 的通道无论是发送还是接收操作都会永久阻塞
func TestNilChan(t *testing.T) {

	// var c chan int
	// c <- 1

	var c chan int
	num, ok := <-c
	fmt.Printf("读chan的协程结束, num=%v, ok=%v\n", num, ok)
}

type Ban struct {
	visitIPs map[string]time.Time
	lock     sync.Mutex
}

func NewBan(ctx context.Context) *Ban {
	o := &Ban{visitIPs: make(map[string]time.Time)}
	go func() {
		timeer := time.NewTicker(time.Minute * 1)
		for {
			select {
			case <-timeer.C:
				o.lock.Lock()
				for k, v := range o.visitIPs {
					if time.Now().Sub(v) > time.Minute*1 {
						delete(o.visitIPs, k)
					}
				}
				o.lock.Unlock()
				timeer.Reset(time.Minute * 1)
			case <-ctx.Done():
				return
			}
		}
	}()
	return o
}
func (o *Ban) visit(ip string) bool {
	o.lock.Lock()
	defer o.lock.Unlock()
	if _, ok := o.visitIPs[ip]; ok {
		return true
	}
	o.visitIPs[ip] = time.Now()
	return false
}
func TestMapConcurrency(t *testing.T) {
	success := int64(0)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ban := NewBan(ctx)

	wait := &sync.WaitGroup{}

	wait.Add(1000 * 100)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func(j int) {
				defer wait.Done()
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.visit(ip) {
					// success++
					atomic.AddInt64(&success, 1)
				}
			}(j)
		}
	}
	wait.Wait()

	fmt.Println("success:", success)
}

func TestGoFuncUseTemp(t *testing.T) {
	//for 循环中启动 goroutine 时变量使用问题 : 看起来现在没有这个问题了
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second * 2)
}

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
	a = iota //0
	b = iota //1
)
const (
	name = "menglu" //0
	c    = iota     //1
	d    = iota     //2
)
const (
	x = iota
	_
	y
	z = "pi"
	k
	p = iota
	q
)

func TestConstant(t *testing.T) {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(x, y, z, k, p, q)
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
