package main

import (
	"math/rand"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan int)
	ch2 := make(chan int)
	go createSlice(&wg, ch)
	go sqrNum(&wg, ch, ch2)
	for i := 0; i < 10; i++ {
		num, ok := <-ch2
		if ok {
			println(num)
		}
		println()
	}
	wg.Wait()
}

func createSlice(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	slice := make([]int, 10)
	for i := 0; i < 10; i++ {
		slice[i] = rand.Intn(101)
	}
	for i := 0; i < 10; i++ {
		ch <- slice[i]
	}
	close(ch)
}

func sqrNum(wg *sync.WaitGroup, ch chan int, ch2 chan int) {
	defer wg.Done()

	for i := 0; i < 10; i++ {

		num, ok := <-ch
		if ok {
			ch2 <- num * num
		}

	}
	close(ch2)

}
