package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		param int
		slice []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "it correctly finds the expected index",
			args: args{
				param: 4,
				slice: []int{6, 8, 2, 4, 6},
			},
			want: 3,
		},
		{
			name: "it returns -1 when out of bounds > length",
			args: args{
				param: 10,
				slice: []int{6, 8, 2, 4, 6},
			},
			want: -1,
		},
		{
			name: "it returns -1 when mid < 0",
			args: args{
				param: -1,
				slice: []int{6, 8, 2, 4, 6},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.param, tt.args.slice); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
