package main

import (
	"encoding/json"
	"fmt"
	"hoanglv/crawler/model"
	"io/ioutil"
	"strings"

	// "strings"
	// "time"

	// "encoding/json"
	"github.com/gocolly/colly"
)
var data []model.Article
var filename string

func crawl(path string, topic string){
	c := colly.NewCollector()

	// domain := "https://vnexpress.net/"

	c.OnHTML(".col-left .title-news a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnHTML(".sidebar-1", func(e *colly.HTMLElement) {
		article := &model.Article{}

		artTime := e.ChildText(".header-content span.date")
		strArr := strings.Split(artTime, ",")
		article.CreatedDate = strings.ReplaceAll(strArr[0], " ", "")

		article.Avatar = e.ChildAttr(".fig-picture .lazy", "data-src")
		article.FeatureImage = e.ChildAttr(".fig-picture img", "data-src")
		article.Content = make([]model.Content, 0)
		
		e.ForEach("p.Normal", func(i int, h *colly.HTMLElement) {
			article.Content = append(article.Content, model.Content{Content: h.Text, ContentType: "text"})
		})

		article.Sapo = e.ChildText("p.description")
		article.Topic = e.ChildText(".title-detail")
		article.Language = "vietnamese"
		article.Newspaper ="VnExpress"
		article.Href = e.Request.URL.String();

		data = append(data, *article)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL)
	})
	c.Visit(fmt.Sprintf("https://vnexpress.net/%s", path))
	
	for i :=2; i<3 ; i++ {
		url := fmt.Sprintf("https://vnexpress.net/%s/p%d", path, i)
		c.Visit(url)
	}
}


func main() {

	filename = "vtv/vtv.json"
	data = make([]model.Article, 0)
	crawl("the-thao", "Thể thao")
	crawl("the-gioi", "The gioi")
	crawl("kinh-doanh", "Kinh doanh")
	crawl("phap-luat", "Phap luat")
	crawl("giao-duc", "Giao duc")
	crawl("suc-khoe", "Suc khoe")
	crawl("khoa-hoc", "Khoa hoc")
	crawl("so-hoa", "SO hoa")

	result, _ := json.Marshal(data)
	_ = ioutil.WriteFile(filename, result, 0644)
}
