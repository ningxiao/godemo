package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

func httpGet(id string) {
	resp, err := http.Get("http://www.01happy.com/demo/accept.php?id=" + id)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func httpPost() {
	resp, err := http.Post("http://www.01happy.com/demo/accept.php", "application/x-www-form-urlencoded", strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func httpPostForm() {
	resp, err := http.PostForm("http://www.01happy.com/demo/accept.php",
		url.Values{"key": {"Value"}, "id": {"123"}})

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}

func httpDo() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://www.01happy.com/demo/accept.php", strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
func add(wg sync.WaitGroup) {
	wg.Add(1)
}
func done(wg *sync.WaitGroup) {
	httpGet("85")
	wg.Done()
}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 打开数据库，sns是我的数据库名字，需要替换你自己的名字，（官网给的没有加tcp，跑不起来，具体有时 间看看源码分析下为何）
	db, err := sql.Open("mysql", "root:19870615@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// topic是我本地数据库的表名，需要替换你自己的表名，这里面的英文注释都是引用github官网的~~
	//  嘿嘿 我只是想跑起来看看
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		panic(err.Error())
	}

	// 获取列名
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	// 切片读取内容
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	//读取行
	for rows.Next() {
		// 从数据中获得原始字节
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
	dstFile, err := os.Create("mysql.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//写入文件
	defer dstFile.Close()
	s := "hello world"
	dstFile.WriteString(s + "\n")

	//读取文件
	fi, err := os.Open("mysql.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	fmt.Println(string(fd))
	//http开发
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
	}
	for i := 0; i < 10; i++ {
		go done(&wg)
	}
	wg.Wait()
	fmt.Println("exit")
	//httpGet()
	// httpPost()
	// httpPostForm()
	// httpDo()
}
