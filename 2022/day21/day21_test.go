package main

import (
	"reflect"
	"testing"
)

func Test_part1(t *testing.T) {

	input1 := `root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32`
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "day 21 part 1",
			args: input1,
			want: 152,
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

	input1 := `root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32`
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "day 21 part 2",
			args: input1,
			want: 301,
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
