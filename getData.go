package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type config struct {
	UserName string
	PassWord string
	Ip       string
	Database string
}

// 获取config配置
func getConfig() config {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	var res config
	json.Unmarshal(file, &res)
	return res
}

// 新建db连接
func newDB() *gorm.DB {
	config := getConfig()
	dsn := config.UserName + ":" + config.PassWord + "@tcp(" + config.Ip + ")/" + config.Database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	return db
}

func getData() {
	var size int64 = 100
	var count int64
	db := newDB()
	time := getHourTimestamp()
	// 获取每个小时前六十分钟的串
	db.Model(&ForumPost{}).Where("time >= ? and time < ?", time-3600, time).Count(&count)
	// db.Model(&ForumPost{}).Count(&count)
	for i := 0; count > 0; i++ {
		var res []ForumPost
		db.Limit(int(size)).Offset(i*int(size)).Order("time asc").Where("time >= ? and time < ?", time-3600, time).Find(&res)
		// db.Limit(int(size)).Offset(i * int(size)).Order("time asc").Find(&res)
		count -= size
		saveToJson(createFileName(i, time), res)
	}
}

// 获取整点
func getHourTimestamp() int64 {
	now := time.Now()
	hour := now.Unix() - int64(now.Second()) - int64(60*now.Minute())
	return hour
}

func (ForumPost) TableName() string {
	return "forum_post"
}
