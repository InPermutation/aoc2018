package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(f error) {
	if f != nil {
		panic(f)
	}
}

func main() {
	f := 0
	seen := make(map[int]bool)
	seen[0] = true

	for {
		dat, ferr := ioutil.ReadFile("input1")
		check(ferr)
		text := string(dat)
		text = strings.Trim(text, "\r\n")
		for _, line := range strings.Split(text, "\n") {
			for _, s := range strings.Split(line, ", ") {
				d, err := strconv.Atoi(s)
				check(err)
				f += d
				if seen[f] {
					fmt.Println(f)
					return
				}
				seen[f] = true
			}
		}
	}
}
