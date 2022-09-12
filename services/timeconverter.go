package services

import (
	"fmt"
	// "strconv"
	// "strings"
	"time"
	"github.com/timberio/go-datemath"
)


func Timeconverter(anchorDate string)(string){
	now := time.Now()
	t,err := datemath.Parse(anchorDate)
	if err != nil {
		panic(err)
	}
	// expr, _ := datemath.Parse("now-15m")

	a := fmt.Sprintln(t.Time(datemath.WithNow(now)))
	// time := t.Time(datemath.WithNow(now))
	return a
}