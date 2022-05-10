package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/robfig/cron/v3"
)

func init() {
	file := "./log/" + "message" + ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[CronSave]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)

	initItem()
	return
}

func main() {
	crons := cron.New()
	_, err := crons.AddFunc("@every 1h", mission)
	if err != nil {
		panic(err)
	}
	crons.Start()
	defer crons.Stop()

	select {}
}

func mission() {
	err := getData(time.Now().Unix())
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Save")
}

func initItem() {
	conf := getConfig()
	recordTime := int64(conf.LastRecordTime)
	// 往回倒一天
	recordTime = recordTime - 86400
	nowTime := time.Now().Unix()
	for {
		if recordTime > nowTime {
			break
		}
		timeLayout := "2006-01-02 15:04:05"
		datetime := time.Unix(recordTime, 0).Format(timeLayout)
		fmt.Println(datetime)
		getData(recordTime)
		recordTime += 3600
	}
}
