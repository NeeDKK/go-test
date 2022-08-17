package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"test/initialize"
)

type BaseLineIpPort struct {
	ID              uint   `gorm:"primarykey" json:"id"` // 主键ID
	Ip              string `json:"ip"`
	Port            string `json:"port"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	About           int    `json:"about"`
	BaseLineGroupId int    `json:"baseLineGroupId"`
}

func (p *BaseLineIpPort) TableName() string {
	return "base_line_ip_port"
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "10.25.10.125:6379",
		Password: "Netvine123#@!", // no password set
		DB:       1,               // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err)
	} else {
		var BaseLineIpPort []BaseLineIpPort
		m := make(map[string]string)
		//存入IP端口的基线白名单
		db := initialize.GormMysql()
		db.Model(&BaseLineIpPort).Where("base_line_group_id = 4").Find(&BaseLineIpPort)
		for _, v := range BaseLineIpPort {
			if v.About == 1 {
				m[v.Ip+"-"+strconv.Itoa(int(v.ID))] = "IP"
			} else {
				m[v.Ip+"-"+v.Port] = "IP-PORT"
			}
		}
		client.HSet(context.Background(), "base_line_warning", m)
	}
}
