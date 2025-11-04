package basis

// 给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
// 将大整数加 1，并返回结果的数字数组。
func Job3(params []int) []int {

	//记录下一位是否加一
	next := true
	//标记是否需要在原数组基础上再添加一位
	addOne := false
	for i := len(params) - 1; i >= 0; i-- {
		if next {
			//如果当前位是9则设置当前为为零
			if params[i] == 9 {
				params[i] = 0
				if i == 0 {
					addOne = true
				}
			} else {
				//否则当前位加一并且设置下一位加1标记为false
				params[i]++
				next = false
			}
		}
	}

	//定义返回切片
	result := []int{}
	if addOne {
		result = append(result, 1)
	}
	//遍历params并将其中的值添加到result中
	for _, v := range params {
		result = append(result, v)
	}
	return result

}
