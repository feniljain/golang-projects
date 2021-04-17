package renamer

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

//Execute acts as the starting point of the package
func Execute() int {
	walker("sample")
	return 0
}

func listAndRenameFiles(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, f := range files {
		if !f.IsDir() {
			fmt.Println(path + "/" + f.Name())
			//fileSuffix := strings.Split(f.Name(), "_")[1]
			////fmt.Println(fileSuffix)
			//os.Rename(path+"/"+f.Name(), path+"/"+"something"+fileSuffix)
			//fmt.Println("New File Name:", path+"/"+"something"+fileSuffix)
			re := regexp.MustCompile(`something00[0-9]\.txt`)
			fmt.Println(re.Match([]byte(f.Name())))
		}
	}
	return nil
}

func walker(filePath string) int {
	err := listAndRenameFiles(filePath)
	if err != nil {
		return 1
	}
	err = filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		currentDir := strings.Split(filePath, "/")[len(strings.Split(filePath, "/"))-1]
		if info.IsDir() && info.Name() != "renamer" && info.Name() != currentDir {
			//fmt.Println("Found the dir:", info.Name())
			//fmt.Println("Current Path:", path)
			//fmt.Println("Current file path:", filePath)
			walker(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}
