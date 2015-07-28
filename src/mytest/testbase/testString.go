package testbase

import (
	"fmt"
	"strings"
)

func TestString() {
	fmt.Println("testString --------------------------------")
	var s string = "abcdefg"
	fmt.Println("s = ", s)

	l := len(s)
	fmt.Println(l)

	ss := s[2:5]
	fmt.Println(ss)

	//转换字符串的内容，先转换a的类型为[]byte
	c := []byte(s)
	fmt.Println("c = ", c)

	c[0] = 'z'
	var d string = string(c)
	fmt.Println("d = ", d)

}

func TestStrings() {
	fmt.Println("testStrings---------------------------------")
	a := "jiangkou jiangkou"
	b := "ang"
	fmt.Println("[a=", a, "][b=", b, "]")
	//查找某个字符是否在这个字符串中存在，存在返回true
	fmt.Println("strings.Contains(a, b) = ", strings.Contains(a, b))
	fmt.Println("strings.Contains(b, a) = ", strings.Contains(b, a))

	//查询字符串中是否包含多个字符
	c := "h&z"
	d := "i&o"
	fmt.Println("[c=", c, "][d=", d, "]")
	fmt.Println("strings.ContainsAny(a, c) = ", strings.ContainsAny(a, c))
	fmt.Println("strings.ContainsAny(a, d) = ", strings.ContainsAny(a, d))

	//一段字符串中有多少匹配到的字符
	fmt.Println("strings.Count(a, b) = ", strings.Count(a, b))

	//两个字符串在完全小写的情况下是否相等
	e := "JiAnGKOu"
	fmt.Println("e=", e)
	fmt.Println("strings.EqualFold(a, e) = ", strings.EqualFold(a, e))

	//空格来分割字符串最后返回的是[]string的切片
	f := "aa bcd efghij"
	fmt.Println("f=", f)
	fmt.Println("strings.Fields(f) = ", strings.Fields(f))

	//根据自定义函数来切分
	i := "aa|bcd_d|fgh_i|ej"
	j := '|'
	fmt.Println("[i=", i, "][j=", j, "]")
	fmt.Println(strings.FieldsFunc(i, func(s rune) bool {
		if s == rune(j) {
			return true
		}
		return false
	}))

	fmt.Println("strings.Split(i, string(j)) = ", strings.Split(i, string(j)))

	//判断是否以某个字符串开头
	g := "jia"
	fmt.Println("g=", g)
	fmt.Println("strings.HasPrefix(a, g) = ", strings.HasPrefix(a, g))

	//判断是否以某个字符串结尾
	h := "ou"
	fmt.Println("h=", h)
	fmt.Println("strings.HasSuffix(a, h) = ", strings.HasSuffix(a, h))
}
