package main

import (
	"fmt"
	"imageTransform/primitive"
	"net/http"
	"path/filepath"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		file, header, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()
		fmt.Println(header.Filename)

		fileName, err := primitive.Primitive(file, filepath.Ext(header.Filename)[1:])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Println(fileName)
		http.Redirect(w, r, "/"+fileName, http.StatusFound)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
		<html>
			<body>
				<label>Upload Photo</label></br>
				<form action="/upload" method="post" enctype="multipart/form-data">
					<input type="file" name="image"/>
					<input type="submit"/>
				</form>
			</body>
		</html>
		`)
	})
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":3000", mux)
	//cmd, err := exec.Command("go", "run /home/fenil/go/src/github.com/fogleman/primitive/main.go").CombinedOutput()
	//cmd := &exec.Cmd{
	//	Path: "/home/fenil/go/src/github.com/fogleman/primitive/main",
	//	Args: []string{"-i", "images/input.png"},
	//	//, "-o s.png -n 100"
	//	//Stdout: os.Stdout,
	//	//Stderr: os.Stderr,
	//}
}
