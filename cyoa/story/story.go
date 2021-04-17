package story

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func CLI(args []string) int {
	var story Story
	data, err := ioutil.ReadFile("./gopher.json")
	if err != nil {
		fmt.Println("Error occurred while reading file")
		return 1
	}
	err = json.Unmarshal([]byte(string(data)), &story)
	if err != nil {
		fmt.Println("Error occurred while parsing data")
		return 1
	}
	fmt.Println("Starting the server on :8080")
	type user struct{ UserName string }
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.New("Some Template")
		t, _ = t.Parse(`
			<html>
				<head>
				</head>
				<body>
					<h1>Welcome!</h1>
					hello {{.Title}}
				</body>
			</html>
		`)
		//user{
		//	UserName: "s",
		//}
		fmt.Println(story["intro"].Paragraphs)
		t.Execute(w, story["intro"])
	})
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi pe: Namastey Duniyaa!")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
	return 0
}

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
