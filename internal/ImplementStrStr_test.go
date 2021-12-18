package main

import "testing"

func Test_needlePointerInHaystack(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{"hello", "ll"}, 2},
		{"test2", args{"aaaaa", "bba"}, -1},
		{"test3", args{"", ""}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := needlePointerInHaystackNaive(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("needlePointerInHaystack() = %v, want %v", got, tt.want)
			}
		})
	}
}
