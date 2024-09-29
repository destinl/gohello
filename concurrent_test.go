package main

import (
	"bytes"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

// 1
var mu sync.Mutex
var chain string

func TestMutex(t *testing.T) {
	chain = "main"
	A()
	fmt.Println(chain)
}
func A() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> A"
	B()
}
func B() {
	chain = chain + " --> B"
	C()
}
func C() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> C"
}

// 2
var rwmu sync.RWMutex
var count int

func TestRWMutex(t *testing.T) {
	go RWMutexA()
	time.Sleep(2 * time.Second)
	rwmu.Lock()
	defer rwmu.Unlock()
	count++
	fmt.Println(count)
}
func RWMutexA() {
	rwmu.RLock()
	defer rwmu.RUnlock()
	RWMutexB()
}
func RWMutexB() {
	time.Sleep(5 * time.Second)
	RWMutexC()
}
func RWMutexC() {
	rwmu.RLock()
	defer rwmu.RUnlock()
}

// 3
func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		wg.Add(1)
	}()
	wg.Wait()
}

// 4 双检锁
type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if o.done == 1 {
		return
	}
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		o.done = 1
		f()
	}
}

// 5
type MyMutex struct {
	count int
	sync.Mutex
}

func TestMyMutex(t *testing.T) {
	var mu MyMutex
	mu.Lock()
	var mu2 = mu //mu2把加锁状态也复制了，再加锁会死锁
	mu.count++
	mu.Unlock()
	mu2.Lock()
	mu2.count++
	mu2.Unlock()
	fmt.Println(mu.count, mu2.count)
}

// 6 可以编译，运行时内存可能暴涨
// 个人理解，在单核CPU中，内存可能会稳定在256MB，如果是多核可能会暴涨。
var pool = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}

func TestPool(t *testing.T) {
	go func() {
		for {
			processRequest(1 << 28) // 256MiB
		}
	}()
	for i := 0; i < 1000; i++ {
		go func() {
			for {
				processRequest(1 << 10) // 1KiB
			}
		}()
	}
	var stats runtime.MemStats
	for i := 0; ; i++ {
		runtime.ReadMemStats(&stats)
		fmt.Printf("Cycle %d: %dB\n", i, stats.Alloc)
		time.Sleep(time.Second)
		runtime.GC()
	}
}
func processRequest(size int) {
	b := pool.Get().(*bytes.Buffer)
	time.Sleep(500 * time.Millisecond)
	b.Grow(size)
	pool.Put(b)
	time.Sleep(1 * time.Millisecond)
}

// 7 #goroutines: 3很多，最后一个#goroutines: 4
// 因为 ch 未初始化，写和读都会阻塞，之后被第一个协程重新赋值，导致写的ch 都阻塞
func TestChannel(t *testing.T) {
	var ch chan int
	go func() {
		ch = make(chan int, 1)
		ch <- 1
	}()
	go func(ch chan int) {
		time.Sleep(time.Second)
		<-ch
	}(ch)
	c := time.Tick(1 * time.Second)
	for range c {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}

// 8 panic: close of nil channel
// ch 没有被初始化，关闭时会报错。
func TestChannel2(t *testing.T) {
	var ch chan int
	var count int
	go func() {
		ch <- 1
	}()
	go func() {
		count++
		close(ch)
	}()
	<-ch
	fmt.Println(count)
}

// 9 编译报错
// sync.Map 没有 Len 方法。
func TestMap(t *testing.T) {
	var m sync.Map
	m.LoadOrStore("a", 1)
	m.Delete("a")
	fmt.Println(m)
}

// 10 输出1
// happens-before 规则保证了 a 的赋值一定在 c <- 0 之前(发送和接收会相互阻塞等对方)，因此输出 1。
var c = make(chan int)
var a int

func f() {
	a = 1
	<-c
}
func TestHappensBefore(t *testing.T) {

	c <- 0
	print(a)
}
