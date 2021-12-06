package linearsearch

import "testing"

func TestPrimitiveSearch(t *testing.T) {
	type args struct {
		param interface{}
		slice interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "it finds the correct indices for a given param",
			args: args{
				param: 5,
				slice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			want: 4,
		},
		{
			name: "it finds the correct indices for a given param that does not exist",
			args: args{
				param: 69,
				slice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrimitiveSearch(tt.args.param, tt.args.slice); got != tt.want {
				t.Errorf("PrimitiveSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
