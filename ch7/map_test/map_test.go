package map_test

import "testing"

func TestMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2], len(m1))
	m2 := map[int]int{}
	t.Log(len(m2))
	m3 := make(map[int]int, 10)
	t.Log(len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 3
	if v, ok := m1[3]; ok {
		t.Log(v)
	} else {
		t.Log(v)
	}
}

func TestTraveMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
