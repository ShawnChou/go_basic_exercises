package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	//var arr [3]int
	//初始默认为0
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 4, 5}
	t.Log(arr1[1], arr2[2])
}

func TestArrayTravel(t *testing.T) {
	arr2 := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr2); i++ {
		t.Log(arr2[i])
	}

	for idx, e := range arr2 {
		t.Log(idx, e)
	}
}
