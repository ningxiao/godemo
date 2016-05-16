package main

import (
	"bufio"
	"fmt"
	"mplayer/pkg/mlib"
	"mplayer/pkg/mp"
	"os"
	"strconv"
	"strings"
)

var musiclib *mlib.MusicManager
var id int = 0
var ctrl, signal chan int

func playmusic(name string) {
	//传递的是一个结构体的引用
	var music *mlib.MusicEntry = musiclib.Find(name)
	if music == nil {
		fmt.Println("音乐", name, "不存在.")
		return
	}
	mp.Play(music.Source, music.Type)
}
func playcommand(tokens []string) {
	switch tokens[0] {
	case "play":
		if len(tokens) != 2 {
			fmt.Println("输入播放名称错误")
			return
		}
		playmusic(tokens[1])
	case "list":
		for i, l := 0, musiclib.Len(); i < l; i++ {
			music, _ := musiclib.Get(i)
			fmt.Println(i+1, ":", music.Name, music.Artist, music.Source, music.Type)
		}
	case "add":
		if len(tokens) == 5 {
			id++
			musiclib.Add(&mlib.MusicEntry{strconv.Itoa(id), tokens[1], tokens[2], tokens[3], tokens[4]})
		} else {
			fmt.Println("添加音乐 <名称> <演唱者> <地址> <类型> 错误")
		}
	case "remove":
		if len(tokens) == 2 {
			musiclib.RemoveByName(tokens[1])
		} else {
			fmt.Println("删除音乐 <名称> 错误")
		}
	default:
		fmt.Printf("无法识别 %s 命令\n", tokens[0])
	}
}
func main() {
	fmt.Println(`命令控制播放器：
	list -- 查看音乐播放列表
	add <名称> <演唱者> <地址> <类型> -- 添加音乐进入列表
	remove <名称> -- 删除指定名称的音乐
	play <名称> -- 播放指定名称音乐
	`)
	musiclib = mlib.NewMusicManager()
	read := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入命令-> ")
		rawline, _, err := read.ReadLine()
		if err != nil {
			fmt.Println("程序出错")
			break
		}
		line := string(rawline)
		line = strings.Trim(line, "\r\n")
		if line == "q" || line == "e" {
			break
		}
		tokens := strings.Split(line, " ")
		playcommand(tokens)
	}
}
