package main

import (
	"reflect"
	"testing"
)

func Test_getSumOfThreeMax(t *testing.T) {

	input1 := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "sample input",
			args: input1,
			want: 45000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSumOfThreeMax(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSumOfThreeMax() = %v, want %v", got, tt.want)
			}
		})
	}

}
