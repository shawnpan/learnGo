package testIO

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
)

func TestMyNet() {
	conn, err := net.Dial("tcp", "baidu.com:80")
	if err != nil {
		// handle error
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	rder := bufio.NewReader(conn)
	status, rerr := rder.ReadString('\n')
	for {
		if rerr != nil {
			fmt.Println("rerr != nil")
		}
		if io.EOF == rerr {
			fmt.Println("io.EOF == rerr")
		}
		if rerr != nil || io.EOF == rerr {
			break
		}
		status, rerr = rder.ReadString('\n')
		fmt.Println(status)
	}

}

func TestGet() {
	resp, err := http.Get("http://jandan.net/pic/page-6871#comments")
	processErr(err)
	printResponse(resp)
}
func printResponse(resp *http.Response) {
	fmt.Println("resp.Status : ", resp.Status)
	fmt.Println("resp.StatusCode : ", resp.StatusCode)
	fmt.Println("resp.ContentLength : ", resp.ContentLength)
	fmt.Println("resp.TransferEncoding : ", resp.TransferEncoding)
	header := resp.Header
	fmt.Println("resp.Header : ", header)
	fmt.Println("header.Get(\"Set-Cookie\") : ", header.Get("Set-Cookie"))

	fmt.Println("resp.Proto : ", resp.Proto)
	fmt.Println("resp.Request : ", resp.Request)

	fmt.Println("resp.Cookies() : ", resp.Cookies())
	cookies := resp.Cookies()
	fmt.Println("cookies len : ", len(cookies))
	for i, v := range cookies {
		fmt.Println("cookies ", i, " : ", v)
	}

	body := resp.Body
	fmt.Println("resp.Body : ", body)
	defer body.Close()
	rder := bufio.NewReader(body)
	msg, rerr := rder.ReadString('\n')
	for {
		if rerr != nil {
			fmt.Println("rerr != nil")
		}
		if io.EOF == rerr {
			fmt.Println("io.EOF == rerr")
		}
		if rerr != nil || io.EOF == rerr {
			break
		}
		msg, rerr = rder.ReadString('\n')
		fmt.Println(msg)
	}

}

func processErr(err interface{}) {
	if err != nil {
		fmt.Println("Error : ", err)
	}
}

func TestNewRequest() {
	client := &http.Client{}
	reqest, _ := http.NewRequest("GET", "http://www.baidu.com", nil)

	reqest.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	reqest.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	reqest.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	reqest.Header.Set("Cache-Control", "max-age=0")
	reqest.Header.Set("Connection", "keep-alive")

	response, _ := client.Do(reqest)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		bodystr := string(body)
		fmt.Println(bodystr)
	}
}
