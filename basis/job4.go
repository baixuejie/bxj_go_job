package basis

import "fmt"

// 给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数
// 考虑 nums 的唯一元素的数量为 k。去重后，返回唯一元素的数量 k。
// nums 的前 k 个元素应包含 排序后 的唯一数字。下标 k - 1 之后的剩余元素可以忽略
func Job4() {
	// 定义有序数字 data
	data := [6]int{1, 1, 2, 2, 3, 3}
	index := 0

	for i := 1; i < len(data); i++ {
		//判断当前数字和有效数字是否相等，不相等为有效，添加到下一个有效数字位置
		if data[i] != data[index] {
			index++
			data[index] = data[i]
		}
	}
	//将无效数据置零
	for i := index + 1; i < len(data); i++ {
		data[i] = 0
	}

	fmt.Println("最终数组为:", data)
	fmt.Println("有效数据个数:", index+1)
}
