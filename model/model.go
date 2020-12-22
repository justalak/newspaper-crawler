package model

import "time"

type Content struct {
	Content string `json:"content"`;
	ContentType string `json:"type"`;
	Link string `json:"link"`;
}

type Article struct {
	Avatar string `json:"avatar"`;
	Content []Content `json:"content"`;
	CreatedDate time.Time `json:"created_date"`;
	FeatureImage string `json:"feature_image"`;
	Href string `json:"href"`;
	Newspaper string `json:"newspaper"`;
	Language string `json:"language"`;
	Sapo string `json:"sapo"`;
	Topic string `json:"topic"`;
}
