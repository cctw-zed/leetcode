package backtracking

func partition(s string) [][]string {
	res := [][]string{}
	subStr := []string{}

	var dfs func(index int)
	dfs = func(index int) {
		if index == len(s) {
			res = append(res, append([]string(nil), subStr...))
			return
		}

		for i := index; i < len(s); i++ {
			if isPalindrome(s[index:i]) {
				subStr = append(subStr, s[index:i])
				dfs(i)
				subStr = subStr[:len(subStr)-1]
			}
		}
	}

	dfs(0)
	return res
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i] {
			return false
		}
	}
	return true
}
