package main

import (
	"fmt"

	"RSS/pkg/rss"
)

func main() {
	fmt.Println("Staring ...")

	rss.Parse([]string{
		// "https://news.google.com/news/rss",
		"https://news.google.com/rss?hl=bg&gl=BG&ceid=BG:bg",
		"https://www.mediapool.bg/rss",
	})
}
