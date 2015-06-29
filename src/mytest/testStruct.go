package mytest

import (
	"fmt"
)

type Player struct {
	id   int
	name string
}

func (p *Player) GetId() int {
	return p.id
}

//该方法可以修改成功
func (p *Player) SetId(id int) {
	p.id = id
}

//该方无法修改对象的id值
func (p Player) SetIId(id int) {
	p.id = id
}

func TestStruct() {
	p := Player{}
	p.id = 10
	p.name = "hello"
	fmt.Println(p)

	p.SetId(11)
	fmt.Println(p)

	p.SetIId(12)
	fmt.Println(p)

	p1 := Player{id: 1, name: "boys"}
	fmt.Println(p1)

	modifyPlayer1(p1)
	fmt.Println(p1)

	modifyPlayer2(&p1)
	fmt.Println(p1)

}
func modifyPlayer1(p Player) {
	p.id = 99
	p.name = "gudo"
}

func modifyPlayer2(p *Player) {
	p.id = 99
	p.name = "gudo"
}
