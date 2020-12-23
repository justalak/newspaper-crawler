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

	c.OnHTML(".list-content h3 a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnHTML(".ArticleDetail", func(e *colly.HTMLElement) {
		article := &model.Article{}

		artTime := e.ChildText(".ArticleDate")
		strArr := strings.Split(artTime, ",")
		article.CreatedDate= strings.ReplaceAll(strArr[0], "\n", " ")

		article.Avatar = e.ChildAttr(".ImageCenterBox img", "src")
		article.FeatureImage = e.ChildAttr(".FmsArticleBoxStyle-Images img", "src")
		article.Content = make([]model.Content, 0)
		
		e.ForEach(".ArticleContent p.t-j", func(i int, h *colly.HTMLElement) {
			article.Content = append(article.Content, model.Content{Content: h.Text, ContentType: "text"})
		})

		article.Sapo = e.ChildText(".ArticleLead p")
		article.Topic = e.ChildText("h1.title")
		article.Language = "vietnamese"
		article.Newspaper ="Vietnamnet"
		article.Href = e.Request.URL.String();

		data = append(data, *article)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL)
	})
	// c.Visit(fmt.Sprintf("https://vietnamnet.vn/vn/%s", path))
	
	for i :=2; i<3 ; i++ {
		url := fmt.Sprintf("https://vietnamnet.vn/vn/%s/trang%d", path, i)
		c.Visit(url)
	}
}


func main() {

	filename = "vietnamnet/vietnamnet.json"
	data = make([]model.Article, 0)
	crawl("the-thao", "Thể thao")
	crawl("the-gioi", "The gioi")
	crawl("kinh-doanh", "Kinh doanh")
	crawl("phap-luat", "Phap luat")
	crawl("giao-duc", "Giao duc")
	crawl("suc-khoe", "Suc khoe")
	crawl("cong-nghe", "Khoa hoc")
	crawl("thoi-su", "SO hoa")

	result, _ := json.Marshal(data)
	_ = ioutil.WriteFile(filename, result, 0644)
}
