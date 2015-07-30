package testbase

import (
	"fmt"
	"reflect"
)

type People struct {
	Name string
}

func (p People) GetName() string {
	return p.Name
}

func TestReflect() {
	fmt.Println("testReflect---------------------------------------------------—")
	i := 10
	to := reflect.TypeOf(i)
	fmt.Println(to, reflect.TypeOf(&i), reflect.ValueOf(i))
	fmt.Println(to.NumMethod())

	p := People{"ppp"}
	fmt.Println(reflect.TypeOf(p).NumMethod(), ",", reflect.TypeOf(p).NumField())

	m := reflect.TypeOf(p).Method(0)
	fmt.Println(m, ",", m.Func, ",", m.Name)

	// 通过反射调用方法
	m1 := reflect.ValueOf(p).Method(0)
	fmt.Println(m1.Call([]reflect.Value{}))

}
