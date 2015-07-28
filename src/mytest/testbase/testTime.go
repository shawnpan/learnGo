package testbase

import (
	"fmt"
	"time"
)

func TestTime() {
	fmt.Println("testTime-----------------------")
	// testTick()
	t := time.Date(2111, 11, 11, 15, 15, 15, 15, time.Local)
	fmt.Println(t)

	t1 := time.Now()
	fmt.Println("t1=", t1)
	fmt.Println("t1=", t1.String())

	// 时间格式化到字符串
	timeToStr := t1.Format(time.RubyDate)
	fmt.Println("timeToStr : ", timeToStr)

	// 自定义格式化，格式请参考：format.go 中定义的stdxxx常量,切记一定要使用该常量字符串
	formatStr := "2006-01-02" + " 15:04:05"
	timeToStr = t1.Format(formatStr)
	fmt.Println("timeToStr : ", timeToStr)

	// 字符串转为time
	t2, err := time.Parse(time.RubyDate, "Mon Jun 29 17:29:10 +0800 2015")
	if err != nil {
		fmt.Println("parse time err", err)
	}
	fmt.Println("t2=", t2)

	// 字符串转为time 使用定义的格式转换
	t2, err = time.Parse(formatStr, "2015-11-02 15:44:05")
	if err != nil {
		fmt.Println("parse time err", err)
	}
	fmt.Println("t2=", t2)

	// 通过time获取年/月/日
	y, m, d := t2.Date()
	fmt.Printf("%d-%d-%d", y, m, d)
	fmt.Println("")

	y, m, d = t2.Year(), t2.Month(), t2.Day()
	h, mi, s := t2.Hour(), t2.Minute(), t2.Second()
	fmt.Printf("%d-%d-%d %d:%d:%d", y, m, d, h, mi, s)
	fmt.Println("")

	// 休眠当前goroutine
	time.Sleep(100 * time.Millisecond)

	// 求时间差
	t3 := time.Now()
	delta := t2.Sub(t1)
	fmt.Println("delta = ", delta)
	fmt.Println("t3.Sub(t1) : ", t3.Sub(t1))

	time.Sleep(time.Millisecond * 100)
	t4 := time.Now()
	fmt.Println("time use : ", t4.Nanosecond()-t3.Nanosecond())

}
func testTick() {
	c := time.Tick(5 * time.Second)
	for now := range c {
		fmt.Printf("%v %s\n", now, "")
	}
}
