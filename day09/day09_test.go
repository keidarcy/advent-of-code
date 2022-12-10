package main

import (
	"reflect"
	"testing"
)

func Test_part1(t *testing.T) {

	input1 := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "day 09 part 1",
			args: input1,
			want: 13,
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

	input1 := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "day 09 part 2",
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
