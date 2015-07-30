package testbase

import (
	"fmt"
	"reflect"
)

// 接口的定义
type Human interface {
	GetName() string
}
type Teacher struct {
	name string
}

// Teacher实现了接口Human的方法
func (p Teacher) GetName() string {
	return p.name
}

type Gamer interface {
	PlayGame(game string)
	Human // 接口组合，将从该接口继承相应的方法过来
}
type Sky struct {
	Name string
}

func (s Sky) GetName() string {
	return s.Name
}
func (s Sky) PlayGame(game string) {
	fmt.Println(s.Name, " play ", game)
}

// Interface 是一组抽象方法(未具体实现的方法/仅包含方法名参数返回值的方法)的集合，有点像但又不同于其他编程语言中的 interface 。
// 如果实现了 interface 中的所有方法，即该类/对象就实现了该接口
func TestInterface() {
	fmt.Println("testInterface-------------------------------------------------")
	t := Teacher{"teacher"}
	var h Human = t
	fmt.Println(h, reflect.TypeOf(h))

	needHuman(t)

	var g Gamer = Sky{"sky"}
	g.PlayGame("super mario")
}

func needHuman(h Human) {
	fmt.Println(h.GetName())
}
