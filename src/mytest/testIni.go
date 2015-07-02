package mytest

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var per map[string]interface{}

func TestIni() {
	per = make(map[string]interface{})
	fp := "config/test.ini"
	f, _ := os.Open(fp)
	buf := bufio.NewReader(f)
	for {
		l, err := buf.ReadString('\n')
		line := strings.TrimSpace(l)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			if len(line) == 0 {
				break
			}
		}
		switch {
		case len(line) == 0:
		case line[0] == '[' && line[len(line)] == ']':
			section := strings.TrimSpace(line[1 : len(line)-1])
			fmt.Println(section)
		default:
			//dnusername = xiaowei 这种的可以匹配存储
			i := strings.IndexAny(line, "=")
			per[strings.TrimSpace(line[0:i])] = strings.TrimSpace(line[i+1:])
		}
	}
	//循环输出结果
	for k, v := range per {
		fmt.Println(k, " = ", v)
	}
}
