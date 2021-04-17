package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide filename!")
		os.Exit(1)
	}

	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening %s: %s", filename, err)
		os.Exit(1)
	}
	defer f.Close()

	// USING IO
	//buf := make([]byte, 8)
	//if _, err := io.ReadFull(f, buf); err != nil {
	//	if err == io.EOF {
	//		err = io.ErrUnexpectedEOF
	//	}
	//}
	//io.WriteString(os.Stdout, string(buf))
	//fmt.Println()

	// USING BUFFERIO
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if scanner.Err() != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(line)
	}
}
