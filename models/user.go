package models

import (
	"time"
	"github.com/kaito071831/blog_server/db"
)

type User struct {
	ID int `gorm:"primarykey" json:"id"`
	Email string `gorm:"type:string;not null" json:"email"`
	Hash_password string `gorm:"type:string;not null" json:"hash_password"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Posts []Post
}

func init() {
	db.Db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&User{})
}

func GetUser(id int)(*User, error) {
	user := User{}
	db.Db.First(&user, id)
	return &user, nil
}

func GetUsers()([]*User, error) {
	users := []*User{}
	if err := db.Db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(u User)(User, error){
	user := u
	if err := db.Db.Create(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

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
