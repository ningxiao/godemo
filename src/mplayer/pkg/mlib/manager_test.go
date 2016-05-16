package mlib

import "testing"

func TestOps(t *testing.T) {
	musics := NewMusicManager()
	if musics == nil {
		t.Error("创建音乐管理器失败。")
	}
	if musics.Len() != 0 {
		t.Error("创建的音乐管理器，不是空的。")
	}
	music := &MusicEntry{"010001", "独家记忆", "陈小春", "http://192.168.203.71/24501234", "mp3"}
	musics.Add(music)
	if musics.Len() != 1 {
		t.Error("音乐管理器添加失败。")
	}
	m := musics.Find(music.Name)
	if m == nil {
		t.Error("音乐管理器名称查找失败。")
	}
	m, err := musics.Get(0)
	if m == nil {
		t.Error("音乐管理器索引查找失败。", err)
	}
	m = musics.Remove(0)
	if m == nil || musics.Len() != 0 {
		t.Error("音乐管理器删除文件失败。")
	}
}
