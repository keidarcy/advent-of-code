package main

import (
	"reflect"
	"testing"
)

func Test_getSumOfThreeMax(t *testing.T) {

	input1 := `A Y
B X
C Z`
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "sample input",
			args: input1,
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateScore(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateScore() = %v, want %v", got, tt.want)
			}
		})
	}

}
