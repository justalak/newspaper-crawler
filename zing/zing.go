package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"hoanglv/crawler/model"
	// "encoding/json"
	"github.com/gocolly/colly"
)

var data []model.Article
var filename string

func crawlDanTri(path string, topic string){
	c := colly.NewCollector()

	domain := "https://dantri.com.vn"

	c.OnHTML(".dt-main-category .news-item__title a", func(e *colly.HTMLElement) {
		e.Request.Visit(domain + e.Attr("href"))
	})

	c.OnHTML(".dt-news__detail", func(e *colly.HTMLElement) {
		article := &model.Article{}

		artTime := e.ChildText(".dt-news__time")
		strArr := strings.Split(artTime, ",")
		article.CreatedDate= strings.ReplaceAll(strings.Split(strArr[1], "-")[0], " ", "")

		article.Avatar = e.ChildAttr("img", "data-original")
		article.Content = make([]model.Content, 0)
		
		e.ForEach("p", func(i int, h *colly.HTMLElement) {
			article.Content = append(article.Content, model.Content{Content: h.Text, ContentType: "text"})
		})

		article.Sapo = e.ChildText(".dt-news__sapo")
		article.Topic = e.ChildText(".dt-news__title")
		article.Language = "vietnamese"
		article.Newspaper ="Dan tri"
		article.Href = e.Request.URL.String();

		data = append(data, *article)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL)
	})

	for i :=1; i<2 ; i++ {
		url := fmt.Sprintf("https://dantri.com.vn/%s/trang-%d.htm", path, i)
		c.Visit(url)
	}
}


func main() {

	filename = "zing/zing.json"
	data = make([]model.Article, 0)
	crawlDanTri("the-thao", "Thể thao")
	crawlDanTri("su-kien", "Sự kiện")
	crawlDanTri("the-gioi", "Thế giới")
	crawlDanTri("xa-hoi", "Thế giới")
	crawlDanTri("suc-khoe", "Suc khoe")
	crawlDanTri("kinh-doanh", "Kinh doanh")
	// crawlVnExpress()
	result, _ := json.Marshal(data)

	_ = ioutil.WriteFile(filename, result, 0644)
}
