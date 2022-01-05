package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{intervals: [][]int{{1, 4}, {0, 2}, {3, 5}}},
			want: [][]int{{1, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quick_sort(t *testing.T) {
	type args struct {
		res   []int
		begin int
		end   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{res: []int{5, 1, 9, -1}, begin: 0, end: 3},
			want: []int{-1, 1, 5, 9},
		},

		{
			args: args{res: []int{1, 2, 3}, begin: 0, end: 2},
			want: []int{1, 2, 3},
		},
		{
			args: args{res: []int{-6, -1, -2}, begin: 0, end: 2},
			want: []int{-6, -2, -1},
		},
		{
			args: args{res: []int{2, 1}, begin: 0, end: 1},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// quick_sort(tt.args.res, tt.args.begin, tt.args.end)
			// if !reflect.DeepEqual(tt.want, tt.args.res) {
			// 	t.Errorf("Merge() = %v, want %v", tt.args.res, tt.want)
			// }
		})
	}
}

func TestRemoveElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "7",
			args: args{nums: []int{4, 5}, val: 4},
			want: 1,
		},
		{
			name: "7",
			args: args{nums: []int{4, 4, 5}, val: 4},
			want: 1,
		},
		{
			name: "7",
			args: args{nums: []int{4, 4, 5, 5}, val: 4},
			want: 2,
		},
		{
			name: "4",
			args: args{nums: []int{1, 1, 3}, val: 1},
			want: 1,
		},
		{
			name: "1",
			args: args{nums: []int{}, val: 1},
			want: 0,
		},
		{
			name: "2",
			args: args{nums: []int{1}, val: 1},
			want: 0,
		},
		{
			name: "3",
			args: args{nums: []int{1, 2, 3}, val: 1},
			want: 2,
		},

		{
			name: "5",
			args: args{nums: []int{1, 1, 1}, val: 1},
			want: 0,
		},
		{
			name: "6",
			args: args{nums: []int{3, 3, 1}, val: 1},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("RemoveElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
