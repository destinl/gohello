package algorithm

import "math/rand/v2"

const MaxLevel = 16

type SkipListNode struct {
	key   int
	Value interface{}
	next  []*SkipListNode
}

type SkipList struct {
	head  *SkipListNode
	level int
}

func NewSkipList() *SkipList {
	return &SkipList{
		head: &SkipListNode{
			next: make([]*SkipListNode, MaxLevel),
		},
		level: 0,
	}
}

func randomLevel() int {
	level := 0
	for rand.Float32() < 0.5 && level < MaxLevel-1 {
		level++
	}
	return level
}

func (s *SkipList) Search(key int) (interface{}, bool) {
	x := s.head
	for i := s.level - 1; i >= 0; i-- {
		for x.next[i] != nil && x.next[i].key < key {
			x = x.next[i]
		}
	}
	x = x.next[0]
	if x == nil || x.key != key {
		return nil, false
	}
	return x.Value, true
}

func (s *SkipList) Insert(key int, value interface{}) {
	update := make([]*SkipListNode, MaxLevel)
	x := s.head
	for i := s.level - 1; i >= 0; i-- {
		for x.next[i] != nil && x.next[i].key < key {
			x = x.next[i]
		}
		update[i] = x
	}
	x = x.next[0]
	if x != nil && x.key == key {
		x.Value = value
		return
	}
	level := randomLevel()
	if level > s.level {
		for i := s.level; i < level; i++ {
			update[i] = s.head
		}
		s.level = level
	}
	newNode := &SkipListNode{
		key:   key,
		Value: value,
		next:  make([]*SkipListNode, level+1),
	}
	for i := 0; i <= level; i++ {
		newNode.next[i] = update[i].next[i]
		update[i].next[i] = newNode
	}
}

func (s *SkipList) Delete(key int) {
	update := make([]*SkipListNode, MaxLevel)
	x := s.head
	for i := s.level - 1; i >= 0; i-- {
		for x.next[i] != nil && x.next[i].key < key {
			x = x.next[i]
		}
		update[i] = x
	}
	x = x.next[0]
	if x == nil || x.key != key {
		return
	}
	for i := 0; i <= s.level; i++ {
		if update[i].next[i] != x {
			break
		}
		update[i].next[i] = x.next[i]
	}
	for s.level > 0 && s.head.next[s.level] == nil {
		s.level--
	}
}
