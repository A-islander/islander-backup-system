package main

import (
	"fmt"
	"testing"
)

// func TestSave(t *testing.T) {
// 	for i := 0; i < 5; i++ {
// 		saveToJson(createFileName(i), []ForumPost{{}, {}})
// 	}
// }

func TestGet(t *testing.T) {
	runCmd("./save", "git", "status")
	runCmd("./save", "git", "status")
	fmt.Println(nowTime())
	// runCmd("./save", "git", "add", "-A")
	// runCmd("./save", "git", "commit", "-m", "\""+"\"")
}

func TestInit(t *testing.T) {
	initItem()
}

func TestUpdateTime(t *testing.T) {
	// now := getHourTimestamp(time.Now().Unix())
	// updateLastRecordTime(now)
}
