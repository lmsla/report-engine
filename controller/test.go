package controller

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

func TestDefer(c *gin.Context) {
	if err := test(); err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(200, nil)
	}
}

func test() (err error) {
	defer func() {
		if panicError := recover(); panicError != nil {
			// 处理错误
			fmt.Println("crash", "current params")
			
			err = errors.New("panic error")
		} else {
			fmt.Println("normal error")
			err = errors.New("normal error")
		}

	}()

	panic("fff")
	// return nil
}
