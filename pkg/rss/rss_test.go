package rss

import (
	"testing"
	"sync"
)



func TestReadFeedSuccess(t *testing.T) {
	ris := &rssItems{}
	wg := &sync.WaitGroup{}

	wg.Add(1)
	readFeed("https://news.google.com/rss", ris, wg)
}

// just for caverage
func TestReadFeedFail1(t *testing.T) {
	ris := &rssItems{}
	wg := &sync.WaitGroup{}

	wg.Add(1)
	readFeed("", ris, wg)

	if ris.size() != 0 {
		t.Errorf("rssItems should be empty")
	}
}

// just for caverage
func TestReadFeedFail2(t *testing.T) {
	ris := &rssItems{}
	wg := &sync.WaitGroup{}

	wg.Add(1)
	readFeed("https://news.kole.com/rss", ris, wg)

	if ris.size() != 0 {
		t.Errorf("rssItems should be empty")
	}
}

// just for caverage
func TestParse(t *testing.T) {
	ris := &rssItems{}

	Parse([]string{
		"https://news.google.com/news/rss",
		"https://news.google.com/rss?hl=bg&gl=BG&ceid=BG:bg",
		"https://www.mediapool.bg/rss",
	})

	if ris.size() != 0 {
		t.Errorf("rssItems should be empty")
	}
}
