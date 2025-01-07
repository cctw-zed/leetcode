package backtracking

import (
	"fmt"
	"testing"

	"github.com/cctw-zed/leetcode/solutions/backtracking"
)

func TestCombinations(t *testing.T) {
	res := backtracking.Combine(4, 2)
	fmt.Println(res)
}
