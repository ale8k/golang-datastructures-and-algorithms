package queue

import (
	"reflect"
	"testing"
)

func TestIntQueue_Enqueue(t *testing.T) {
	type args struct {
		elements []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "it queues items",
			args: args{
				elements: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &IntQueue{}
			q.Enqueue(tt.args.elements...)
			if !reflect.DeepEqual(q.queue, tt.want) {
				t.Errorf("IntQueue.Dequeue() = %v, want %v", q.queue, tt.want)
			}
		})
	}
}

func TestIntQueue_Dequeue(t *testing.T) {
	type fields struct {
		queue []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "it dequeues the last element correctly",
			fields: fields{[]int{40, 20, 10}},
			want:   10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &IntQueue{
				queue: tt.fields.queue,
			}
			if got := q.Dequeue(); got != tt.want {
				t.Errorf("IntQueue.Dequeue() = %v, want %v", got, tt.want)
			}
		})
	}
}
