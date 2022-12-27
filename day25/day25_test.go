package main

import (
	"reflect"
	"testing"
)

func Test_part1(t *testing.T) {

	input1 := `1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122`
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "day 25 part 1",
			args: input1,
			want: "2=-1=0",
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

