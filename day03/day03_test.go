package main

import (
	"reflect"
	"testing"
)

func Test_getPriority1(t *testing.T) {

	input1 := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "sample input",
			args: input1,
			want: 157,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPriority1(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPriority1() = %v, want %v", got, tt.want)
			}
		})
	}

}
func Test_getPriority2(t *testing.T) {

	input1 := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "sample input",
			args: input1,
			want: 70,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPriority2(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPriority2() = %v, want %v", got, tt.want)
			}
		})
	}

}
