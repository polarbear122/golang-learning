package main

import "container/list"

const base = 769

type MyHashSet struct {
	data []list.List
}

func HashConstructor() MyHashSet {
	return MyHashSet{make([]list.List, base)}
}
func (s *MyHashSet) Hash(key int) int {
	return key % base
}
func (s *MyHashSet) Add(key int) {
	if !s.Contain(key) {
		h := s.Hash(key)
		s.data[h].PushBack(key)
	}
}
func (s *MyHashSet) Remove(key int) {
	if s.Contain(key) {
		h := s.Hash(key)
		for e := s.data[h].Front(); e != nil; e = e.Next() {
			if e.Value.(int) == key {
				s.data[h].Remove(e)
			}
		}
	}
}
func (s *MyHashSet) Contain(key int) bool {
	h := s.Hash(key)
	for e := s.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			return true
		}
	}
	return false
}
func main() {

}
