package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"time"
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

func input() []string {
	dat, ferr := ioutil.ReadFile("input")
	check(ferr)
	text := string(dat)
	text = strings.Trim(text, "\r\n")
	return strings.Split(text, "\n")
}

func lex(row string) (dt time.Time, what string) {
	sdt := strings.SplitN(row[1:], "]", 2)[0]
	dt, err := time.Parse("2006-01-02 15:04", sdt)
	check(err)
	what = strings.SplitN(row, " ", 3)[2]
	return
}

type sleepers = map[string][]time.Time

func collate(text []string) (logs sleepers) {
	var duty string
	logs = make(map[string][]time.Time)
	for _, event := range text {
		when, what := lex(event)
		if strings.HasPrefix(what, "Guard #") {
			duty = what
		} else if what == "falls asleep" {
			logs[duty] = append(logs[duty], when)
		} else if what == "wakes up" {
			logs[duty] = append(logs[duty], when)
		} else {
			panic(event)
		}
	}
	return
}

func sleepiest(logs sleepers) (guard string) {
	longest := 0.0
	for id, schedule := range logs {
		totalMinutes := 0.0
		for i := 0; i < len(schedule); i += 2 {
			start, stop := schedule[i], schedule[i+1]
			d := stop.Sub(start)
			totalMinutes += d.Minutes()
		}
		if totalMinutes > longest {
			longest = totalMinutes
			guard = id
		}
	}
	return guard
}

func mostCommon(schedule []time.Time) (minute int) {
	slept := 0
	for i := 0; i < 60; i++ {
		sum := 0
		for j := 0; j < len(schedule); j += 2 {
			start, stop := schedule[j], schedule[j+1]
			if start.Minute() <= i && stop.Minute() >= i {
				sum++
			}

		}
		if sum > slept {
			slept, minute = sum, i
		}
	}
	return
}

func main() {
	text := input()
	sort.Strings(text)

	logs := collate(text)
	chose := sleepiest(logs)
	minute := mostCommon(logs[chose])
	fmt.Println(chose, minute)
}
