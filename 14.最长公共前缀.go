package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	sort.Strings(strs)
	fmt.Printf("min = %s , max = %s", strs[0], strs[len(strs)-1])

	minStrs, maxStrs := strs[0], strs[len(strs)-1]
	commonStrs := make([]byte, 0)
	for i := 0; i < len(minStrs); i++ {
		if minStrs[i] == maxStrs[i] {
			commonStrs = append(commonStrs, minStrs[i])
		} else {
			break
		}
	}
	return string(commonStrs)
}

// @lc code=end
