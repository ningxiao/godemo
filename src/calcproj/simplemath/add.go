// add.go
package simplemath

/*
 * 计算两个值相加同类型 函数名称大写包外可见
 * @param  {int} a 传入名称
 * @param  {int} b 传入名称
 * @return {int} sum 相加结果
 */
func Add(a, b int) int {
	if a < 50 { //如果a小于50的时候调用内部函数执行翻倍处理
		a = sum(a)
	}
	return a + b
}

/*
 * 计算两个值相加同类型 函数名称小写包外不可见
 * @param  {int} a 传入名称
 * @return {int} sum 相乘结果
 */
func sum(a int) int {
	return a * 2
}
