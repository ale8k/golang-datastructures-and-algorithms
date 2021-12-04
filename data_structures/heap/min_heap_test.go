package heap

import (
	"reflect"
	"testing"
)

func TestMinHeap_ParentIndex(t *testing.T) {
	type fields struct {
		nodes []int
	}
	type args struct {
		childIndex int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "it gets the correct index",
			fields: fields{
				nodes: []int{1, 5, 3, 9, 6, 8},
			},
			args: args{
				childIndex: 5,
			},
			want: 2,
		},
		{
			name: "it gets the correct index from root",
			fields: fields{
				nodes: []int{1, 5, 3, 9, 6, 8},
			},
			args: args{
				childIndex: 0,
			},
			want: 0,
		},
		{
			name: "it gets the correct index one node down",
			fields: fields{
				nodes: []int{1, 5, 3, 9, 6, 8},
			},
			args: args{
				childIndex: 1,
			},
			want: 0,
		},
		{
			name: "it gets the correct index one node down(2)",
			fields: fields{
				nodes: []int{1, 5, 3, 9, 6, 8},
			},
			args: args{
				childIndex: 2,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MinHeap{
				nodes: tt.fields.nodes,
			}
			if got := h.ParentIndex(tt.args.childIndex); got != tt.want {
				t.Errorf("MinHeap.ParentIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinHeap_LeftChildIndex(t *testing.T) {
	type fields struct {
		nodes []int
	}
	type args struct {
		parentNodeIndex int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "gets correct left index for a given child",
			fields: fields{
				nodes: []int{1, 5, 3, 9, 6, 8},
			},
			args: args{
				parentNodeIndex: 1,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MinHeap{
				nodes: tt.fields.nodes,
			}
			if got := h.LeftChildIndex(tt.args.parentNodeIndex); got != tt.want {
				t.Errorf("MinHeap.LeftChildIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinHeap_InsertNode(t *testing.T) {
	type args struct {
		node []int
	}
	tests := []struct {
		name string
		args args
		want *MinHeap
	}{
		{
			name: "it inserts and reshuffles the nodes correctly",
			args: args{
				node: []int{8, 9, 5, 6, 1, 3},
			},
			want: &MinHeap{[]int{1, 5, 3, 9, 6, 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MinHeap{}
			if got := h.InsertNode(tt.args.node...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinHeap.InsertNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinHeap_swapNodes(t *testing.T) {

	t.Run("it swaps nodes correctly", func(t *testing.T) {
		h := &MinHeap{
			nodes: []int{50, 100},
		}
		h.swapNodes(0, 1)
		if h.nodes[0] != 100 && h.nodes[1] != 50 {
			t.Error("did not swap correctly")
		}
	})

}
