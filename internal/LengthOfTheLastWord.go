package main

import (
	"fmt"
	"strings"
)

func lengthOfTheLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}

	l := len(s) - 1

	//trim the whitespaces from last
	for l >= 0 && s[l] == ' ' {
		l--
	}

	c := 0
	for l >= 0 && s[l] != ' ' {
		l--
		c++
	}

	return c
}

func main() {
	s := "Hello World"
	fmt.Println("length of last word of the string is : ", lengthOfTheLastWord(s))

	s = "   fly me   to   the moon  "
	fmt.Println("length of last word of the string is : ", lengthOfTheLastWord(s))

	s = "luffy is still joyboy"
	fmt.Println("length of last word of the string is : ", lengthOfTheLastWord(s))

	fmt.Println(strings.TrimSpace("  Hello     there  "))
}
