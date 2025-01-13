package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"report-backend-golang/global"
	"time"
)

func WriteErrorLog(c *gin.Context, msg string) {
	fileName := fmt.Sprintf("%s/apiError_%s.log",global.EnvConfig.Files.LogPath, time.Now().Format("200601"))

	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("error opening file: %v", err)
	}

	defer f.Close()
	log.SetOutput(f)

	logmsg := fmt.Sprintf("Method=\"%s\" URL=\"%s\" msg=\"%s\"", c.Request.Method, c.Request.RequestURI, msg)
	log.Println(logmsg)
}
