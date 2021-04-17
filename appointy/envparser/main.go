package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("../deployment.raw.env")
	if err != nil {
		fmt.Println(err)
	}

	str := string(b)

	var envs string
	s := strings.Split(str, ";")
	for _, env := range s {
		envs += env + "\n"
	}

	f, err := os.Create(".env")
	if err != nil {
		fmt.Println(err)
	}

	_, err = f.Write([]byte(envs))
	if err != nil {
		fmt.Println(err)
	}
}
