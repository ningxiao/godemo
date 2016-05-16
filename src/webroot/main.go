package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	//"webroot/daemon"

	_ "github.com/go-sql-driver/mysql"
)

var tmpl string
var sqldb *sql.DB
var port *string = flag.String("port", "8080", "设置端口")

type userinfo struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Sex   int    `json:"sex"`
	Age   int    `json:"age"`
	Level int    `json:"level"`
}

func initdb() {
	sqldb, _ = sql.Open("mysql", "root:19870615@tcp(localhost:3306)/test?charset=utf8")
	sqldb.SetMaxOpenConns(200)
	sqldb.SetMaxIdleConns(100)
	sqldb.Ping()
}
func createip() string {
	var ip string = "127.0.0.1"
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ip
	}
	for _, address := range addrs { // 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
			}
		}
	}
	return ip
}
func hello(rw http.ResponseWriter, req *http.Request) {
	rows, err := sqldb.Query("SELECT * FROM user")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}
	users := []userinfo{}
	for rows.Next() {
		user := userinfo{}
		rows.Columns()
		err = rows.Scan(&user.Id, &user.Name, &user.Sex, &user.Age, &user.Level)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}
	res2B, _ := json.Marshal(users)
	io.WriteString(rw, string(res2B))
}
func view(name string) *template.Template {
	//读取文件
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	str, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	tpl, err := template.New("play.html").Parse(string(str))
	return tpl
}
func rendertmpl(rw http.ResponseWriter, req *http.Request) {
	rows, err := sqldb.Query("SELECT * FROM user")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}
	onlineuser := []userinfo{}
	for rows.Next() {
		user := userinfo{}
		rows.Columns()
		err = rows.Scan(&user.Id, &user.Name, &user.Sex, &user.Age, &user.Level)
		if err != nil {
			panic(err.Error())
		}
		onlineuser = append(onlineuser, user)
	}
	//tpl := view("play.html")
	tpl, err := template.ParseFiles("play.html")
	err = tpl.Execute(rw, onlineuser)
	if err != nil {
		panic(err.Error())
	}
}
func test(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "callback({status: 200,data:})")
}
func playtmpl(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, string(tmpl))
}
func main() {
	//daemon.Daemon(0, 0)
	flag.Parse()
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	path := strings.Replace(dir, "\\", "/", -1)
	//读取文件
	file, err := os.OpenFile(path+"/play.html", os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	str, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	tmpl = string(str)
	initdb()
	fmt.Print("http://" + createip() + ":" + *port)
	http.HandleFunc("/", hello)          //设定访问的路径
	http.HandleFunc("/test", test)       //设定访问的路径
	http.HandleFunc("/tmpl", rendertmpl) //设定访问的路径
	http.HandleFunc("/play", playtmpl)   //设定访问的路径
	http.ListenAndServe(":"+*port, nil)  //设定端口和handler
}
