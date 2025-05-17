package largearrayaccess

import (
	"math"
	"math/rand"
	"sync"
	"testing"
)

func TestFindMaxConcurrently(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{
			name: "all positive numbers",
			arr:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: 10,
		},
		{
			name: "contains negative numbers",
			arr:  []int{-10, -5, -1, -20},
			want: -1,
		},
		{
			name: "mixed positive and negative",
			arr:  []int{3, -2, 7, 0, -8, 5},
			want: 7,
		},
		{
			name: "single element",
			arr:  []int{42},
			want: 42,
		},
		{
			name: "empty array",
			arr:  nil,
			want: math.MinInt32,
		},
		{
			name: "all same elements",
			arr:  []int{5, 5, 5, 5, 5},
			want: 5,
		},
		{
			name: "large numbers",
			arr:  []int{math.MinInt32, 0, math.MaxInt32},
			want: math.MaxInt32,
		},
		{
			name: "all zeros",
			arr:  []int{0, 0, 0, 0, 0},
			want: 0,
		},
		{
			name: "alternating min and max",
			arr:  []int{math.MinInt32, math.MaxInt32, math.MinInt32, math.MaxInt32},
			want: math.MaxInt32,
		},
		{
			name: "descending order",
			arr:  []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			want: 10,
		},
		{
			name: "ascending order",
			arr:  []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1},
			want: -1,
		},
		{
			name: "two elements, negative and positive",
			arr:  []int{-100, 100},
			want: 100,
		},
		{
			name: "very large array",
			arr: func() []int {
				a := make([]int, 10000)
				for i := range a {
					a[i] = i - 5000
				}
				return a
			}(),
			want: 4999,
		},
		{
			name: "max at the beginning",
			arr:  []int{100, 1, 2, 3, 4, 5},
			want: 100,
		},
		{
			name: "max at the end",
			arr:  []int{1, 2, 3, 4, 5, 100},
			want: 100,
		},
		{
			name: "max in the middle",
			arr:  []int{1, 2, 100, 3, 4, 5},
			want: 100,
		},
		{
			name: "all negative, same value",
			arr:  []int{-7, -7, -7, -7},
			want: -7,
		},
		{
			name: "single negative element",
			arr:  []int{-42},
			want: -42,
		},
		{
			name: "large random array",
			arr: func() []int {
				rand.Seed(42)
				a := make([]int, 10000)
				maxV := math.MinInt32
				for i := range a {
					a[i] = rand.Intn(20000) - 10000
					if a[i] > maxV {
						maxV = a[i]
					}
				}
				a[5000] = 20000 // ensure a known max
				return a
			}(),
			want: 20000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaxConcurrently(30, tt.arr)
			if got != tt.want {
				t.Errorf("findMaxConcurrently(%v) = %v, want %v", tt.arr, got, tt.want)
			}
		})
	}
}

// Race condition check: run findMaxConcurrently in parallel
func TestFindMaxConcurrently_Race(t *testing.T) {
	arr := make([]int, 10000)
	for i := range arr {
		arr[i] = i
	}
	want := 9999

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				got := findMaxConcurrently(30, arr)
				if got != want {
					t.Errorf("findMaxConcurrently(arr) = %v, want %v", got, want)
				}
			}
		}()
	}
	wg.Wait()
}

func BenchmarkFindMaxConcurrently(b *testing.B) {
	arr := make([]int, 1000000)
	for i := range arr {
		arr[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findMaxConcurrently(30, arr)
	}
}
