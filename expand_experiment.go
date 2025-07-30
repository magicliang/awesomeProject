package main

// 这个算法的本质是：
// 1. 从一个空的可迭代 [][]T{{}} 切片套切片开始
// 2. 然后用 originTypes 的每个元素进行叉乘
// 3. 每一轮就把这个 result 升级一次
// 4. 所以要在 originTypes 的每一行里准备一个 temp 替换新的 result
// 5. 遍历完整个 originType 切片以后，才实现替换
// 6. 在遍历 originType 内部的时候，从 result 的内部乘出来我们需要的元素，然后作为结果的新行插进 temp 结果里
// 7. 然后完成替换
func expandToCombinations[T any](originTypes [][]T) [][]T {
	// 一定要双 {{}} 才能触发第一次空循环
	result := [][]T{{}}

	for _, originType := range originTypes {
		// 一定要双 {{}} 才能触发第一次空循环
		var temp [][]T
		for _, t := range originType {
			// 遍历 result 以前生成一个 temp，这样就可以用 result 生成 temp
			for _, r := range result {
				newRow := make([]T, len(r))
				copy(newRow, r)
				newRow = append(newRow, t)
				temp = append(temp, newRow)
			}
		}
		result = temp
	}

	return result
}
