package mytest

import (
	"fmt"
	"io"
	"os"
)

func TestFile() {
	fmt.Println("-testFile------------------------------------------")
	r, _ := os.Open("testMap.go")
	buf := make([]byte, 20)
	io.ReadFull(r, buf)
	fmt.Println(buf)

}
