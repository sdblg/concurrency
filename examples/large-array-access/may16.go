package largearrayaccess

import (
	"fmt"
	"math"
	"sync"
)

/*

Problem: You are given a large array of integers. Find the maximum value using concurrent search.
Requirements:
Divide the array into smaller sub-arrays, with each goroutine processing a sub-array.
Use channels to send back the maximum value from each goroutine.
Use a waitgroup to wait for all goroutines to finish.
Finally, aggregate the results and return the largest value.

*/

func Do() {
	example1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, -1, 5, 13}
	ans := findMaxConcurrently(30, example1)
	fmt.Println(ans)
}

func findMaxConcurrently(workerNumber int, arr []int) int {
	findMax := func(s, e int, mx *sync.Mutex) int {
		// defer mx.Unlock()
		maxN := math.MinInt32
		// mx.Lock()
		for j := s; j < e; j++ {
			maxN = max(maxN, arr[j])
		}
		return maxN
	}

	worker := func(s, e int, wg *sync.WaitGroup, rsp chan int, mx *sync.Mutex) {
		defer wg.Done()
		rsp <- findMax(s, e, mx)
	}

	n := len(arr)
	rng := n / workerNumber
	if rng == 0 {
		rng++
	}

	wg := &sync.WaitGroup{}
	rsp := make(chan int, workerNumber)
	mx := &sync.Mutex{}
	s := 0
	for i := 0; i < workerNumber; i++ {
		wg.Add(1)
		e := s + rng
		if i == workerNumber-1 || n <= e { // last range
			e = n
		}

		go worker(s, e, wg, rsp, mx)
		s += rng
	}

	go func() {
		wg.Wait()
		close(rsp)
	}()

	ans := math.MinInt32
	for m := range rsp {
		ans = max(ans, m)
	}
	return ans
}
