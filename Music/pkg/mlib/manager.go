package mlib

import "errors"
// 定义音乐结构
type MusicEntry struct {
	Id     string
	Name   string
	Artist string
	Source string
	Type   string
}

// 定义音乐库类
type MusicManger struct {
	musics []MusicEntry
}

// 实例化音乐库类
func NewMusicManager() *MusicManger {
	return &MusicManger{musics: make([]MusicEntry, 0)}
}

// 为MusicManage 添加Len方法
func (m *MusicManger) Len() int {
	return len(m.musics)
}
// 为MusicManage 添加Get方法
func (m *MusicManger) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("Index out of range")
	}
	return &m.musics[index], nil
}
// 为MusicManage 添加Find方法
func (m *MusicManger) Find(name string) *MusicEntry {
	if 0 == len(m.musics) {
		return nil
	}

	for _, m := range m.musics {
		if name == m.Name {
			return &m
		}
	}

	return nil
}

func (m *MusicManger) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}
// 为MusicManage 添加Remove方法
func (m *MusicManger) Remove(index int) *MusicEntry {
	if index < 0 || index > len(m.musics) {
		return nil
	}
	removeMusic := &m.musics[index]

	if index < len(m.musics)-1 {
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
	} else if index == 0 {
		m.musics = make([]MusicEntry, 0)
	} else {
		m.musics = m.musics[:index-1]
	}

	return removeMusic
}



