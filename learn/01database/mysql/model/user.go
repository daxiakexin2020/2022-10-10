package model

import (
	"fmt"
	"time"
)

type User struct {
	Id         int       `json:"id" gorm:"id"`
	Name       string    `json:"name" gorm:"name"`
	Age        int       `json:"age" gorm:"age"`
	Sex        int       `json:"sex" gorm:"sex"`
	CreateTime time.Time `json:"create_time" gorm:"create_time"`
	UpdateTime time.Time `json:"update_time" gorm:"update_time"`
}

func (u *User) Get() *User {
	user := User{}
	TestDB.Table("users").Find(&user)
	return &user
}

func (u *User) Update() {
	TestDB.Model(&User{}).Where("name=?", "test1").Update("name", "test2")
}

func (u *User) Delete() {

}

func (u *User) Create() {
	user := &User{
		Name:       "test1",
		Age:        20,
		Sex:        1,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	ret := TestDB.Table("users").Create(user)
	fmt.Print(ret.Error)
}
