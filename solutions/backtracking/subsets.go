package backtracking

func subsets(nums []int) [][]int {
	res := [][]int{}
	path := []int{}

	var dfs func(index int)
	dfs = func(index int) {
		if index == len(nums) {
			res = append(res, append([]int(nil), path...))
			return
		}

		// 不选择当前数字
		dfs(index + 1)

		// 选择当前数字
		path = append(path, nums[index])
		dfs(index + 1)
		path = path[:len(path)-1]
	}

	dfs(0)
	return res
}

func subsets2(nums []int) [][]int {
	res := [][]int{}
	path := []int{}

	var dfs func(index int)
	dfs = func(index int) {
		res = append(res, append([]int(nil), path...))

		for i := index; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(0)
	return res
}
