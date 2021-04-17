package hn

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	apiBase = "https://hacker-news.firebaseio.com/v0"
)

//Client represents client to use http.Client
type Client struct {
	apiBase string
}

//Init function initializes the api
func (c *Client) Init(url string) {
	if url != "" {
		c.apiBase = url
	}
	c.apiBase = apiBase
}

func (c *Client) getIds(n int) ([]int, error) {
	resp, err := http.Get(fmt.Sprintf("%s/topstories.json", c.apiBase))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var ids []int
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&ids)
	if err != nil {
		return nil, err
	}
	var filteredIds []int
	for i, id := range ids {
		if i < n {
			filteredIds = append(filteredIds, id)
		}
	}
	return filteredIds, nil
}

func (c *Client) getItem(id int, ch chan NewsItem) {
	resp, err := http.Get(fmt.Sprintf("%s/item/%d.json", c.apiBase, id))
	var nilNewsItem NewsItem
	if err != nil {
		ch <- nilNewsItem
	}
	defer resp.Body.Close()
	nb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- nilNewsItem
	}
	var news NewsItem
	err = json.Unmarshal(nb, &news)
	if err != nil {
		ch <- nilNewsItem
	}
	ch <- news
}

//NewsItem represents a news item
type NewsItem struct {
	By       string `json:"by"`
	ID       int    `json:"id"`
	Title    string `json:"title"`
	NewsType string `json:"type"`
	URL      string `json:"url"`
}
