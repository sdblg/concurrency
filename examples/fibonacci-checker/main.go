package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	ex := []int{1, 2, 3, 4}
	fmt.Println(getSubArrHavingSumIsFibNum(ex))
}
/*
Problem:
Given an array, check concurrently if the sum of any subarray is a Fibonacci number.
Requirements:
	•	Use goroutines for checking subarrays.
	•	Use channels to collect matches.
	. 	no memory

Input - [1,2,3,4]
Output - [][]int{{1}, {1, 2}, {2}, {2, 3}, {3}}
*/
func getSubArrHavingSumIsFibNum(arr []int) [][]int {
	wg := &sync.WaitGroup{}
	n := len(arr)
	rsp := make(chan *response)
	for l := 0; l < n; l++ {
		for r := l + 1; r < n; r++ {
			wg.Add(1)
			go worker(arr[l:r], rsp, wg)
		}
	}
	go func() {
		wg.Wait()
		close(rsp)
	}()
	var ans [][]int
	for r := range rsp {
		if r.ok {
			ans = append(ans, r.arr)
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		if ans[i][0] == ans[j][0] {
			return len(ans[i]) < len(ans[j])
		}
		return ans[i][0] < ans[j][0]
	})
	return ans
}

func worker(arr []int, rsp chan *response, wg *sync.WaitGroup) {
	defer wg.Done()
	s := 0
	for _, a := range arr {
		s += a
	}

	a, b := 0, 1
	for b < s {
		a, b = b, a+b
	}

	rsp <- &response{
		ok:  b == s,
		arr: arr,
	}
}

type response struct {
	ok  bool
	arr []int
}
