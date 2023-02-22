package services


import (
	// "fmt"
	// "report-backend-golang/entities"
	"report-backend-golang/global"
	// "report-backend-golang/models"
	"fmt"
	"time"
)


type EntryID int

type Entry struct {
	ID EntryID
	// Schedule Schedule
	Next time.Time
	Prev time.Time
	// WrappedJob Job
	// Job Job
}

func Print() {
	fmt.Println(time.Now().String()+"  hello")
}

func Print1() {
	fmt.Println(time.Now().String()+"   hi你媽")
}
func CronTest(){

	EntryID,err := global.Crontab.AddFunc("*/1 * * * *",func(){
		Print()
	})
	fmt.Println(EntryID,err)

	EntryID1,err := global.Crontab.AddFunc("*/1 * * * *",func(){
		Print1()
	}) 
	fmt.Println(EntryID1,err)
	global.Crontab.Entries()
	fmt.Println(global.Crontab.Entries())
	// global.Crontab.Remove(EntryID)
	if err != nil {
		fmt.Println("crontab test 初始化失敗")
		// log.Logrecord("排程 ","PDF排程 初始化失敗")
		fmt.Println(err.Error())
		// log.Logrecord("ERROR ",err.Error())
	} else {
		fmt.Println("crontab test 初始化成功")
		// log.Logrecord("排程 ","PDF排程 初始化成功")
		// c.Start()
		global.Crontab.Start()
		time.Sleep(time.Minute * 30)
	}
}