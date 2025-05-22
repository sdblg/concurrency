package main

import (
	"reflect"
	"testing"
)

func Test_getSubArrHavingSumIsFibNum(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "regular",
			args: args{
				arr: []int{1, 2, 3, 4},
			},
			want: [][]int{{1}, {1, 2}, {2}, {2, 3}, {3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSubArrHavingSumIsFibNum(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSubArrHavingSumIsFibNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
