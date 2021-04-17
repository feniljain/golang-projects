package parser

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func CLI() int {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error!", err)
		return 1
	}
	//node, err := html.Parse(resp.Body)
	//if err != nil {
	//	fmt.Println("Error!", err)
	//	return 1
	//}
	//fmt.Println(extractNodes(node))

	//defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(bodyBytes)
	if err != nil {
		fmt.Println("Error!")
		return 1
	}
	fmt.Println(string(bodyBytes))
	x, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("-", x)
	//sampleLink := `
	//<html>
	//<body>
	//	<a href="/dog">
	//		<span>Something</span>
	//	</a>
	//</body>
	//</html>
	//`
	bs, err := ioutil.ReadFile("ex1.html")
	if err != nil {
		log.Fatal("Gone!")
		return 1
	}
	r := bytes.NewReader(bs)
	//str1 := bytes.NewReader([]byte(sampleLink))
	//strParse, err := html.Parse(str1)
	//if err != nil {
	//	log.Fatal(err)
	//	fmt.Println(err)
	//	return 1
	//}
	//fmt.Println(extractNodes(strParse))
	doc, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
		return 1
	}
	fmt.Println(extractNodes(doc))
	return 0
}

func extractNodes(n *html.Node) string {
	fmt.Println("-", n.Type)
	if n.Type == html.ElementNode && n.Data == "a" {
		return "Done"
	}
	return "No"
}
