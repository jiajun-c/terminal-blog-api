package model

import _ "gorm.io/gorm"

type Blog struct {
	Bid      string `gorm:"primaryKey"` // The id for the blog, generated by the uuid
	Time     string
	title    string
	abstract string `gorm:"type:text"` // The abstract for the article
	content  string `gorm:"type:text"` // The content for the article
}
