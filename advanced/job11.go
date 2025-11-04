package advanced

func Job11(param []*int) {
	for i := 0; i < len(param); i++ {
		*param[i] = *param[i] * 2
	}
}
