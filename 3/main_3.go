package main

import "fmt"

type StringIntMap struct {
	m map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{make(map[string]int)}
}

func (sm *StringIntMap) Add(key string, value int) {
	sm.m[key] = value
}

func (sm *StringIntMap) Remove(key string) {
	delete(sm.m, key)
}

func (sm *StringIntMap) Copy() map[string]int {
	res := make(map[string]int)
	for k, v := range sm.m {
		res[k] = v
	}
	return res
}

func (sm *StringIntMap) Exists(key string) bool {
	_, exists := sm.m[key]
	return exists
}

func (sm *StringIntMap) Get(key string) (int, bool) {
	v, exist := sm.m[key]
	return v, exist
}

func main() {
	m := NewStringIntMap()

	m.Add("apple", 10)
	m.Add("banana", 20)
	fmt.Println(m.m)

	value, ok := m.Get("apple")
	fmt.Printf("get apple: value: %v, ok: %v\n", value, ok)

	exists := m.Exists("banana")
	fmt.Printf("exists banana: %v\n", exists)

	m.Remove("banana")
	fmt.Printf("remove banana: %v\n", m.m)

	copyMap := m.Copy()
	fmt.Printf("copy map: %v\n", copyMap)
}
