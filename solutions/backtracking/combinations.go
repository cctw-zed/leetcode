package backtracking

func Combine(n int, k int) [][]int {
	res := [][]int{}

	// 使用一个路径数组，大小为 k，可以容纳所有的组合
	path := make([]int, k)

	var dfs func(begin, index int)
	dfs = func(begin, index int) {
		if index == k {
			// 完成了一个组合，直接将 path 添加到结果
			res = append(res, append([]int(nil), path...)) // 避免共享引用，创建一个新的切片
			return
		}

		// 从当前 begin 开始，选择后续的数字
		for i := begin; i <= n+1+index-k; i++ {
			path[index] = i   // 直接在 path 中设置当前选择的数字
			dfs(i+1, index+1) // 递归，下一层选择，更新 index
		}
	}

	dfs(1, 0) // 从 1 开始，index 从 0 开始
	return res
}
