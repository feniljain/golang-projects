package primitive

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
)

func createTempFile(r io.Reader, ext string) (*os.File, error) {
	f, err := ioutil.TempFile("static/input/", "")
	if err != nil {
		return nil, err
	}
	defer os.Remove(f.Name())
	file, err := os.Create(fmt.Sprintf("%s.%s", f.Name(), ext))
	if err != nil {
		return nil, err
	}
	io.Copy(file, r)
	return file, nil
}

func primitate(fileName string) error {
	cmd := exec.Command("/home/fenil/go/src/github.com/fogleman/primitive/main", "-i", "static/input/"+fileName, "-o", "static/output/"+fileName, "-n", "100")
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

//Primitive takes in the writer form of file and returns the name of the output file
func Primitive(r io.Reader, ext string) (string, error) {
	w, err := createTempFile(r, ext)
	if err != nil {
		return "", err
	}
	err = primitate(strings.Split(w.Name(), "/")[2])
	if err != nil {
		return "", err
	}
	go func() {
		timer := time.NewTimer(15 * time.Second)
		<-timer.C
		os.Remove(w.Name())
	}()
	return "static/output/" + strings.Split(w.Name(), "/")[2], nil
}
