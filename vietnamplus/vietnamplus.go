package main

import (
	"encoding/json"
	"fmt"
	"hoanglv/crawler/model"
	"io/ioutil"

	// "strings"
	// "time"

	// "encoding/json"
	"github.com/gocolly/colly"
)
var data []model.Article
var filename string

func crawl(path string, topic string){
	c := colly.NewCollector()

	domain := "https://www.vietnamplus.vn"

	c.OnHTML(".zone--timeline .story .story__title", func(e *colly.HTMLElement) {
		fmt.Println(e.Attr("href"))
		e.Request.Visit(domain + e.Attr("href"))
	})

	c.OnHTML(".details", func(e *colly.HTMLElement) {
		article := &model.Article{}

		article.CreatedDate = e.ChildAttr(".details__meta .cms-date", "content")

		article.Avatar = e.ChildAttr(".cms-photo", "data-photo-original-src")
		article.FeatureImage = e.ChildAttr("img.cms-photo", "data-photo-original-src")
		article.Content = make([]model.Content, 0)
		
		e.ForEach(".details__content .article-body p", func(i int, h *colly.HTMLElement) {
			article.Content = append(article.Content, model.Content{Content: h.Text, ContentType: "text"})
		})

		article.Sapo = e.ChildText(".details__summary")
		article.Topic = e.ChildText("h1.details__headline")
		article.Language = "vietnamese"
		article.Newspaper ="Vietnamplus"
		article.Href = e.Request.URL.String();

		data = append(data, *article)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL)
	})
	
	for i :=2; i<3 ; i++ {
		url := fmt.Sprintf("https://www.vietnamplus.vn/%s/trang%d.vnp", path, i)
		c.Visit(url)
	}
}


func main() {

	filename = "vietnamplus/vietnamplus.json"
	data = make([]model.Article, 0)
	crawl("thethao", "Thể thao")
	crawl("thegioi", "The gioi")
	crawl("kinhte", "Kinh te")
	// crawl("phap-luat", "Phap luat")
	// crawl("giao-duc", "Giao duc")
	// crawl("suc-khoe", "Suc khoe")
	crawl("congnghe", "Cong nghe")
	// crawl("thoi-su", "SO hoa")

	result, _ := json.Marshal(data)
	_ = ioutil.WriteFile(filename, result, 0644)
}
