package mytest

import (
	"bufio"
	"fmt"
	"net"
)

func TestNet() {
	fmt.Println("testNet")
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("testNet")
	read := bufio.NewReader(conn)
	rs, rerr := read.ReadString('\n')
	for {
		if rerr != nil {
			fmt.Println(rs)
			rs, rerr = read.ReadString('\n')
		} else {
			fmt.Println("Read Error : ", rerr)
			break
		}
	}

}
func TestMyNet() {
	fmt.Println("testNet")
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("testNet")
	fmt.Println(conn)
	// read := bufio.NewReader(conn)
	// rs, rerr := read.ReadString('\n')
	// for {
	// 	if rerr != nil {
	// 		fmt.Println(rs)
	// 		rs, rerr = read.ReadString('\n')
	// 	} else {
	// 		fmt.Println("Read Error : ", rerr)
	// 		break
	// 	}
	// }

}
