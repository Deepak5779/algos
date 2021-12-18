package binarySearchBased

import "testing"

func Test_findNthRoot(t *testing.T) {
	type args struct {
		n int
		r int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"test1", args{n: 8, r: 3}, 2},
		{"test2", args{n: 16, r: 4}, 2},
		{"test3", args{n: 1000, r: 6}, 3.162},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNthRoot(tt.args.n, tt.args.r); got != tt.want {
				t.Errorf("findNthRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}
