package primenumbergenerate

import (
	"reflect"
	"testing"
)

func Test_findPrimeNumbers(t *testing.T) {
	type args struct {
		numbersPerGoRoutine int
		n                   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "negative",
			args: args{numbersPerGoRoutine: 8, n: -1},
			want: nil,
		},
		{
			name: "two is lower that numbersPerGoRoutine",
			args: args{numbersPerGoRoutine: 8, n: 2},
			want: []int{2},
		},
		{
			name: "eight",
			args: args{numbersPerGoRoutine: 8, n: 8},
			want: []int{2, 3, 5, 7},
		},
		{
			name: "twenty",
			args: args{numbersPerGoRoutine: 8, n: 20},
			want: []int{2, 3, 5, 7, 11, 13, 17, 19},
		},
		{
			name: "hundred",
			args: args{numbersPerGoRoutine: 8, n: 100},
			want: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97},
		},		
		{
			name: "n < numbersPerGoRoutine",
			args: args{numbersPerGoRoutine: 500, n: 100},
			want: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPrimeNumbers(tt.args.numbersPerGoRoutine, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findPrimeNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
