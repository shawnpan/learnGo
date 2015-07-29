package testbase

import (
	"fmt"
)

// 声明
type Player struct {
	id   int
	name string
}
type Boss struct {
	// 匿名字段，那么默认Boss就包含了Player的所有字段
	Player
	name string
}

// 声明该函数为Player对象的函数，且默认传入方法的对象是，调用该方法的对像的指针
func (p *Player) GetId() int {
	return p.id
}

//该方法可以修改成功，因为默认传入的是调用者的指针
func (p *Player) SetId(id int) {
	p.id = id
}

func (p *Player) SetName(name string) {
	p.name = name
}

//该方无法修改对象的id值，因为默认传入的是调用者的的一份拷贝
func (p Player) SetIId(id int) {
	p.id = id
}

//对象传递，会复制一个对象副本，对副本操作不影响原对象
func modifyPlayer1(p Player) {
	p.id = 99
	p.name = "gudo"
}

//指针传递
func modifyPlayer2(p *Player) {
	p.id = 99
	p.name = "gudo"
}

func TestStruct() {
	// 声明和赋值
	p := Player{}
	// p := new(Player) 返回的是*Player
	// var p Player
	// p:= Player{id: 10, name: "hello"}
	// p:= Player{10, "hello"}
	// p:= &Player{10, "hello"} 返回的是*Player

	p.id = 10
	p.name = "hello"
	fmt.Println(p)

	p.SetId(11)
	fmt.Println(p)

	p.SetIId(12)
	fmt.Println(p)

	p2 := Player{id: 1, name: "boys"}
	fmt.Println(p2)

	// 参数以值传递
	modifyPlayer1(p2)
	fmt.Println(p2)

	// 参数以指针传递
	modifyPlayer2(&p2)
	fmt.Println(p2)

	// 匿名字段就是这样，能够实现字段的继承
	boss1 := Boss{Player{1, "player"}, "boss"}
	printBoss(&boss1)

	// 还可以直接访问方法，该方法属于Player，只会修改boss.Player.name
	boss1.SetName("SB")
	printBoss(&boss1)

}

func printBoss(boss *Boss) {
	// 还可以这样访问
	// 当Boss与Player都拥有name字段时，使用boss.name访问的是Boss的name字段，访问Player的name字段需要使用boss.Player.name
	fmt.Println("id=", boss.id, ",boss.Player.name=", boss.Player.name, ",boss.name=", boss.name)
}
