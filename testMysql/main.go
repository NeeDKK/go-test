package main

import (
	"fmt"
	"gorm.io/gorm"
	"test/initialize"
)

type AllowGroup struct {
	ID           int64  `gorm:"primary_key"`
	GroupName    string `json:"groupName"`
	Description  string `json:"description"`
	Status       int64  `json:"status"`
	ProtoGroupId int    `json:"protoGroupId"`
}

func (a *AllowGroup) TableName() string {
	return "allow_group"
}

func main() {
	db := initialize.GormMysql()
	list := []AllowGroup{
		{GroupName: "test1", Description: "", Status: 1, ProtoGroupId: 1},
		{GroupName: "test2", Description: "", Status: 1, ProtoGroupId: 1},
		{GroupName: "test3", Description: "", Status: 1, ProtoGroupId: 1},
		{GroupName: "test4", Description: "", Status: 1, ProtoGroupId: 1},
		{GroupName: "test5", Description: "", Status: 1, ProtoGroupId: 1},
	}
	db.Create(list)
}

func (a *AllowGroup) AfterSave(tx *gorm.DB) (err error) {
	fmt.Println(a.ID)
	return
}
