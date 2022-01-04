package linkedlist

import (
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}

	pub := New([]int{7, 8})
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "有交叉点",
			args: args{
				headA: NewWith([]int{1, 2}, pub),
				headB: NewWith([]int{3}, pub),
			},
			want: pub,
		},
		{
			name: "没有",
			args: args{
				headA: NewWith([]int{1, 2}, pub),
				headB: NewWith([]int{1}, nil),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetIntersectionNode(tt.args.headA, tt.args.headB); got != tt.want {
				t.Errorf("GetIntersectionNode() = %v, want %v", got.Print(), tt.want.Print())
			}
		})
	}
}
