package backtracking

func permute(nums []int) [][]int {

	res := [][]int{}
	path := make([]int, len(nums))		
	visited := make([]int, len(nums))	// 初始化 全部没有访问过

	var dfs func(index int)
	dfs = func(index int) {
		if index == len(nums) {
			// 完成了一个组合，直接将 path 添加到结果
			res = append(res, append([]int(nil), path...)) // 避免共享引用，创建一个新的切片
			return
		}

		for i:=0; i<len(nums); i++ {
			if visited[i] == 1 {
				continue
			}
			visited[i] = 1
			dfs(index+1)
			visited[i] = 0
		}
	}

	for i := 0; i < len(nums); i++ {
		dfs(i)
	}

	return res
}