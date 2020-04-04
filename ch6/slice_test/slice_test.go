package slice_test

import (
	"fmt"
	"testing"
)

func TestSliceData(t *testing.T) {
	//数组的相关测试
	arr := [...]int{1, 2, 3, 4} //即使设置...一旦确定数组值的长度，则不能修改
	fmt.Print(arr, arr[1])
	//测试是否共享内存
	arr1 := arr[2:4]
	fmt.Print(arr1)
	arr1[0] = 0
	fmt.Print(arr, len(arr1), cap(arr1))
	//测试是否可以变长 结论：不可变长
	arr = append(arr, 6)
	fmt.Print(arr)
	//slice
	slice := []string{"1", "2", "3"}
	slice = append(slice, "4")
	fmt.Print(slice, len(slice), cap(slice))
}
