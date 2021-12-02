package linkedlist

import (
	"reflect"
	"testing"
)

func TestLinkedList_Append(t *testing.T) {
	t.Run("appends an item to the list", func(t *testing.T) {
		ll := &LinkedList{}
		ll.Append(30)
		if ll.firstNode == nil {
			t.Error("no first node exists")
			if ll.firstNode.Value != 30 {
				t.Error("node missing value")
			}
		}
	})

	t.Run("appends two items to the list", func(t *testing.T) {
		ll := &LinkedList{}
		ll.Append(30)
		ll.Append(40)
		if ll.firstNode.NextNode == nil {
			t.Error("no first node exists")
			if ll.firstNode.NextNode.Value != 40 {
				t.Error("node missing value")
			}
		}
	})

}

func TestLinkedList_Prepend(t *testing.T) {

	t.Run("it prepends an item to an list with >1 count", func(t *testing.T) {
		ll := &LinkedList{}
		ll.Append(10)
		ll.Prepend(20)
		if ll.firstNode.Value != 20 {
			t.Error("did not prepend to list")
		}
	})

}

func TestLinkedList_Get(t *testing.T) {

	tests := []struct {
		name    string
		args    int
		want    int
		wantErr bool
	}{
		{
			name:    "throws an error when index is < 0",
			args:    -3,
			want:    0,
			wantErr: true,
		},
		{
			name:    "gets item value at index 0",
			args:    0,
			want:    10,
			wantErr: false,
		},
		{
			name:    "gets item value at index 1",
			args:    1,
			want:    20,
			wantErr: false,
		},
		{
			name:    "returns -1 for an item that doesn't exist and index is >0",
			args:    5,
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := &LinkedList{}
			ll.Append(10)
			ll.Append(20)
			ll.Append(30)
			got, err := ll.Get(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("LinkedList.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got.Value, tt.want) {
				t.Errorf("LinkedList.Get() = %v, want %v", got.Value, tt.want)
			}
		})
	}
}

func TestLinkedList_Contains(t *testing.T) {
	tests := []struct {
		name string
		args interface{}
		want int
	}{
		{
			name: "contains returns an index for value 20",
			args: 20,
			want: 5,
		},
		{
			name: "contains returns -1 for an value that doesn't exist",
			args: 9000,
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := &LinkedList{}
			ll.Append(0).Append(10).Append(60).Append(30).Append(40).Append(20)
			if got := ll.Contains(tt.args); got != tt.want {
				t.Errorf("LinkedList.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Remove(t *testing.T) {
	tests := []struct {
		name string
		args int
		want bool
	}{
		{
			name: "remove index 2 successfully",
			args: 2,
			want: true,
		},
		{
			name: "returns false when it cannot remove an item due to it not existing",
			args: 5,
			want: false,
		},
		{
			name: "returns false when a negative index is passed",
			args: -1,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := &LinkedList{}
			ll.Append(40).Append(20).Append(10).Append(5).Append(0)
			if got := ll.Remove(tt.args); got != tt.want {
				t.Errorf("LinkedList.Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Count(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "counts the amount of items correctly",
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := &LinkedList{}
			ll.Append(1).Append(2).Append(3)
			if got := ll.Count(); got != tt.want {
				t.Errorf("LinkedList.Count() = %v, want %v", got, tt.want)
			}
		})
	}
}
