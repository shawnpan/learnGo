package mytest

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func TestFile() {
	fmt.Println("-testFile------------------------------------------")
	fn := "config/config.txt"
	fmt.Println(filepath.Base(fn))

	testReadFull(fn)
	testReadAtLeast(fn)
	testIoutilReadFile(fn)

}
func testReadFull(fn string) {
	fmt.Println("testReadFull---------------------------------------")
	r, err := os.Open(fn)
	if err != nil {
		fmt.Println("open file error:", fn, err)
		return
	}
	buf := make([]byte, 20)
	defer r.Close()
	_, rerr := io.ReadFull(r, buf)
	for {
		if rerr == nil {
			// fmt.Println(buf)
			fmt.Println(string(buf))
			fmt.Println("######", len(buf))
			_, rerr = io.ReadFull(r, buf)
		} else {
			if rerr == io.EOF {
				fmt.Println(" read at the end!!!")
			} else {
				fmt.Println("read error : ", rerr)
			}
			break
		}
	}
}
func testReadAtLeast(fn string) {
	fmt.Println("testReadAtLeast---------------------------------------")
	r, err := os.Open(fn)
	if err != nil {
		fmt.Println("open file error:", fn, err)
		return
	}
	buf := make([]byte, 20)
	defer r.Close()
	_, rerr := io.ReadAtLeast(r, buf, 1)
	for {
		if rerr == nil {
			// fmt.Println(buf)
			fmt.Println(string(buf))
			fmt.Println("######", len(buf))
			_, rerr = io.ReadAtLeast(r, buf, 1)
		} else {
			if rerr == io.EOF {
				fmt.Println(" read at the end!!!")
			} else {
				fmt.Println("read error : ", rerr)
			}
			break
		}
	}
}
func testIoutilReadFile(fn string) {
	fmt.Println("testIoutilReadFile---------------------------------------")
	f, err := ioutil.ReadFile(fn)
	if err != nil {
		fmt.Println("open file error : ", fn, err)
		return
	}
	fmt.Println(string(f))
}
