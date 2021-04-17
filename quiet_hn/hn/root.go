package hn

import (
	"fmt"
	"time"
)

var (
	cache           []NewsItem
	cacheExpiration time.Time = time.Now()
)

//Execute acts as a entry point for the client
func Execute() int {

	if time.Now().Sub(cacheExpiration) < 0 {
		fmt.Println("Used cached data")
		return 0
	}

	newsCount := 30
	var client Client

	client.Init("")

	ids, err := client.getIds(newsCount)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	news, err := getAllNews(newsCount, client, ids)

	if err != nil {
		fmt.Println(err)
		return 1
	}

	filteredNews := make([]NewsItem, newsCount)
	var nilNewsItem NewsItem

	cnt := 0
	for i, l := range news {
		if l != nilNewsItem {
			filteredNews[cnt] = news[i]
			cnt++
		}
	}
	fmt.Println(len(filteredNews))

	//timeChan := make(chan time.Time)
	//go cacheTimer(timeChan)
	cache = filteredNews
	cacheExpiration = time.Now().Add(1 * time.Minute)
	return 0
}

func cacheTimer(c chan time.Time) {
}

func getAllNews(newsCount int, client Client, ids []int) ([]NewsItem, error) {
	c := make(chan NewsItem)

	for i := 0; i < newsCount; i++ {
		go client.getItem(ids[i], c)
	}
	news := make([]NewsItem, newsCount)
	var nilNewsItem NewsItem
	invalidNews := 0
	for i := 0; i < len(news); i++ {
		l := <-c
		index := indexOf(ids, l.ID)
		news[index] = l
		if l.NewsType != "story" || i == 1 {
			news[index] = nilNewsItem
			invalidNews++
		}
	}
	if invalidNews > 0 {
		idsNew, err := client.getIds(invalidNews)
		if err != nil {
			return nil, err
		}
		newNews, err := getAllNews(invalidNews, client, idsNew)
		if err != nil {
			return nil, err
		}
		for _, e := range newNews {
			news = append(news, e)
		}
	}
	return news, nil
}

func removeIntAtIndex(s []int, i int) []int {
	return (append(s[:i], s[i+1:]...))
}

func removeNewsItemAtIndex(s []NewsItem, i int) []NewsItem {
	return (append(s[:i], s[i+1:]...))
}

func indexOf(si []int, e int) int {
	for i, element := range si {
		if element == e {
			return i
		}
	}
	return -1
}
