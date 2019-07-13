package mlib

import "errors"

type MusicEntry struct {
	Id     string
	Name   string
	Artist string
	Source string
	Type   string
}

type MusicManger struct {
	musics []MusicEntry
}

func NewMusicManager() *MusicManger {
	return &MusicManger{musics: make([]MusicEntry, 0)}
}

func (m *MusicManger) Len() int {
	return len(m.musics)
}

func (m *MusicManger) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("Index out of range")
	}
	return &m.musics[index], nil
}

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



