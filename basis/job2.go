package basis

func IsValid(value string) bool {
	// 存储右括号与对应左括号的映射
	parenMap := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	valueStrs := []string{}
	for i := 0; i < len(value); i++ {
		//将左括号存入切片中
		if "(" == string(value[i]) || "[" == string(value[i]) || "{" == string(value[i]) {
			valueStrs = append(valueStrs, string(value[i]))
		} else {
			//右括号出现时如果左括号还未出现则直接返回false
			if len(valueStrs) == 0 {
				return false
			}
			//判断右括号是否对应
			if parenMap[string(value[i])] != valueStrs[len(valueStrs)-1] {
				return false
			}
			//删除对应的左括号继续匹配
			valueStrs = valueStrs[:len(valueStrs)-1]
		}
	}

	return len(valueStrs) == 0
}
