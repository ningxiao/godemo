package main

/**
*启动命令：go run main.go -i unsorted.dat -o sorted.dat -a qsort
*输出结果：
* infile:  unsorted.dat  // 因为指定了路径，所以覆盖了默认路径
* outfile:  sorted.dat         // 启动命令里带 -d 参数，所以启用默认值
* algorithm qsort        //排序类型
**/
import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"sorter/algorithms/bubblesort"
	"sorter/algorithms/qsort"
	"strconv"
	"time"
)

var infile *string = flag.String("i", "unsorted.dat", "文件包含排序值")
var outfile *string = flag.String("o", "sorted.dat", "文件接收排序的值")
var algorithm *string = flag.String("a", "qsort", "排序算法")

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
func readvalues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("无法打开输入文件", infile)
		return
	}
	//返回之前关闭文件流
	defer file.Close()
	//读取文件
	br := bufio.NewReader(file)
	//创建一个int类型的数组切片长度为0 切片回更加数据内容动态扩容
	values = make([]int, 0)
	for {
		line, isprefix, ioerr := br.ReadLine()
		if ioerr != nil {
			if ioerr != io.EOF {
				err = ioerr
			}
			break
		}
		if isprefix {
			fmt.Println("数据太大")
			return
		}
		str := string(line)               //将字符转为字符串
		value, ioerr := strconv.Atoi(str) //将字符串转为int类型
		if ioerr != nil {
			err = ioerr
			return
		}
		values = append(values, value) //单个元素追加
	}
	return
}
func writevalues(values []int, outfile string) error {
	var (
		file *os.File
		err  error
	)
	if checkFileIsExist(outfile) {
		//开启写入模式并且追加 如果不开启写完 linux 报异常无法写入文件
		file, err = os.OpenFile(outfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	} else {
		file, err = os.Create(outfile)
	}
	if err != nil {
		fmt.Println("写入排序结果错误", outfile)
		return err
	}
	defer file.Close()
	buffer := bytes.NewBufferString("") //创建字符串缓冲
	for _, value := range values {
		buffer.WriteString(strconv.Itoa(value) + "\n") //写入字符串缓冲
	}
	_, err = file.WriteString(buffer.String())
	return err
}
func main() {
	flag.Parse()
	if infile != nil {
		fmt.Println("infile =", *infile, "outfile = ", *outfile, "algorithm = ", *algorithm)
	}
	values, err := readvalues(*infile)
	if err == nil {
		t1 := time.Now()
		switch *algorithm {
		case "qsort":
			qsort.QuickSort(values)
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default:
			fmt.Println("排序算法", *algorithm, "不存在")
		}
		err = writevalues(values, *outfile)
		t2 := time.Now()
		if err != nil {
			fmt.Println("写入排序文件异常", err, "\n")
		}
		fmt.Println("排序耗时", t2.Sub(t1), "毫秒\n排序结果", values)
	} else {
		fmt.Println("读取文件错误", err)
	}
}
