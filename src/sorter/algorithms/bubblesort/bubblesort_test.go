package bubblesort

/**
*冒泡排序的单元测试
**/
import (
	"testing"
)

/*
 * 不存在相同数据单元测试
 * @param  {t} testing.T 单元测试
 */
func TestBubbleSort1(t *testing.T) {
	//初始化创建一个数组切片
	values := []int{5, 4, 3, 2, 1}
	arr := [5]int{1, 2, 3, 4, 5}
	BubbleSort(values)
	for i, _ := range values {
		if values[i] != arr[i] {
			t.Error("BubbleSort", values, "结果应该为", arr)
		}
	}
}

/*
 * 存在相同数据单元测试
 * @param  {t} testing.T 单元测试
 */
func TestBubbleSort2(t *testing.T) {
	//初始化创建一个数组切片
	values := []int{5, 5, 3, 2, 1}
	arr := [5]int{1, 2, 3, 5, 5}
	BubbleSort(values)
	for i, _ := range values {
		if values[i] != arr[i] {
			t.Error("BubbleSort", values, "结果应该为", arr)
		}
	}
}

/*
 * 只有一个数据进行测试
 * @param  {t} testing.T 单元测试
 */
func TestBubbleSort3(t *testing.T) {
	//初始化创建一个数组切片
	values := []int{5}
	arr := [1]int{5}
	BubbleSort(values)
	for i, _ := range values {
		if values[i] != arr[i] {
			t.Error("BubbleSort", values, "结果应该为", arr)
		}
	}
}
