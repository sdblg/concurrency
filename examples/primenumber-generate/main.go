package primenumbergenerate

import (
	"fmt"
	"sort"
	"sync"
)

/*
Problem: You need to find all prime numbers less than or equal to a given number N using multiple goroutines.
Requirements:
Use goroutines to check for primality concurrently for different ranges of numbers.
Use channels to collect the prime numbers found by each goroutine.
Use a waitgroup to ensure all goroutines finish before collecting results.

n := 100

[2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97]
*/
func Do() {
	fmt.Println(findPrimeNumbers(10, 100))
}

func findPrimeNumbers(numbersPerGoRoutine, n int) []int {
	goroutineNumber := n / numbersPerGoRoutine
	if goroutineNumber == 0 || n%numbersPerGoRoutine != 0 {
		goroutineNumber++
	}
	resp := make(chan []int, goroutineNumber)
	wg := &sync.WaitGroup{}
	var s, e int = 0, numbersPerGoRoutine
	for i := 0; i < goroutineNumber; i++ {
		wg.Add(1)
		if n < e {
			e = n
		}
		go producer(resp, wg, s, e)
		s += numbersPerGoRoutine
		e += numbersPerGoRoutine
	}

	wg.Wait()
	close(resp)

	var ans [][]int
	for arr := range resp {
		ans = append(ans, arr)
	}

	sort.Slice(ans, func(i, j int) bool {
		return  ans[i] != nil && ans[j] != nil && ans[i][0] < ans[j][0]
	})
	var result []int
	for i := range ans {
		result = append(result, ans[i]...)
	}
	return result
}

func producer(rsp chan []int, wg *sync.WaitGroup, s, e int) {
	defer wg.Done()
	rsp <- prime(s, e)
}

func prime(s, e int) []int {
	var ans []int
	if e < 2 {
		return ans
	}

	if s < 2 && 2 <= e {
		ans = append(ans, 2)
	}
	for i := s; i <= e; i++ {
		if i < 2 || i%2 == 0 {
			continue
		}
		itIsP := true
		for j := 3; j*j <= i; j += 2 {
			if i%j == 0 {
				itIsP = false
				break
			}
		}
		if itIsP {
			ans = append(ans, i)
		}
	}
	return ans
}
