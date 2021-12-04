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
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{
			name:    "it dequeues the last element correctly",
			fields:  fields{[]int{40, 20, 10}},
			want:    10,
			wantErr: false,
		},
		{
			name:    "it attempts to dequeue empty list",
			fields:  fields{[]int{}},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &IntQueue{
				queue: tt.fields.queue,
			}
			got, err := q.Dequeue()
			if err == nil && got != tt.want {
				t.Errorf("IntQueue.Dequeue() = %v, want %v", got, tt.want)
			}
			if err == nil && tt.wantErr {
				t.Errorf("IntQueue.Dequeue() did not error as expected")
			}
		})
	}
}
