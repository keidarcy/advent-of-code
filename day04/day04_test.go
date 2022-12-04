package main

import (
	"reflect"
	"testing"
)

func Test_countContain(t *testing.T) {

	input1 := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "sample input",
			args: input1,
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countContain(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countContain() = %v, want %v", got, tt.want)
			}
		})
	}

}
func Test_countOverlap(t *testing.T) {

	input1 := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "sample input",
			args: input1,
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOverlap(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countOverlap() = %v, want %v", got, tt.want)
			}
		})
	}

}
