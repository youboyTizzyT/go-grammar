/**
 * @program: go-grammar
 *
 * @description:
 *
 * @author: Mr.weicong
 *
 * @create: 2019-10-18 15:39
 **/
package main

/**
给定一个没有重复数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func getNum(num int) int {
	if num == 1 {
		return 1
	}
	return num * getNum(num-1)
}
func permute(nums []int) [][]int {
	l := len(nums)
	for i := 0; i < getNum(l); i++ {

	}
}
