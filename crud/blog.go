package crud

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"terminal-blog-api/model"
	"time"
)

func InsertBlog(blog model.Blog, db *gorm.DB) {
	blog.Bid = uuid.New().String()
	blog.Time = time.Now().Format(time.UnixDate)
	db.Create(blog)
}

func DeleteBlog(bid string, db *gorm.DB) {
	db.Where("bid = ?", bid).Delete(&model.Blog{})
}

func GetBlog(bid string, db *gorm.DB) model.Blog {
	var blog model.Blog
	db.Find(&blog, "bid = ?", bid)
	return blog
}

func ListBlog(db *gorm.DB) []model.Blog {
	var blogs []model.Blog
	db.Select([]string{"bid", "time", "title"}).Find(blogs)
	return blogs
}
