package bubblesort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBubbleSortInt(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "it sorts the slice",
			args: []int{6, 2, 8, 3, 1, 10},
			want: []int{1, 2, 3, 6, 8, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSortInt(tt.args)
			fmt.Print(tt.args)
			if !reflect.DeepEqual(tt.args, tt.want) {
				t.Error("the sort failed")
			}
		})
	}
}
