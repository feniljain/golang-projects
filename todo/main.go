package main

import (
	//"fmt"
	"todo/cmd"
	"todo/db"
)

func init() {
	db.Init()
}

func main() {
	cmd.Execute()
}
