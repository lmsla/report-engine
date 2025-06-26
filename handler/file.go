package handler

import (
	"fmt"
	"os"
	"report-backend-golang/global"
	"time"
	"log"
)


func WriteFrontendLog(log_type, usr, msg string) error {
	fileName := fmt.Sprintf("%s/apiAuth_%s.log", global.EnvConfig.Files.LogPath, time.Now().Format("2006_01_02"))

	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	defer f.Close()
	log.SetOutput(f)

	log.Println(fmt.Sprintf("[%s]", log_type), usr, msg)
	return err
}