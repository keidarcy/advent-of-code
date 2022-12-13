package main

import (
	"reflect"
	"testing"
)

func Test_part1(t *testing.T) {

	input1 := `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "day 12 part 1",
			args: input1,
			want: 31,
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

	input1 := `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "day 12 part 2",
			args: input1,
			want: 29,
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
