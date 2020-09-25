package rss

import (
	"fmt"

	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
)

func Parse(urls []string) []RssItem {
	items := &rssItems{}
	wg := &sync.WaitGroup{}

	for _, u := range urls {
		wg.Add(1)
		go readFeed(u, items, wg)
	}

	wg.Wait()

	// TODO: remove
	fmt.Printf("ITEMS: %+v\n", items)
	fmt.Println(items.size())

	return items.getItems()
}

// The readFeed function reads feed from specific url
func readFeed(_url string, ris *rssItems, wg *sync.WaitGroup) {
	defer wg.Done()

	// check if url valid
	u, err := url.ParseRequestURI(_url)
	if err != nil {
		fmt.Println("URL: %s \t Error: %s", _url, err.Error())
		return
	}

	// get response
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println("URL: %s \t Error: %s", _url, err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	rss := &rss{}
	err = xml.Unmarshal(body, rss)
	if err != nil {
		fmt.Println("Unmarshal error: ", err)
		return
	}

	ris.addFromRss(rss)
}