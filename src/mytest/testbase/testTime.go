package testbase

import (
	"fmt"
	"time"
)

func TestTime() {
	fmt.Println("testTime-----------------------")
	//	testTick()
	t := time.Date(2111, 11, 11, 15, 15, 15, 15, time.Local)
	fmt.Println(t)

	t1 := time.Now()
	fmt.Println("t1=", t1)
	fmt.Println("t1=", t1.String())

	fromtStr := t1.Format(time.RubyDate)
	fmt.Println("fromtStr", fromtStr)

	t2, err := time.Parse(time.RubyDate, "Mon Jun 29 17:29:10 +0800 2015")
	if err != nil {
		fmt.Println("parse time err", err)
	}
	fmt.Println("t2=", t2)

	y, m, d := t2.Date()
	fmt.Printf("%d-%d-%d", y, m, d)
	fmt.Println("")

	y, m, d = t2.Year(), t2.Month(), t2.Day()
	h, mi, s := t2.Hour(), t2.Minute(), t2.Second()
	fmt.Printf("%d-%d-%d %d:%d:%d", y, m, d, h, mi, s)
	fmt.Println("")

	time.Sleep(2 * time.Second)
}
func testTick() {
	c := time.Tick(5 * time.Second)
	for now := range c {
		fmt.Printf("%v %s\n", now, "")
	}
}
