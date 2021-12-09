package shellsort

import (
	"reflect"
	"testing"
)

func TestShellSort(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "it sorts the slice correctly",
			args: args{
				slice: []int{6, 1, 8, 9},
			},
			want: []int{1, 6, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ShellSort(tt.args.slice)
			if !reflect.DeepEqual(tt.want, tt.args.slice) {
				t.Errorf("the sort failed, want: %v, got: %v", tt.want, tt.args.slice)
			}
		})
	}
}
