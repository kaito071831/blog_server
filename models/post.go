package models

import (
	"time"

	"github.com/kaito071831/blog_server/db"
)

type Post struct {
	ID int `gorm:"primarykey" json:"id"`
	Title string `gorm:"type:string;not null" json:"title"`
	Body string `gorm:"type:string" json:"body"`
	Created_at time.Time `gorm:"type:timestamp;not null;autoCreateTime" json:"created_at"`
	Updated_at time.Time `gorm:"type:timestamp;not null;autoUpdateTime" json:"updated_at"`
	UserID int `json:"user_id"`
}

func init() {
	db.Db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&Post{})
}

func GetPost(id int)(*Post, error) {
	post := Post{}
	db.Db.First(&post, id)
	return &post, nil
}

func GetPosts()([]*Post, error) {
	posts := []*Post{}
	db.Db.Find(&posts)
	if posts == nil {
		panic("記事一覧の取得に失敗しました")
	}
	return posts, nil
}

func CreatePost(p Post)(Post, error){
	post := p
	if err := db.Db.Create(&post).Error; err != nil {
		return Post{}, nil
	}
	return post, nil
}

func UpdatePost(p Post)(Post, error){
	post := Post{}
	if err := db.Db.First(&post, p.ID).Error; err != nil {
		return Post{}, err
	}
	if err := db.Db.Model(&post).Updates(p).Error; err != nil {
		return Post{}, err
	}
	return post, nil
}
