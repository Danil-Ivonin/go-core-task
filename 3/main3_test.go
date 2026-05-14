package main

import "testing"

func TestAddAndGet(t *testing.T) {
	t.Parallel()
	m := NewStringIntMap()
	m.Add("apple", 100)
	value, exists := m.Get("apple")
	if !exists {
		t.Errorf("key should exist")
	}
	if value != 100 {
		t.Errorf("expected 100, got %d", value)
	}
}

func TestRemove(t *testing.T) {
	t.Parallel()
	m := NewStringIntMap()
	m.Add("apple", 100)
	m.Remove("apple")
	_, exists := m.Get("apple")
	if exists {
		t.Errorf("key should be removed")
	}
}

func TestExists(t *testing.T) {
	t.Parallel()
	m := NewStringIntMap()
	m.Add("banana", 200)
	if !m.Exists("banana") {
		t.Errorf("key should exist")
	}
	if m.Exists("orange") {
		t.Errorf("key should not exist")
	}
}

func TestCopy(t *testing.T) {
	t.Parallel()
	m := NewStringIntMap()
	m.Add("a", 1)
	m.Add("b", 2)
	copyMap := m.Copy()
	if len(copyMap) != 2 {
		t.Errorf("expected map length 2, got %d", len(copyMap))
	}
	if copyMap["a"] != 1 {
		t.Errorf("expected value 1")
	}

	copyMap["a"] = 999
	originalValue, _ := m.Get("a")
	if originalValue != 1 {
		t.Errorf("original map should not change")
	}
}

func TestGetNonExistingKey(t *testing.T) {
	t.Parallel()
	m := NewStringIntMap()
	_, exists := m.Get("unknown")
	if exists {
		t.Errorf("key should not exist")
	}
}
