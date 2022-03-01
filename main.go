package main

import (
	"log"
	"os"

	"github.com/robfig/cron/v3"
)

func init() {
	file := "./" + "message" + ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[CronSave]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
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
	getData()
	log.Println("Save")
}
