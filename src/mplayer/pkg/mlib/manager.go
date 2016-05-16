package mlib

import "errors"

/**
* 音乐列表控制器结构体
* 数据源为一个数组musics 里面元素类型为MusicEntry
**/
type MusicManager struct {
	musics []MusicEntry
}

/**
* 给MusicManager结构体添加获取长度方法
**/
func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("获取文件索引超出音乐列表范围")
	}
	return &m.musics[index], nil
}
func (m *MusicManager) Find(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}
	for _, m := range m.musics {
		if m.Name == name {
			return &m
		}
	}
	return nil
}
func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}
func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= len(m.musics) {
		return nil
	}
	music := &m.musics[index]
	m.musics = append(m.musics[:index], m.musics[index+1:]...)
	return music
}
func (m *MusicManager) RemoveByName(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}

	for i, v := range m.musics {
		if v.Name == name {
			return m.Remove(i)
		}
	}
	return nil
}

/**
 * 音乐播放列表构造函数创建一个列表的引用
 * 数据为一个MusicEntry的数组切片
 * @return {*MusicManager} 返回一个MusicManager结构体引用
 */
func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}
