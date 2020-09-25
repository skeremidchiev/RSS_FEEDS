package rss

import (
	"encoding/xml"
	"time"
	"sync"
)

// The RssItem represents a single feed
type RssItem struct {
	Title       string    `json:"title"`
	Source      string    `json:"source"`
	SourceURL   string    `json:"sourceURL"`
	Link        string    `json:"link"`
	PublishDate time.Time `json:"publishDate"`
	Description string    `json:"description"`
}

// The rssItems is a thread safe structure of multiple feeds
type rssItems struct {
	sync.RWMutex
	items []RssItem
}

// The addItem is thread safe set function for rssItems
func (ris *rssItems) addItem(item *RssItem) {
	ris.Lock()
	ris.items = append(ris.items, *item)
	ris.Unlock()
}

// The getItems is thread safe accessor function for rssItems
func (ris *rssItems) getItems() []RssItem {
	ris.RLock()
	defer ris.RUnlock()
	return ris.items

}

// The addFromRss importing all items from rss feed to rssItems in a thread safe manner
func (ris *rssItems) addFromRss(r *rss) {
	source := r.Channel.Title
	sourceURL := r.Channel.Link

	for _, item := range r.Channel.Items {
		rssItem := &RssItem{}

		if item.Source.Source == "" {
			rssItem.Source = source
		} else {
			rssItem.Source =  item.Source.Source
		}

		if item.Source.SourceURL == "" {
			rssItem.SourceURL = sourceURL
		} else {
			rssItem.SourceURL =  item.Source.SourceURL
		}

		rssItem.Title = item.Title
		rssItem.Link = item.Link
		rssItem.PublishDate = item.PublishDate.Time
		rssItem.Description = item.Description

		ris.addItem(rssItem)
	}
}

// The addFromRss importing all items from rss feed to rssItems in a thread safe manner
func (ris *rssItems) size() int {
	return len(ris.items)
}

// The rss is the top structure from rss feed
type rss struct {
	Channel channel `xml:"channel"`
}

// The channel hold rss feed data and related feed items
type channel struct {
	Title     string     `xml:"title"`
	Copyright string     `xml:"copyright"`
	Link      string     `xml:"link"`

	LastBuildDate customTime `xml:"lastBuildDate"`
	PublishDate   customTime `xml:"pubDate"`

	Items []item `xml:"item"`
}

// item
type item struct {
	Title       string     `xml:"title"`
	Link        string     `xml:"link"`
	Description string     `xml:"description"`
	PublishDate customTime `xml:"pubDate"`
	Source      source     `xml:"source"`
}

// customTime - used for RFC1123Z conversion
type customTime struct {
	time.Time
}

// UnmarshalXML - converts time from RFC1123Z
func (c *customTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse(time.RFC1123/*Z*/, v)
	if err != nil {
		return err
	}

	*c = customTime{parse}
	return nil
}

type source struct {
	SourceURL    string   `xml:"url,attr"`
	Source       string   `xml:"source"`
}