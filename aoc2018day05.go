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

func input() string {
	dat, ferr := ioutil.ReadFile("input")
	check(ferr)
	text := string(dat)
	text = strings.Trim(text, "\r\n")
	return text
}

func react(polymer string) string {
	i := 0
	for i < len(polymer)-1 {
		if reactable(polymer[i:i+1], polymer[i+1:i+2]) {
			polymer = polymer[0:i] + polymer[i+2:]
			i--
			if i < 0 {
				i = 0
			}
		} else {
			i++
		}
	}
	return polymer
}

func reactable(l string, r string) bool {
	return l != r &&
		(l == strings.ToUpper(r) ||
			r == strings.ToUpper(l))
}

func main() {
	text := input()
	for {
		next := react(text)
		if next == text {
			break
		}
		text = next
		fmt.Println(len(text))
	}
	fmt.Println(len(text))
}
