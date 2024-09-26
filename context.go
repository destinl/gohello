package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func task(name string, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Task %s started\n", name)

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Task %s stopped\n", name)
			return
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {

	ctxA, cancelA := context.WithCancel(context.Background())

	ctxB, cancelB := context.WithCancel(ctxA)
	ctxC, cancelC := context.WithCancel(ctxA)
	ctxD, _ := context.WithCancel(ctxA)

	ctxE, _ := context.WithCancel(ctxB)
	ctxF, _ := context.WithCancel(ctxB)

	ctxG, _ := context.WithCancel(ctxC)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go task("A", ctxA, &wg)

	wg.Add(1)
	go task("B", ctxB, &wg)

	wg.Add(1)
	go task("C", ctxC, &wg)

	wg.Add(1)
	go task("D", ctxD, &wg)

	wg.Add(1)
	go task("E", ctxE, &wg)

	wg.Add(1)
	go task("F", ctxF, &wg)

	wg.Add(1)
	go task("G", ctxG, &wg)

	time.Sleep(2 * time.Second)

	cancelB()
	time.Sleep(1 * time.Second)

	cancelC()
	time.Sleep(1 * time.Second)

	cancelA()
	time.Sleep(1 * time.Second)

	wg.Wait()
	fmt.Println("All tasks stopped")
}
