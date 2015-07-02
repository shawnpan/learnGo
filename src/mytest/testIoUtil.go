package mytest

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

func TestIoUtil() {
	dp := "config"
	f_list, err := ioutil.ReadDir(dp)
	if err != nil {
		fmt.Println("Error : ", err)
	}
	for i, v := range f_list {
		fmt.Println(i, " : ", v)
	}

	fp := "config/config.txt"
	f, rerr := ioutil.ReadFile(fp)
	if rerr != nil {
		fmt.Println("Error : ", err)
	}
	fmt.Println(string(f))

	reader := strings.NewReader("hi man ,what 's this")
	fmt.Println(reflect.TypeOf(reader))
	data, _ := ioutil.ReadAll(reader)
	fmt.Println(string(data))

}
