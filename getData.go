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
	UserName       string
	PassWord       string
	Ip             string
	Database       string
	LastRecordTime int
}

var db = newDB()

// 获取config配置
func getConfig() config {
	file, err := ioutil.ReadFile("./conf/config.json")
	if err != nil {
		log.Fatal(err)
	}
	var res config
	json.Unmarshal(file, &res)
	return res
}

func updateLastRecordTime(nowTime int64) {
	file, err := ioutil.ReadFile("./conf/config.json")
	if err != nil {
		log.Fatal(err)
	}
	var res config
	json.Unmarshal(file, &res)
	res.LastRecordTime = int(nowTime)
	str, err := json.Marshal(res)
	ioutil.WriteFile("./conf/config.json", str, 0666)
}

// 新建db连接
func newDB() *gorm.DB {
	config := getConfig()
	dsn := config.UserName + ":" + config.PassWord + "@tcp(" + config.Ip + ")/" + config.Database + "?timeout=30s"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	return db
}

func getData(timeStamp int64) error {
	var size int64 = 100
	var count int64
	nowTime := getHourTimestamp(timeStamp)
	// 获取每个小时前六十分钟的串
	db.Model(&ForumPost{}).Where("time >= ? and time < ?", nowTime-3600, nowTime).Count(&count)
	// db.Model(&ForumPost{}).Count(&count)
	for i := 0; count > 0; i++ {
		var res []ForumPost
		db.Limit(int(size)).Offset(i*int(size)).Order("time asc").Where("time >= ? and time < ?", nowTime-3600, nowTime).Find(&res)
		// db.Limit(int(size)).Offset(i * int(size)).Order("time asc").Find(&res)
		count -= size
		saveToJson(createFileName(i, nowTime), res)
	}
	updateLastRecordTime(nowTime)
	return nil
}

// 获取整点
func getHourTimestamp(timeStamp int64) int64 {
	now := time.Unix(timeStamp, 0)
	hour := now.Unix() - int64(now.Second()) - int64(60*now.Minute())
	return hour
}

func (ForumPost) TableName() string {
	return "forum_post"
}
