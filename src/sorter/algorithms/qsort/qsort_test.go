package qsort

import (
	"testing"
)

func TestQuickSort1(t *testing.T) {
	//初始化创建一个数组切片
	values := []int{5, 4, 3, 2, 1}
	arr := [5]int{1, 2, 3, 4, 5}
	QuickSort(values)
	for i, _ := range values {
		if values[i] != arr[i] {
			t.Error("BubbleSort", values, "结果应该为", arr)
		}
	}
}

func TestQuickSort2(t *testing.T) {
	//初始化创建一个数组切片
	values := []int{5, 5, 3, 2, 1}
	arr := [5]int{1, 2, 3, 5, 5}
	QuickSort(values)
	for i, _ := range values {
		if values[i] != arr[i] {
			t.Error("BubbleSort", values, "结果应该为", arr)
		}
	}
}

func TestQuickSort3(t *testing.T) {
	//初始化创建一个数组切片
	values := []int{5}
	arr := [1]int{5}
	QuickSort(values)
	for i, _ := range values {
		if values[i] != arr[i] {
			t.Error("BubbleSort", values, "结果应该为", arr)
		}
	}
}
