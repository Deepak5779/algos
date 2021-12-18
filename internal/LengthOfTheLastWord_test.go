package main

import "testing"

func Test_lengthOfTheLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{"Hello World"}, 5},
		{"test2", args{"   fly me   to   the moon  "}, 4},
		{"test3", args{"luffy is still joyboy"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfTheLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfTheLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
