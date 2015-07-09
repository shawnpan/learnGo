package testbase

import (
	"fmt"
	"os"
)

func TestOS() {
	testFileInfo()
	testGetwd()
	testENV()
	// testMkDir() // ÇëÏÈmkÔÙrm
	// testRmDir()
	testRename()
}
func testGetwd() {
	dir, err := os.Getwd()
	olddir := dir
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("dir = ", dir)

	// ÇÐ»»Ä¿Â¼
	os.Chdir(dir + "/mytest")

	dir, err = os.Getwd()
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("dir = ", dir)

	// ÇÐ»»Ä¿Â¼
	os.Chdir(olddir)

}
func testENV() {
	key := "GOPATH"
	v := os.Getenv(key)
	if v == "" {
		fmt.Println("no env value for key:", key)
	} else {
		fmt.Println("the env value ", key, ":", v)
	}

	key = "path"
	v = os.Getenv(key)
	if v == "" {
		fmt.Println("no env value for key:", key)
	} else {
		fmt.Println("the env value ", key, ":", v)
	}

	err := os.Setenv("temp1", "temp1")
	if err != nil {
		fmt.Println("Error err : ", err)
		return
	}

	data := os.Environ()
	fmt.Println("total env : \n", data)

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("hostname : ", hostname)

}
func testFileInfo() {
	fp := "config/config.txt"
	fi, err := os.Stat(fp)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("FileInfo of file ", fp, " : ", fi)
	fmt.Println("fi.ModTime() : ", fi.ModTime())
	fmt.Println("fi.Name() : ", fi.Name())
	fmt.Println("fi.IsDir() : ", fi.IsDir())
	fmt.Println("fi.Mode() : ", fi.Mode())
	fmt.Println("fi.Size() : ", fi.Size())
	fmt.Println("fi.Sys() : ", fi.Sys())

}
func testMkDir() {
	dir, _ := os.Getwd()
	var path string
	if os.IsPathSeparator('\\') {
		path = "\\"
	} else {
		path = "/"
	}
	fmt.Println("PathSeparator : ", path)
	err := os.Mkdir(dir+path+"temp", os.ModePerm)
	if err != nil {
		fmt.Println("create dir err : ", err)
		return
	}
	err = os.MkdirAll(dir+path+"temp1"+path+"temp2", os.ModePerm)
	if err != nil {
		fmt.Println("create dir err : ", err)
		return
	}

}

func testRmDir() {
	dir, _ := os.Getwd()
	fmt.Println("dir = ", dir)
	err := os.Remove("temp")
	if err != nil {
		fmt.Println("rmove dir err : ", err)
		return
	}
	err = os.RemoveAll("temp1")
	if err != nil {
		fmt.Println("rmove dir err : ", err)
		return
	}
}

func testRename() {
	fp1 := "config/config.txt"
	fp2 := "config/config1.txt"
	fp3 := "config/temp.txt"

	err := os.Rename(fp1, fp2)
	if err != nil {
		if os.IsExist(err) {
			fmt.Println("fp2 is exist", err)
		}
	}
	os.Rename(fp1, fp3)
	fmt.Println("Rename Ok!!")
	os.Rename(fp3, fp1)
	var fi1 os.FileInfo
	var fi2 os.FileInfo
	fi1, err = os.Stat(fp1)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fi2, err = os.Stat(fp2)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("os.SameFile : ", os.SameFile(fi1, fi2))
	fmt.Println("fi1 : ", fi1)
	fmt.Println("fi2 : ", fi2)
}