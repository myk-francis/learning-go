package main

type RSSFeed struct {
	Channel struct {
		Title string `xml:"title"`
		Link string `xml:"link"`
		Description string `xml:"description"`
		Language string `xml:"language"`
		Item string `xml:"item"`
	}
}