package qsort

/*
 * 快排内部方法
 * 相同类型left, right int 可以直接这样声明
 * @param  {values} []int 传入数组切片
 * @param  {left} int 传入数组
 * @param  {right} int 传入数组
 */
func quickSort(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && values[j] >= temp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}
		if values[i] <= temp && i <= p {
			i++
		}

		if i <= p {
			values[p] = values[i]
			p = i
		}
	}
	values[p] = temp
	if p-left > 1 {
		quickSort(values, left, p-1)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
	}
}

/*
 * 快排对外部方法
 * @param  {values} []int 传入数组切片
 */
func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}
