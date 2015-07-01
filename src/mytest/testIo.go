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
	testIoutilReadAll(fn)
	testWriteFile(fn)

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
func testIoutilReadAll(fn string) {
	fmt.Println("testIoutilReadAll---------------------------------------")
	r, rerr := os.Open(fn)
	defer r.Close()
	if rerr != nil {
		fmt.Println("open file error : ", fn, rerr)
		return
	}
	f, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println("open file error : ", fn, err)
		return
	}
	fmt.Println(string(f))
	r.Close()
}
func testWriteFile(fn string) {
	fmt.Println("testWriteFile---------------------------------------")
	f, rerr := os.OpenFile(fn, os.O_RDWR, os.ModePerm)

	defer f.Close()
	if rerr != nil {
		fmt.Println("open file error : ", fn, rerr)
		return
	}
	n, werr := io.WriteString(f, "helloGood")
	if werr != nil {
		fmt.Println("write file error : ", fn, werr, n)
		return
	}
	// ioutil.WriteFile(filename, data, perm)
	f.Close()
	testIoutilReadFile(fn)
}
