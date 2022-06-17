package models

import (
	"log"
	"time"

	"github.com/kaito071831/blog_server/db"
)

// ユーザーの構造体
type User struct {
	ID int `gorm:"primarykey" json:"id"`
	Email string `gorm:"type:string;not null" json:"email"`
	Hash_password string `gorm:"type:string;not null" json:"hash_password"`
	Created_at time.Time `gorm:"type:timestamp;not null;autoCreateTime" json:"created_at"`
	Updated_at time.Time `gorm:"type:timestamp;not null;autoUpdateTime" json:"updated_at"`
	Posts []Post
}

// 自動でテーブル作成＆マイグレーション
func init() {
	db.Db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&User{})
}

// ユーザーを取得
func GetUser(id int)(*User, error) {
	user := User{}
	db.Db.First(&user, id)
	return &user, nil
}

// 全ユーザーを取得
func GetUsers()([]*User, error) {
	users := []*User{}
	if err := db.Db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// ユーザーを作成する
func CreateUser(u User)(User, error){
	user := u
	if err := db.Db.Create(&user).Error; err != nil {
		log.Println(err)
		return User{}, err
	}
	return user, nil
}

// ユーザーを更新する
func UpdateUser(u User)(User, error){
	user := User{}
	if err := db.Db.First(&user, u.ID).Error; err != nil {
		return User{}, err
	}
	if err := db.Db.Model(&user).Updates(u).Error; err != nil {
		return User{}, err
	}
	return user, nil
}
