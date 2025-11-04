package basis

// 给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素
func OnlyOneInt(values []int) int {
	//创建map记录values中元素出现的次数
	intMap := make(map[int]int)
	for _, v := range values {
		count := intMap[v] + 1
		intMap[v] = count
	}
	//遍历map，找到出现次数为1的元素，假定我们输入的数组中只有一个元素出现一次，不考虑数组中多个元素出现一次的情况
	for k, v := range intMap {
		if v == 1 {
			return k
		}
	}
	// 输入的数组中不存在出现次数为1的元素
	return -1
}
