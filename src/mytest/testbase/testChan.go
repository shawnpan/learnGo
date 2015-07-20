package testbase

import (
	"fmt"
)

func TestChan() {
	ic := make(chan int)
	go func() {
		ic <- 10
	}()

	num := <-ic
	fmt.Println(num)
}
