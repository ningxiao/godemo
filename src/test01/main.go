package main

import (
	"fmt"
	"log"
	"os"
)

//自定义一个矩阵复合体 并且添加对应方法
type Rect struct {
	x, y          float64
	width, height float64
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}

//声明一个结构体
type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

//声明一个新类型
type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

//指针类型方法
func (a *Integer) Add(b Integer) {
	*a += b
}

//值传递类型方法
func (a Integer) Adds(b Integer) {
	a += b
}

//在一个结构体里面使用日志方法
type Job struct {
	Command string
	*log.Logger
}

func (job *Job) Log(log string) {
	job.Logger.Println(log)
}
func (job *Job) Start() {
	job.Log("start now -----")
	job.Log("started")
}

/*
 * 数组值传递修改不会影响原始数据
 * @param  {array} array 传入数组
 */
func modify(array [5]int) {
	array[0] = 10
	fmt.Println("In modify(), array values:", array)
}

/*
 *切片引用传递会修改原始数据
 * @param  {array} list 传入数组
 */
func quoteArray(list []int) {
	list[0] = 15
	fmt.Println("In quoteArray(), array values:", list)
}

/*
 * 传递多个指定类型参数
 * @param  {args} int 传入数组
 */
func myfunc(args ...int) {
	for _, arg := range args {
		fmt.Print(arg)
	}
}

/*
 * 传递多个不定类型参数
 * @param  {args} interface 接口
 */
func myprintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) { //获取数据类型
		case int:
			fmt.Print("输入数据为int类型", "\n")
		case string:
			fmt.Print("输入数据为string类型", "\n")
		case int64:
			fmt.Print("输入数据为int64类型", "\n")
		default:
			fmt.Print("输入数据未知F类型", "\n")
		}
	}
}
func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}
func NewJob(command string, logger *log.Logger) *Job {
	return &Job{command, logger}
}
func main() {
	str := "Hello 世界"
	ch := str[0]       //类型为byte
	chs := []rune(str) //转为字符串切片
	fmt.Printf("输出内容 \"%s\" 字节长度 %d 字符个数 %d \n", str, len(str), len(chs))
	fmt.Printf("输出内容 \"%s\" 第一个字符 %c %c \n", str, ch, chs[7])
	for i, ch := range str {
		fmt.Printf("%d\" %c\n", i, ch)
	}
	array := [5]int{1, 2, 3, 4, 5} //定义数组
	modify(array)
	list := []int{1, 2, 3, 4, 5} //定义数组切片
	quoteArray(list)
	fmt.Println("In main(), array values:", array)
	fmt.Println("In main(), array values:", list)

	//数组切片
	myArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//mySlice := myArray[:5] //获取前五个数据 切片是数据引用
	mySlice := myArray[5:] //获取后五个数据 切片是数据引用
	mySlice[2] = 45
	fmt.Println("Elements of myArray")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}
	fmt.Println("\nElements of mySlice")
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Println()
	slice := make([]int, 5, 10) //初始化个数为5 并且预留10个元素空间
	fmt.Printf("slice切片 元素数量 %d 空间大小 %d \n", len(slice), cap(slice))
	slice = append(slice, 1, 2, 3) //单个元素追加
	slice1 := []int{4, 5, 6}
	//上面切片预留了十个空间如果超出的时候回动态分配一块
	slice = append(slice, slice1...) //将一个切片打散追加进去
	fmt.Println("slice values:", slice)
	copy(slice1, mySlice) //将切片mySlice复制到slice1
	fmt.Println("slice values:", slice1)
	//var personDB map[string]PersonInfo
	personDB := make(map[string]PersonInfo)
	personDB["511324"] = PersonInfo{"511324", "宁肖", "北京市朝阳区"}
	personDB["511334"] = PersonInfo{"511334", "张三", "北京市海淀区"}
	person, ok := personDB["511324"] //查找到指定的key
	if ok {                          //判断是否ok就行
		fmt.Print("ID     姓名       地址\n")
		fmt.Print(person.ID, " ", person.Name, "       ", person.Address, "\n")
	} else {
		fmt.Print("对不起查无此人\n")
	}
	delete(personDB, "511334") //删除指定键值
	a := []int8{1, 2, 3, 4, 5, 6}
	for i, l := 0, len(a); i < l; i++ { //平行赋值
		fmt.Print(a[i], "\n")
	}
	//传入多个指定类型参数
	myfunc(1, 2, 3, 4)
	fmt.Print("\n")
	myfunc(1, 2, 3, 4, 5, 6, 7)
	fmt.Print("\n")
	//传入多个未知类型参数
	var (
		v1 int     = 1
		v2 int64   = 234
		v3 string  = "hello"
		v4 float32 = 1.2345
	)
	myprintf(v1, v2, v3, v4)
	//闭包使用执行匿名函数
	var j int = 5
	fun := func() func() {
		var i int = 10
		return func() {
			fmt.Printf("i, j:%d, %d\n", i, j)
		}
	}()
	fun()
	j *= 2
	fun()
	//给自定义类型添加方法
	fmt.Printf("给自定义类型添加方法\n")
	var aut Integer = 1
	if aut.Less(2) {
		fmt.Println(aut, "Less 2")
	}
	aut.Add(2)
	fmt.Println(aut, "Add 指针\n")
	aut.Adds(2)
	fmt.Println(aut, "Add 值计算\n")
	//大多数类型属于值传递
	testa := [3]int{1, 2, 3}
	testb := testa
	testb[1]++
	testc := &testa //指针传递会影响原值
	fmt.Println(testa, testb, "\n")
	testc[1]++
	fmt.Println(testa, testc, "\n")
	//自定义复合类型
	rect1 := new(Rect)
	rect2 := &Rect{}
	rect3 := &Rect{0, 0, 100, 200}
	rect4 := &Rect{width: 100, height: 200}
	rect5 := NewRect(10, 10, 110, 210)
	fmt.Println(rect5, rect4, rect1, rect2, rect3, "\n")
	//自定义复合类型调用日志
	job := NewJob("测试日志输出", log.New(os.Stderr, "Job: ", log.Ldate))
	job.Start()
}
