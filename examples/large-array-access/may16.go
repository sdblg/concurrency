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
	fmt.Println("Max value of the slice:", ans)
}

type concurrentSlice struct {
	arr []int
	mtx *sync.Mutex
}

func (cs *concurrentSlice) Get(i int) int {
	cs.mtx.Lock()
	defer cs.mtx.Unlock()
	if 0 <= i && i < len(cs.arr) {
		return cs.arr[i]
	}
	return 0
}

func findMaxConcurrently(workerNumber int, arr []int) int {
	findMax := func(s, e int, cs *concurrentSlice) int {
		maxN := math.MinInt32
		for j := s; j < e; j++ {
			maxN = max(maxN, cs.Get(j))
		}
		return maxN
	}

	worker := func(s, e int, cs *concurrentSlice, wg *sync.WaitGroup, rsp chan int) {
		defer wg.Done()
		rsp <- findMax(s, e, cs)
	}

	n := len(arr)
	rng := n / workerNumber
	if rng == 0 {
		rng++
	}

	wg := &sync.WaitGroup{}
	rsp := make(chan int, workerNumber)
	cs := &concurrentSlice{arr: arr, mtx: &sync.Mutex{}}
	s := 0
	for i := 0; i < workerNumber; i++ {
		wg.Add(1)
		e := s + rng
		if i == workerNumber-1 || n <= e { // last range
			e = n
		}

		go worker(s, e, cs, wg, rsp)
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
