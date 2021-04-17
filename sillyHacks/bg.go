package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/reujab/wallpaper"
)

//Resp represents the response structure
type Resp struct {
	Type    string
	Success string
	Message string
	Data    string
}

func main() {
	background, err := wallpaper.Get()

	if err != nil {
		panic(err)
	}

	resp, err := http.Get("https://boiling-dawn-19818.herokuapp.com/api/getImage")
	if err != nil {
		log.Fatal("Error getting background")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading image URL")
	}

	var respBody Resp
	fmt.Print(string(body))
	fmt.Print(json.Unmarshal(body, &respBody))
	fmt.Print(respBody.Data)

	fmt.Println("Current wallpaper:", background)
	//wallpaper.SetFromFile("/usr/share/backgrounds/gnome/adwaita-day.jpg")
	//wallpaper.SetFromURL("https://i.imgur.com/pIwrYeM.jpg")
	wallpaper.SetFromURL(respBody.Data)
}
