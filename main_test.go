package main

import (
	"reflect"
	"testing"
)

func Test_filter(t *testing.T) {
	type args struct {
		array []int
		f     func(int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{[]int{1, 2, 3, 4, 5}, func(n int) bool { return n%2 == 0 }},
			want: []int{2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filter(tt.args.array, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tryDecodeJson(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want Message
	}{
    {
    name: "JSON Test",
    args: args{[]byte(`{"chatroom": 1, "text": "Hello", "username": "John"}`)},
    want: Message{1, "Hello", "John"},
    },
  }
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tryDecodeJson(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tryDecodeJson() = %v, want %v", got, tt.want)
			}
		})
	}
}
