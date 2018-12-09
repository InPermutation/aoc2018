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

func oneCharDifferent(l string, r string) bool {
	if len(l) != len(r) {
		panic("length mismatch " + l + " " + r)
	}
	d := 0
	for i := 0; i < len(l); i++ {
		if l[i] != r[i] {
			d++
		}
	}
	return d == 1
}

func main() {
	dat, ferr := ioutil.ReadFile("input")
	check(ferr)
	text := string(dat)
	text = strings.Trim(text, "\r\n")
	ids := strings.Split(text, "\n")
	for ix, id := range ids {
		for i := ix + 1; i < len(ids); i++ {
			if oneCharDifferent(id, ids[i]) {
				fmt.Println(id, ids[i])
				return
			}
		}
	}
}
