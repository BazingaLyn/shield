package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestMap001(t *testing.T) {

	//var k = map[string]string  相当于map k = null;
	var k = make(map[string]string) //相当于new
	k["hello"] = "world"
	k["smile"] = "gogo"
	k["delete"] = "haha"
	fmt.Println(k["hello"])

	for key, value := range k {
		fmt.Println(key + value)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

func TestMap002(t *testing.T) {

	var k map[string]string //nil 会报错
	//var k = make(map[string]string) //相当于new
	k["hello"] = "world"
	k["smile"] = "gogo"
	k["delete"] = "haha"
	fmt.Println(k["hello"])

	for key, value := range k {
		fmt.Println(key + value)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

//~~~~~~~~~~strconv~~~~~~~~
//str to int
func TestStrConv1(t *testing.T) {
	s1 := "100"
	i, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Println(i)
	}

}

//扩展阅读】这是C语言遗留下的典故。C语言中没有string类型而是用字符数组(array)表示字符串，所以Itoa对很多C系的程序员很好理解。
func TestStrConv2(t *testing.T) {
	itoa := strconv.Itoa(20)
	fmt.Println(itoa)
}
