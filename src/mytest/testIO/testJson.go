package testIO

import (
	"encoding/json"
	"fmt"
)

type PlayerJson struct {
	Id   int    `json:"p_id"` // 还可以自定义输出json的字段名称
	Name string `json:"p_name"`
	Age  int
	Il   []Item
}
type Item struct {
	Id    int `json:"i_id"`
	Price int
}

func TestJSON() {
	itemList := []Item{Item{Id: 1, Price: 22}, Item{Id: 2, Price: 33}, Item{Id: 3, Price: 44}}
	var p PlayerJson = PlayerJson{
		Id:   10,
		Name: "hello",
		Age:  25,
		Il:   itemList}
	//Encoding
	pj, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	jsonString := string(pj)
	fmt.Println(jsonString)

	var ps PlayerJson
	//decoding
	json.Unmarshal([]byte(jsonString), &ps)
	fmt.Println(ps)
}
