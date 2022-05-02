package main

import "strings"

/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	words := strings.Split(strings.TrimSpace(s), " ")

	return len(words[len(words)-1])
}

// @lc code=end
