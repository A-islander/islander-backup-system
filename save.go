package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type ForumPost struct {
	Id            int
	Title         string
	Value         string
	FollowId      int
	PlateId       int
	Status        int
	ReplyArr      string
	UserId        int
	Time          int
	MediaUrl      string
	ReplyCount    int
	TopStatus     int
	LastReplyTime int
	SageAddId     string
	SageSubId     string
	Name          string
}

// 写入文件
func saveToJson(filePath string, data []ForumPost) {
	file, err := os.OpenFile("./save/"+filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	str, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	write.WriteString(string(str))
	write.Flush()
}

// 创建文件ming
func createFileName(times int, timestamp int64) string {
	var res string
	timeObj := time.Unix(timestamp, 0)
	year := timeObj.Year()   //年
	month := timeObj.Month() //月
	day := timeObj.Day()     //日
	hour := timeObj.Hour()   //小时
	res = fmt.Sprintf("%d_%02d_%02d_%02d_%d.json", year, month, day, hour, times)
	return res
}
