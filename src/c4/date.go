package c4

import (
	"fmt"
	"time"
)

// DateAndTime 时间和日期
func DateAndTime() {

	nowTime := time.Now()
	fmt.Println(nowTime)


	fmt.Printf("%d-%02d-%d %d:%d:%d \n",
		nowTime.Year(),
		nowTime.Month(),
		nowTime.Day(),
		nowTime.Hour(),
		nowTime.Minute(),
		nowTime.Second())

	//一周后时间 week单位为纳秒
	var week time.Duration
	week = 60 * 60 * 24 * 7 * 1e9
	nextWeek := nowTime.Add(week)
	fmt.Println(nextWeek)

	fmt.Println(nowTime.Format("2006-01-02 15:04:05"))

}