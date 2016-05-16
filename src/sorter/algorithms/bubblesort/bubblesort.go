package bubblesort

/*
 * BubbleSort 冒泡排序 首字母大写包外可见
 * 因为是切片所以数据是引用类型在排序方法修改之后会影响原始数据
 * @param  {values} []int int类型的数组切片
 */
func BubbleSort(values []int) {
	flag := true
	for i, l := 0, len(values)-1; i < l; i++ {
		flag = true
		for j, k := 0, l-i; j < k; j++ {
			if values[j] > values[j+1] {
				//比较完毕之后进行数据交换
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
			}
		}
		if flag == true {
			break
		}
	}
}
