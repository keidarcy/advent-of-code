package main

import (
	"reflect"
	"testing"
)

func Test_part1(t *testing.T) {

	input1 := `30373
25512
65332
33549
35390`
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "day 08 part 1",
			args: input1,
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPart1(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPart1() = %v, want %v", got, tt.want)
			}
		})
	}

}

func Test_part2(t *testing.T) {

	input1 := `30373
25512
65332
33549
35390`
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "day 08 part 2",
			args: input1,
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPart2(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPart2() = %v, want %v", got, tt.want)
			}
		})
	}

}
