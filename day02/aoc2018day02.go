package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(f error) {
	if f != nil {
		panic(f)
	}
}

const chars = "abcdefghijklmnopqrstuvwxyz"

func main() {
	var twos, threes int

	dat, ferr := ioutil.ReadFile("input")
	check(ferr)
	text := string(dat)
	text = strings.Trim(text, "\r\n")
	for _, line := range strings.Split(text, "\n") {
		var hasTwo, hasThree bool
		for _, ch := range chars {
			switch strings.Count(line, string(ch)) {
			case 2:
				hasTwo = true
			case 3:
				hasThree = true
			}
		}
		if hasTwo {
			twos++
		}
		if hasThree {
			threes++
		}
	}
	fmt.Println(twos * threes)
}
