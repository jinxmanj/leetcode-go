package main

import (
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for range tests {
		main()
	}
}

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. twoSum() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
