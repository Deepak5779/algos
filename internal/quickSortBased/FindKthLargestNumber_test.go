package quickSortBased

import "testing"

func Test_findKthLargestNumberInArray(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{nums: []int{7, 10, 4, 3, 20, 15}, k: 4}, 7},
		{"test2", args{nums: []int{1, 2, 4, 3, 5, 6}, k: 1}, 6},
		{"test3", args{nums: []int{1, 2, 4, 3, 5, 6}, k: 6}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargestNumberInArray(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargestNumberInArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPivotIndex(t *testing.T) {
	type args struct {
		nums  []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPivotIndex(tt.args.nums, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("getPivotIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swap(t *testing.T) {
	type args struct {
		nums []int
		i    int
		j    int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			swap(tt.args.nums, tt.args.i, tt.args.j)
		})
	}
}
