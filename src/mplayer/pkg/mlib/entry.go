package mlib

/**
* 自定义一个关于音乐结构体
* 结构体第一个字母大写外部可见
* 包含了音乐的相关信息
**/
type MusicEntry struct {
	Id     string //唯一ID
	Name   string //音乐名称
	Artist string //演唱者
	Source string //文件位置
	Type   string //文件类型(mp3或者wav)
}
