package mytest

import (
	"fmt"
)

type Player struct {
	id   int
	name string
}

func (p Player) getId() int {
	return p.id
}

func TestStruct() {
	p := Player{}
	p.id = 10
	p.name = "hello"
	fmt.Println(p)
	fmt.Println(p.getId())

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
