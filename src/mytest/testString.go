package mytest

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
	a := "jiangkou"
	b := "ang"
	fmt.Println("a=", a, "b=", b)
	//查找某个字符是否在这个字符串中存在，存在返回true
	fmt.Println(strings.Contains(a, b))
	fmt.Println(strings.Contains(b, a))

	//查询字符串中是否包含多个字符
	fmt.Println(strings.ContainsAny(a, "h&z"))
	fmt.Println(strings.ContainsAny(a, "i&o"))

	//一段字符串中有多少匹配到的字符
	fmt.Println(strings.Count(a, b))
	fmt.Println(strings.Count("hello hello", "l"))

	//两个字符串在完全小写的情况下是否相等uff8编码
	fmt.Println(strings.EqualFold(a, "JiAnGKOu"))

	//空格来分割字符串最后返回的是[]string的切片
	fmt.Println(strings.Fields("aa bcd efghij"))

	//根据自定义函数来切分
	fmt.Println(strings.FieldsFunc("aa|bcd_d|fgh_i|ej", func(s rune) bool {
		if s == rune('|') {
			return true
		}
		return false
	}))
	fmt.Println(strings.Split("aa|bcd_d|fgh_i|ej", "|"))

	//判断是否以某个字符串开头
	fmt.Println(strings.HasPrefix(a, "jia"))

	//判断是否以某个字符串结尾
	fmt.Println(strings.HasSuffix(a, "ou"))
}
