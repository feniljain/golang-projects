package main

import (
	"fmt"
	"sort"
)

type aStruct struct {
	person string
	height int
	width  int
}

func main() {
	mySlice := make([]aStruct, 0)

	a := aStruct{person: "Someone", height: 10, width: 10}
	mySlice = append(mySlice, a)
	a = aStruct{person: "Someone1", height: 13, width: 10}
	mySlice = append(mySlice, a)
	a = aStruct{person: "Someone2", height: 18, width: 10}
	mySlice = append(mySlice, a)
	a = aStruct{person: "Someone3", height: 1, width: 10}
	mySlice = append(mySlice, a)

	fmt.Println("0:", mySlice)
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].width < mySlice[j].width
	})
	fmt.Println("<:", mySlice)
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].width > mySlice[j].width
	})
	fmt.Println(">:", mySlice)
}
