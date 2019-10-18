/**
 * @program: go-grammar
 *
 * @description:
 *
 * @author: Mr.weicong
 *
 * @create: 2019-10-17 15:40
 **/
package main

import (
	"fmt"
	"sort"
)

/**
给定一个无序的数组 nums，将它重新排列成 nums[0] < nums[1] > nums[2] < nums[3]... 的顺序。

示例 1:

输入: nums = [1, 5, 1, 1, 6, 4]
输出: 一个可能的答案是 [1, 4, 1, 5, 1, 6]
示例 2:

输入: nums = [1, 3, 2, 2, 3, 1]
输出: 一个可能的答案是 [2, 3, 1, 3, 1, 2]
说明:
你可以假设所有输入都会得到有效的结果。

进阶:
你能用 O(n) 时间复杂度和 / 或原地 O(1) 额外空间来实现吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/wiggle-sort-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func wiggleSort(nums []int) []int {
	sort.Ints(nums)
	l := len(nums)
	ret := make([]int, l)
	copy(ret, nums)
	for i := 0; i < l; i++ {
		if i < l/2 {
			nums[2*i+1] = ret[i]
		} else {
			nums[2*(i-l/2)] = ret[i]
		}
	}
	return nums
}
func main() {
	fmt.Println(wiggleSort([]int{1, 5, 1, 1, 6, 4}))
}
