package testTimer

import (
	"fmt"
	"time"
)

func TestTimer() {
	fmt.Println("TestTimer ---------------------------", time.Now())
	time.AfterFunc(time.Second*5, timefunc1)
	time.Sleep(time.Second * 10)
	fmt.Println("final -------------------------------", time.Now())

}

func timefunc1() {
	fmt.Println("timefunc1", time.Now())
}
