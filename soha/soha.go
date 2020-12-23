package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"hoanglv/crawler/model"
	// "encoding/json"
	"github.com/gocolly/colly"
)

var data []model.Article
var filename string

func crawl(path string, topic string){
	c := colly.NewCollector()

	domain := "https://soha.vn"

	c.OnHTML(".info-new-cate a", func(e *colly.HTMLElement) {
		e.Request.Visit(domain + e.Attr("href"))
	})

	c.OnHTML(".news-detail", func(e *colly.HTMLElement) {
		article := &model.Article{}

		article.CreatedDate = e.ChildText(".op-published")

		article.Avatar = e.ChildAttr(".avatarBrand img", "src")
		article.Content = make([]model.Content, 0)
		
		e.ForEach(".news-content p", func(i int, h *colly.HTMLElement) {
			article.Content = append(article.Content, model.Content{Content: h.Text, ContentType: "text"})
		})

		article.Sapo = e.ChildText(".news-sapo")
		article.Topic = e.ChildText(".news-title")
		article.Language = "vietnamese"
		article.Newspaper ="Soha"
		article.Href = e.Request.URL.String();

		data = append(data, *article)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL)
	})

	for i :=1; i<2 ; i++ {
		url := fmt.Sprintf("https://soha.vn/timeline/%s/trang-%d.htm", path, i)
		c.Visit(url)
	}
}


func main() {

	filename = "soha/soha.json"
	data = make([]model.Article, 0)
	crawl("1001", "")
	crawl("1002", "")
	crawl("1000", "")
	// crawlVnExpress()
	result, _ := json.Marshal(data)

	_ = ioutil.WriteFile(filename, result, 0644)
}
