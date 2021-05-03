package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	randomSlice := rand.Perm(1e7)
	t := time.Now()
	sum := add(randomSlice)
	fmt.Printf("sequential ADD, sum: %d \t time taken: %s\n", sum, time.Since(t))

	t1 := time.Now()
	sum1 := addConcurrent(randomSlice)
	fmt.Printf("concurrent ADD, sum: %d \t time taken: %s\n", sum1, time.Since(t1))
}

func add(numbers []int) int64 {
	var sum int64
	for _, v := range numbers {
		sum += int64(v)
	}
	return sum
}

func addConcurrent(numbers []int) int64 {
	numOfCores := runtime.NumCPU()
	runtime.GOMAXPROCS(numOfCores)
	var sum int64
	max := len(numbers)

	sizeofparts := max / numOfCores

	var wg sync.WaitGroup
	for i := 0; i < numOfCores; i++ {
		start := i * sizeofparts
		end := start + sizeofparts
		part := numbers[start:end] // [0:1000] [1000:2000] ..

		wg.Add(1)
		go func(nums []int) {
			defer wg.Done()
			var partsum int64
			for _, n := range nums {
				partsum += int64(n)
			}
			atomic.AddInt64(&sum, partsum)
		}(part)
	}
	wg.Wait()
	return sum
}
