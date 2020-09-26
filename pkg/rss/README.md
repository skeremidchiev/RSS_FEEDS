## Installation

```
go get -v github.com/skeremidchiev/RSS_FEEDS/pkg/rss
```

## How To Use

### Example

```
import "RSS/pkg/rss"

rss.Parse([]string{
    "https://news.google.com/news/rss",
    "https://news.google.com/rss?hl=bg&gl=BG&ceid=BG:bg",
    "https://www.mediapool.bg/rss",
})
```

**Parse** method returns array of **RssItem**.

## Important:

This package will work only for Google-like feeds.
Most other feeds have different tag names or tag structure.
(For example **_https://stackoverflow.com/feeds/_** uses **<entry>** instead of **<item>**.)

## Tests

Run ./test.sh

Or :

```
go test -coverprofile=rssPkg.out
go tool cover -html=rssPkg.out
```
