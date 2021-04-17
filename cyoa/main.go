package main

import (
	"cyoa/story"
	"os"
)

func main() {
	os.Exit(story.CLI(os.Args[1:]))
}

//--Parse Json
//--Create a simple web server for rendering data
//--Figure out how to use routes
//Figure out how to use html for rendering
//How to redirect on other links
//Join everything
