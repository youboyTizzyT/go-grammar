/**
 * @program: go-grammar
 *
 * @description:
 *
 * @author: Mr.weicong
 *
 * @create: 2019-10-15 19:01
 **/
package main

import (
	"fmt"
)

/**
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	// 最小子集长度
	ret := 1
	right := 2
	left := 0
	for {
		if i := right - left; i > ret && tmpFunc(s[left:right]) {
			ret = i
			right++
		} else {
			left++
			if right-left <= ret {
				right++
			}
		}
		if right > len(s) {
			break
		}
	}
	return ret
}
func tmpFunc(s string) bool {
	a := make(map[int32]int)
	for _, ss := range s {
		if _, ok := a[ss]; ok {
			return false
		}
		a[ss] = 1
	}
	return true
}
func main() {
	fmt.Println(lengthOfLongestSubstring("aab"))
}
