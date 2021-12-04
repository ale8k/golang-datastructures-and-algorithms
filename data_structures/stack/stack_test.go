package stack

import (
	"reflect"
	"testing"
)

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name string
		want interface{}
	}{
		{
			name: "it pops an item off the stack correctly",
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{}
			s.Push(3)
			if got := s.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	type fields struct {
		store []interface{}
	}
	type args struct {
		data []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				store: tt.fields.store,
			}
			s.Push(tt.args.data...)
		})
	}
}

func TestStack_Count(t *testing.T) {
	type fields struct {
		store []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				store: tt.fields.store,
			}
			if got := s.Count(); got != tt.want {
				t.Errorf("Stack.Count() = %v, want %v", got, tt.want)
			}
		})
	}
}
