package rss

import (
	"testing"
	"encoding/xml"
)


func TestRssItems(t *testing.T) {
	ris := &rssItems{}
	ri := &RssItem{}

	ris.addItem(ri)
	s := ris.size()
	if s != 1 {
		t.Errorf("rssItems size, got %d, expected %d", s, 1)
	}
}

func TestAddFromRssOK(t *testing.T) {
	data :=
	`<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
	<rss version="2.0" xmlns:media="http://search.yahoo.com/mrss/">
	<channel>
	<generator>NFE/5.0</generator>
	<title>Top stories - Google News</title>
	<link>https://news.google.com/?hl=en-US&amp;gl=US&amp;ceid=US:en</link>
	<language>en-US</language><webMaster>news-webmaster@google.com</webMaster>
	<copyright>2020 Google Inc.</copyright>
	<lastBuildDate>Sat, 26 Sep 2020 08:50:05 GMT</lastBuildDate>
	<description>Google News</description>
	<item>
		<title>Maher rips Barrett, Trump's expected Supreme Court pick: 'She's a f---ing nut' - Fox News</title>
		<link>https://news.google.com/__i/rss/rd/articles/CBMibGh0dHBzOi8vd3d3LmZveG5ld3MuY29tL2VudGVydGFpbm1lbnQvbWFoZXItcmlwcy1iYXJyZXR0LXRydW1wcy1leHBlY3RlZC1zdXByZW1lLWNvdXJ0LXBpY2stc2hlcy1hLWYtaW5nLW51dNIBcGh0dHBzOi8vd3d3LmZveG5ld3MuY29tL2VudGVydGFpbm1lbnQvbWFoZXItcmlwcy1iYXJyZXR0LXRydW1wcy1leHBlY3RlZC1zdXByZW1lLWNvdXJ0LXBpY2stc2hlcy1hLWYtaW5nLW51dC5hbXA?oc=5</link>
		<guid isPermaLink="false">52781070340229</guid>
		<pubDate>Sat, 26 Sep 2020 05:04:18 GMT</pubDate>
		<description>&lt;ol&gt;&lt;li&gt;&lt;a href="https://news.google.com/__i/rss/rd/articles/CBMibGh0dHBzOi8vd3d3LmZveG5ld3MuY29tL2VudGVydGFpbm1lbnQvbWFoZXItcmlwcy1iYXJyZXR0LXRydW1wcy1leHBlY3RlZC1zdXByZW1lLWNvdXJ0LXBpY2stc2hlcy1hLWYtaW5nLW51dNIBcGh0dHBzOi8vd3d3LmZveG5ld3MuY29tL2VudGVydGFpbm1lbnQvbWFoZXItcmlwcy1iYXJyZXR0LXRydW1wcy1leHBlY3RlZC1zdXByZW1lLWNvdXJ0LXBpY2stc2hlcy1hLWYtaW5nLW51dC5hbXA?oc=5" target="_blank"&gt;Maher rips Barrett, Trump's expected Supreme Court pick: 'She's a f---ing nut'&lt;/a&gt;&amp;nbsp;&amp;nbsp;&lt;font color="#6f6f6f"&gt;Fox News&lt;/font&gt;&lt;/li&gt;&lt;li&gt;&lt;a href="https://news.google.com/__i/rss/rd/articles/CBMiK2h0dHBzOi8vd3d3LnlvdXR1YmUuY29tL3dhdGNoP3Y9U1hib2xNcnhXVk3SAQA?oc=5" target="_blank"&gt;Trump Plans to Name Barrett to Replace Ginsburg on Supreme Court&lt;/a&gt;&amp;nbsp;&amp;nbsp;&lt;font color="#6f6f6f"&gt;Bloomberg Politics&lt;/font&gt;&lt;/li&gt;&lt;li&gt;&lt;a href="https://news.google.com/__i/rss/rd/articles/CBMiTmh0dHBzOi8vd3d3LmZveG5ld3MuY29tL29waW5pb24vYW15LWNvbmV5LWJhcnJldHQtZmFpdGgtY291bnRyeS1tYWdnaWUtZ2FybmV0dNIBUmh0dHBzOi8vd3d3LmZveG5ld3MuY29tL29waW5pb24vYW15LWNvbmV5LWJhcnJldHQtZmFpdGgtY291bnRyeS1tYWdnaWUtZ2FybmV0dC5hbXA?oc=5" target="_blank"&gt;Maggie Garnett: Judge Amy Coney Barrett has shown me how to both live my faith and serve my country&lt;/a&gt;&amp;nbsp;&amp;nbsp;&lt;font color="#6f6f6f"&gt;Fox News&lt;/font&gt;&lt;/li&gt;&lt;li&gt;&lt;a href="https://news.google.com/__i/rss/rd/articles/CBMiSGh0dHBzOi8vd3d3Lm55dGltZXMuY29tLzIwMjAvMDkvMjUvb3Bpbmlvbi9hbXktY29uZXktYmFycmV0dC1zY290dXMuaHRtbNIBTGh0dHBzOi8vd3d3Lm55dGltZXMuY29tLzIwMjAvMDkvMjUvb3Bpbmlvbi9hbXktY29uZXktYmFycmV0dC1zY290dXMuYW1wLmh0bWw?oc=5" target="_blank"&gt;Would Amy Coney Barrett on the Supreme Court Cost the G.O.P. the Senate?&lt;/a&gt;&amp;nbsp;&amp;nbsp;&lt;font color="#6f6f6f"&gt;The New York Times&lt;/font&gt;&lt;/li&gt;&lt;li&gt;&lt;a href="https://news.google.com/__i/rss/rd/articles/CBMikgFodHRwczovL3d3dy53YXNoaW5ndG9ucG9zdC5jb20vb3BpbmlvbnMvdGhlLWJvbWJzaGVsbC1jb25zZXF1ZW5jZXMtb2YtYW15LWNvbmV5LWJhcnJldHQvMjAyMC8wOS8yNS8zNTMxYWI5Yy1mZjZmLTExZWEtOGQwNS05YmVhYWE5MWM3MWZfc3RvcnkuaHRtbNIBoQFodHRwczovL3d3dy53YXNoaW5ndG9ucG9zdC5jb20vb3BpbmlvbnMvdGhlLWJvbWJzaGVsbC1jb25zZXF1ZW5jZXMtb2YtYW15LWNvbmV5LWJhcnJldHQvMjAyMC8wOS8yNS8zNTMxYWI5Yy1mZjZmLTExZWEtOGQwNS05YmVhYWE5MWM3MWZfc3RvcnkuaHRtbD9vdXRwdXRUeXBlPWFtcA?oc=5" target="_blank"&gt;The bombshell consequences of Amy Coney Barrett&lt;/a&gt;&amp;nbsp;&amp;nbsp;&lt;font color="#6f6f6f"&gt;The Washington Post&lt;/font&gt;&lt;/li&gt;&lt;li&gt;&lt;strong&gt;&lt;a href="https://news.google.com/stories/CAAqfQgKIndDQklTVXpvSmMzUnZjbmt0TXpZd1NrWUtFUWlGb2N2bmtJQU1FWGxwTTJtRVFsN01FakZTWlhCdmNuUTZJRlJ5ZFcxd0lIUnZJRzV2YldsdVlYUmxJRUZ0ZVNCRGIyNWxlU0JDWVhKeVpYUjBJSFJ2SUZOREtBQVAB?oc=5" target="_blank"&gt;View Full Coverage on Google News&lt;/a&gt;&lt;/strong&gt;&lt;/li&gt;&lt;/ol&gt;</description>
		<source url="https://www.foxnews.com">Fox News</source>
	</item>
	</channel>
	</rss>`

	rss := &rss{}
	err := xml.Unmarshal([]byte(data), rss)
	if err != nil {
		t.Errorf("Unmarshal error: %s", err.Error())
		return
	}

	rssItems := &rssItems{}
	rssItems.addFromRss(rss)
	rssItemSlice := rssItems.getItems()

	if rssItemSlice[0].Title == "" {
		t.Errorf("Item Title missing")
	}

	if rssItemSlice[0].Source == "" {
		t.Errorf("Item Source missing")
	}

	if rssItemSlice[0].SourceURL == "" {
		t.Errorf("Item SourceURL missing")
	}

	if rssItemSlice[0].Link == "" {
		t.Errorf("Item Link missing")
	}

	if rssItemSlice[0].Description == "" {
		t.Errorf("Item Description missing")
	}
}

