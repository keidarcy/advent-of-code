package main

import (
	"reflect"
	"testing"
)

func Test_getSumOfThreeMax(t *testing.T) {
	input := `1000
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
	want := 45000
	t.Run("test input", func(t *testing.T) {
		if got := getSumOfThreeMax(input); !reflect.DeepEqual(got, want) {
			t.Errorf("getSumOfThreeMax() = %v, want %v", got, want)
		}
	})

}
