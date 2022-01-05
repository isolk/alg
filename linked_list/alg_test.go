package linkedlist

import (
	"testing"
)

func TestHasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "has",
			args: args{head: NewCycle([]int{1, 2, 3})},
			want: true,
		},
		{
			name: "not",
			args: args{head: New([]int{1, 2, 3})},
			want: false,
		},
		{
			name: "nil",
			args: args{head: nil},
			want: false,
		},
		{
			name: "1 false",
			args: args{head: New([]int{1})},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasCycle(tt.args.head); got != tt.want {
				t.Errorf("HasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseBetween(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1,1,1",
			args: args{
				head:  New([]int{1}),
				left:  1,
				right: 1,
			},
			want: New([]int{1}),
		},
		{
			name: "12,1,1",
			args: args{
				head:  New([]int{1, 2}),
				left:  1,
				right: 1,
			},
			want: New([]int{1, 2}),
		},
		{
			name: "12,1,2",
			args: args{
				head:  New([]int{1, 2}),
				left:  1,
				right: 2,
			},
			want: New([]int{2, 1}),
		},
		{
			name: "123,1,3",
			args: args{
				head:  New([]int{1, 2, 3}),
				left:  1,
				right: 3,
			},
			want: New([]int{3, 2, 1}),
		},
		{
			name: "123,2,3",
			args: args{
				head:  New([]int{1, 2, 3}),
				left:  2,
				right: 3,
			},
			want: New([]int{1, 3, 2}),
		},
		{
			name: "12345,3,4",
			args: args{
				head:  New([]int{1, 2, 3, 4, 5}),
				left:  3,
				right: 4,
			},
			want: New([]int{1, 2, 4, 3, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseBetween(tt.args.head, tt.args.left, tt.args.right); !got.Eq(tt.want) {
				t.Errorf("ReverseBetween() = %v, want %v", got.Print(), tt.want.Print())
			}
		})
	}
}

func TestSortList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
			args: args{head: New([]int{1})},
			want: New([]int{1}),
		},
		{
			name: "12",
			args: args{head: New([]int{1, 2})},
			want: New([]int{1, 2}),
		},
		{
			name: "21",
			args: args{head: New([]int{2, 1})},
			want: New([]int{1, 2}),
		},
		{
			name: "123",
			args: args{head: New([]int{1, 2, 3})},
			want: New([]int{1, 2, 3}),
		},
		{
			name: "132",
			args: args{head: New([]int{1, 3, 2})},
			want: New([]int{1, 2, 3}),
		},
		{
			name: "213",
			args: args{head: New([]int{2, 1, 3})},
			want: New([]int{1, 2, 3}),
		},
		{
			name: "231",
			args: args{head: New([]int{2, 3, 1})},
			want: New([]int{1, 2, 3}),
		},
		{
			name: "312",
			args: args{head: New([]int{3, 1, 2})},
			want: New([]int{1, 2, 3}),
		},
		{
			name: "321",
			args: args{head: New([]int{3, 2, 1})},
			want: New([]int{1, 2, 3}),
		},
		{
			name: "1 -1 -1 1",
			args: args{head: New([]int{1, -1, -1, 1})},
			want: New([]int{-1, -1, 1, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortList(tt.args.head); !got.Eq(tt.want) {
				t.Errorf("SortList() = %v, want %v", got.Print(), tt.want.Print())
			}
		})
	}
}