func TestAddFromRssPartial(t *testing.T) {
	data :=
	`<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
	<rss version="2.0" xmlns:media="http://search.yahoo.com/mrss/">
	<channel>
	<generator>NFE/5.0</generator>
	<title>Top stories - Google News</title>
	<link>https://news.google.com/?hl=en-US&amp;gl=US&amp;ceid=US:en</link>
	<language>en-US</language><webMaster>news-webmaster@google.com</webMaster>
	<copyright>2020 Google Inc.</copyright>
	<lastBuildDate>Sat, 26 Sep 2020 08:50:05 -0700</lastBuildDate>
	<description>Google News</description>
	<item>
		<title>Maher rips Barrett, Trump's expected Supreme Court pick: 'She's a f---ing nut' - Fox News</title>
		<link>https://news.google.com/__i/rss/rd/articles/CBMibGh0dHBzOi8vd3d3LmZveG5ld3MuY29tL2VudGVydGFpbm1lbnQvbWFoZXItcmlwcy1iYXJyZXR0LXRydW1wcy1leHBlY3RlZC1zdXByZW1lLWNvdXJ0LXBpY2stc2hlcy1hLWYtaW5nLW51dNIBcGh0dHBzOi8vd3d3LmZveG5ld3MuY29tL2VudGVydGFpbm1lbnQvbWFoZXItcmlwcy1iYXJyZXR0LXRydW1wcy1leHBlY3RlZC1zdXByZW1lLWNvdXJ0LXBpY2stc2hlcy1hLWYtaW5nLW51dC5hbXA?oc=5</link>
		<guid isPermaLink="false">52781070340229</guid>
		<pubDate>Sat, 26 Sep 2020 05:04:18 -0700</pubDate>
		<description>&lt;ol&gt;&lt;li&gt;&lt;a href="https://news.google.com/__i/rss/rd/articles/CBMibGh0dHBzOi8vd3d3LmZveG5ld3MuY29tL2VudGVydGFpbm1lbnQvbWFoZXItcmlwcy1iYXJyZXR0LXRydW1wcy1leHBlY3RlZC1zdXByZW1lLWNvdXJ0LXBpY2stc2hlcy1hLWYtaW5nLW51dNIBcGh0dHBzOi8vd3d3LmZveG5ld3MuY29tL2VudGVydGFpbm1lbnQvbWFoZXItcmlwcy1iYXJyZXR0LXRydW1wcy1leHBlY3RlZC1zdXByZW1lLWNvdXJ0LXBpY2stc2hlcy1hLWYtaW5nLW51dC5hbXA?oc=5" target="_blank"&gt;Maher rips Barrett, Trump's expected Supreme Court pick: 'She's a f---ing nut'&lt;/a&gt;&amp;nbsp;&amp;nbsp;&lt;font color="#6f6f6f"&gt;Fox News&lt;/font&gt;&lt;/li&gt;&lt;li&gt;&lt;a href="https://news.google.com/__i/rss/rd/articles/CBMiK2h0dHBzOi8vd3d3LnlvdXR1YmUuY29tL3dhdGNoP3Y9U1hib2xNcnhXVk3SAQA?oc=5" target="_blank"&gt;Trump Plans to Name Barrett to Replace Ginsburg on Supreme Court&lt;/a&gt;&amp;nbsp;&amp;nbsp;&lt;font color="#6f6f6f"&gt;Bloomberg Politics&lt;/font&gt;&lt;/li&gt;&lt;li&gt;&lt;a href="https://news.google.com/__i/rss/rd/articles/CBMiTmh0dHBzOi8vd3d3LmZveG5ld3MuY29tL29waW5pb24vYW15LWNvbmV5LWJhcnJldHQtZmFpdGgtY291bnRyeS1tYWdnaWUtZ2FybmV0dNIBUmh0dHBzOi8vd3d3LmZveG5ld3MuY29tL29waW5pb24vYW15LWNvbmV5LWJhcnJldHQtZmFpdGgtY291bnRyeS1tYWdnaWUtZ2FybmV0dC5hbXA?oc=5" target="_blank"&gt;Maggie Garnett: Judge Amy Coney Barrett has shown me how to both live my faith and serve my country&lt;/a&gt;&amp;nbsp;&amp;nbsp;&lt;font color="#6f6f6f"&gt;Fox News&lt;/font&gt;&lt;/li&gt;&lt;li&gt;&lt;a href="https://news.google.com/__i/rss/rd/articles/CBMiSGh0dHBzOi8vd3d3Lm55dGltZXMuY29tLzIwMjAvMDkvMjUvb3Bpbmlvbi9hbXktY29uZXktYmFycmV0dC1zY290dXMuaHRtbNIBTGh0dHBzOi8vd3d3Lm55dGltZXMuY29tLzIwMjAvMDkvMjUvb3Bpbmlvbi9hbXktY29uZXktYmFycmV0dC1zY290dXMuYW1wLmh0bWw?oc=5" target="_blank"&gt;Would Amy Coney Barrett on the Supreme Court Cost the G.O.P. the Senate?&lt;/a&gt;&amp;nbsp;&amp;nbsp;&lt;font color="#6f6f6f"&gt;The New York Times&lt;/font&gt;&lt;/li&gt;&lt;li&gt;&lt;a href="https://news.google.com/__i/rss/rd/articles/CBMikgFodHRwczovL3d3dy53YXNoaW5ndG9ucG9zdC5jb20vb3BpbmlvbnMvdGhlLWJvbWJzaGVsbC1jb25zZXF1ZW5jZXMtb2YtYW15LWNvbmV5LWJhcnJldHQvMjAyMC8wOS8yNS8zNTMxYWI5Yy1mZjZmLTExZWEtOGQwNS05YmVhYWE5MWM3MWZfc3RvcnkuaHRtbNIBoQFodHRwczovL3d3dy53YXNoaW5ndG9ucG9zdC5jb20vb3BpbmlvbnMvdGhlLWJvbWJzaGVsbC1jb25zZXF1ZW5jZXMtb2YtYW15LWNvbmV5LWJhcnJldHQvMjAyMC8wOS8yNS8zNTMxYWI5Yy1mZjZmLTExZWEtOGQwNS05YmVhYWE5MWM3MWZfc3RvcnkuaHRtbD9vdXRwdXRUeXBlPWFtcA?oc=5" target="_blank"&gt;The bombshell consequences of Amy Coney Barrett&lt;/a&gt;&amp;nbsp;&amp;nbsp;&lt;font color="#6f6f6f"&gt;The Washington Post&lt;/font&gt;&lt;/li&gt;&lt;li&gt;&lt;strong&gt;&lt;a href="https://news.google.com/stories/CAAqfQgKIndDQklTVXpvSmMzUnZjbmt0TXpZd1NrWUtFUWlGb2N2bmtJQU1FWGxwTTJtRVFsN01FakZTWlhCdmNuUTZJRlJ5ZFcxd0lIUnZJRzV2YldsdVlYUmxJRUZ0ZVNCRGIyNWxlU0JDWVhKeVpYUjBJSFJ2SUZOREtBQVAB?oc=5" target="_blank"&gt;View Full Coverage on Google News&lt;/a&gt;&lt;/strong&gt;&lt;/li&gt;&lt;/ol&gt;</description>
	</item>
	</channel>
	</rss>`

	rss := &rss{}
	err := xml.Unmarshal([]byte(data), rss)
	if err != nil {
		t.Errorf("Unmarshal error: %s", err.Error())
		return
	}

	rssItems := &rssItems{}
	rssItems.addFromRss(rss)
	rssItemSlice := rssItems.getItems()

	if rssItemSlice[0].Title == "" {
		t.Errorf("Item Title missing")
	}

	if rssItemSlice[0].Source == "" {
		t.Errorf("Item Source missing")
	}

	if rssItemSlice[0].SourceURL == "" {
		t.Errorf("Item SourceURL missing")
	}

	if rssItemSlice[0].Link == "" {
		t.Errorf("Item Link missing")
	}

	if rssItemSlice[0].Description == "" {
		t.Errorf("Item Description missing")
	}
}
