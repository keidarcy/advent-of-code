package main

import (
	"reflect"
	"testing"
)

func Test_part1(t *testing.T) {

	input1 := `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1`
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "day 11 part 1",
			args: input1,
			want: 10605,
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

	input1 := `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1`
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "day 11 part 2",
			args: input1,
			want: 2713310158,
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
