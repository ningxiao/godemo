package main

import (
	"calcproj/simplemath"
	"flag"
	"fmt"
	"strconv"
)

var Usage = func() {
	fmt.Println("USAGE: calc command [arguments] ...")
	fmt.Println("\nThe commands are:\n\tadd\tAddition of two values.\n\tsqrt\tSquare root of a non-negative value.")
}

/*
 * 匿名变量与返回多个变量
 * @param  {string} name 传入名称
 * @return {string,string,string} firstName,lastName,nickName 返回多个值
 */
func GetName(name string) (firstName string, lastName bool, nickName string) {
	b := (1 != 0) //动态推导出数据类型
	return "May", b, "chibi" + name
}
func main() {
	flag.Parse()
	args := flag.Args()
	_, b, nickName := GetName("宁肖")
	fmt.Print(b, nickName, "\n")
	if args == nil || len(args) < 2 {
		Usage()
		return
	}
	switch args[0] {
	case "add":
		if len(args) != 3 {
			fmt.Println("USAGE: calc add <integer1> <integer2>")
			return
		}
		v1, err1 := strconv.Atoi(args[1])
		v2, err2 := strconv.Atoi(args[2])
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: calc add <integer1> <integer2>")
			return
		}
		ret := simplemath.Add(v1, v2)
		fmt.Println("Result: ", ret)
	case "sqrt":
		if len(args) != 2 {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}
		v, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}
		ret := simplemath.Sqrt(v)
		fmt.Println("Result: ", ret)
	default:
		Usage()
	}
}
