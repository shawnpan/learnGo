package testIO

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func TestWeb() {
	fmt.Println("-----------testWeb--------------------------------------------")

	// 为uri指定处理器，表示访问/hello使用handleHello方法来处理
	http.HandleFunc("/hello", handleHello)
	// 为uri指定处理器，表示访问/stop使用handleStop方法来处理
	http.HandleFunc("/stop", handleStop)

	fmt.Println("serving on http://localhost:8888/hello")
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}
func handleHello(w http.ResponseWriter, req *http.Request) {
	log.Println("serving", req.URL)
	fmt.Fprintln(w, "Hello, 世界!")
	fmt.Println("access")
}
func handleStop(w http.ResponseWriter, req *http.Request) {
	log.Println("serving", req.URL)
	fmt.Fprintf(w, "我也想停机！！")
	os.Exit(1)

}
