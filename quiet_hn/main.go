package main

import (
	"fmt"
	"net/http"
	"quiet_hn/hn"
)

var (
	userCount int = 0
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Namastey Duniyaa!")
		hn.Execute()
		userCount++
		fmt.Printf("%d\n", userCount)
	})

	http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil)
}
