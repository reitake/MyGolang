package main

import (
	"fmt"
	"time"
)


/*// 用法：
func main() {
	clockch := make(chan time.Time)
	go SetClock(clockch)
	select {
	case <-clockch:
		dosomething()
	}
}*/

func SetClock(clockch chan time.Time) {
	now := time.Now()

	// 如果当前时间小于当天的闹钟时间，则定下今天的闹钟
	// 如果当前时间大于当天的闹钟时间，则定下明天的闹钟
	if now.Hour()<12 || (now.Hour()=12&&now.Minute()<0){
		next:=now
	}else{
		next := now.Add(time.Hour * 24)	
	}
	
	// 设定定时任务的时间，四个填入值是：时、分、秒、纳秒
	next = time.Date(next.Year(), next.Month(), next.Day(), 12, 00, 0, 0, next.Location())
	t := time.NewTimer(next.Sub(now))
	clockch <- <-t.C
	return
}


/*func dosomething() {
	now := time.Now()
	fmt.Println("DO SOMETHING AT:", now)
}*/
