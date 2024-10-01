package main

import (
	"fmt"
	"sync"
)

func main() {
	letter, number := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
			}
		}
	}()
	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		i := 'A'
		for {
			select {
			case <-letter:
				if i >= 'Z' {
					wait.Done()
					return
				}

				fmt.Print(string(i))
				i++
				fmt.Print(string(i))
				i++
				number <- true
			}
		}
	}(&wait)
	number <- true
	wait.Wait()
}

// 也可以分别使用三个 channel 来控制数字，字母以及终止信号的输入.
// package main

// import "fmt"

// func main() {
// 	number := make(chan bool)
// 	letter := make(chan bool)
// 	done := make(chan bool)

// 	go func() {
// 		i := 1
// 		for {
// 			select {
// 			case <-number:
// 				fmt.Print(i)
// 				i++
// 				fmt.Print(i)
// 				i++
// 				letter <- true
// 			}
// 		}
// 	}()

// 	go func() {
// 		j := 'A'
// 		for {
// 			select {
// 			case <-letter:
// 				if j >= 'Z' {
// 					done <- true
// 				} else {
// 					fmt.Print(string(j))
// 					j++
// 					fmt.Print(string(j))
// 					j++
// 					number <- true
// 				}
// 			}
// 		}
// 	}()

// 	number <- true

// 	for {
// 		select {
// 		case <-done:
// 			return
// 		}
// 	}
// }
