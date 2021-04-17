package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {
	correctString := "Correct"
	incorrectString := "Incorrect"
	timesUpString := "Times Up"

	var csvName string
	var timeLimit string
	flag.StringVar(&csvName, "csv", "problems.csv", "specify csv name")
	flag.StringVar(&timeLimit, "time", "2", "specify time limit")
	flag.Parse()

	csvfile, err := os.Open(csvName)
	if err != nil {
		exit("Error while reading file!")
	}

	r := csv.NewReader(csvfile)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Error parsing csv")
	}

	problems := parseLines(lines)

	correct := 0
	for i, p := range problems {
		timelimit, _ := strconv.Atoi(timeLimit)
		timer := time.NewTimer(time.Duration(int64(timelimit)) * time.Second)

		fmt.Printf("Problem #%d: %s= \n", i+1, p.q)

		c := make(chan string)

		var answer string
		go func(channel chan string, p problem) {
			fmt.Scanf("%s\n", &answer)
			if answer == p.a {
				fmt.Println(correctString)
				correct++
				c <- correctString
			} else {
				fmt.Println(incorrectString)
				c <- incorrectString
			}
		}(c, p)

		go func(channel chan string) {
			<-timer.C
			c <- timesUpString
		}(c)

		resp := <-c
		if resp == timesUpString {
			exit(timesUpString)
		} else {
			continue
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
