package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func check(f error) {
	if f != nil {
		panic(f)
	}
}

func atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}

func input() string {
	dat, ferr := ioutil.ReadFile("input")
	check(ferr)
	text := string(dat)
	text = strings.Trim(text, "\r\n")
	return text
}

func layout(text string) (fabric [1000][1000][]int, maxId int) {
	for _, claim := range strings.Split(text, "\n") {
		rg := strings.Split(claim, " ")
		source := strings.Split(strings.Trim(rg[2], ":"), ",")
		id := atoi(strings.Trim(rg[0], "#"))
		if id > maxId {
			maxId = id
		}
		x, y := atoi(source[0]), atoi(source[1])
		size := strings.Split(rg[3], "x")
		w, h := atoi(size[0]), atoi(size[1])
		for j := y ; j < y + h ; j++ {
			for i := x ; i < x + w ; i++ {
				fabric[i][j] = append(fabric[i][j], id)
			}
		}
	}
	return
}

func process(fabric [1000][1000][]int, maxId int) (c int, idWithOverlaps map[int]bool) {
	idWithOverlaps = make(map[int]bool, 1000)
	for _, y := range fabric {
		for _, x := range y {
			if len(x) > 1 {
				c++
				for _, id := range x {
					idWithOverlaps[id] = true
				}
			}
		}
	}
	return
}

func main() {
	text := input()
	fabric, maxId := layout(text)
	c, idWithOverlaps := process(fabric, maxId)

	fmt.Println("Overlapping square inches", c)
	for id := 1 ; id <= maxId ; id++ {
		if !idWithOverlaps[id] {
			fmt.Println(id, "has no overlaps")
		}
	}
}
