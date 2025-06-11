package log

import (
	"log"
	"os"
	"report-backend-golang/global"
	//"xdr/utils"
	"time"
    "fmt"
)

// func Logrecord1(title,msg string) string{

//     // open file and create if non-existent
//     file, err := os.OpenFile( global.EnvConfig.Files.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
//     if err != nil {
//         log.Fatal(err)
//     }
//     defer file.Close()

//     logger := log.New(file, title + " ", log.LstdFlags)
//     logger.Println(msg)
// 	return msg
//     //time.Sleep(5 * time.Second)
//     //logger.Println("A new log, 5 seconds later")
// }


func Logrecord(title,msg string) string{

    fileName := fmt.Sprintf("%s/ReportEngine-%s.log", global.EnvConfig.Files.LogPath, time.Now().Format("200601"))    
    // open file and create if non-existent
    file, err := os.OpenFile( fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    logger := log.New(file, title + " ", log.LstdFlags)
    logger.Println(msg)
	return msg

}