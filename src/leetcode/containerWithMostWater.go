/**
 * @program: go-grammar
 *
 * @description:
 *
 * @author: Mr.weicong
 *
 * @create: 2019-10-19 15:44
 **/
package main

import (
	"fmt"
	"math"
	"sort"
)

/**
给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/container-with-most-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maxArea(height []int) int {
	l := len(height)
	if l <= 1 {
		return 0
	}
	copyHeight := make([]int, l)
	copy(copyHeight, height)
	sort.Ints(copyHeight)
	var keyHeight []int
	for _, v := range copyHeight {
		if index := isInArray(height, v); index != -1 {
			keyHeight = append(keyHeight, index)
			height = append(height[:index], height[index+1:]...)
		}
	}
	max := 0
	for k, v := range keyHeight {
		hight1 := height[v]
		for j := k + 1; j < l; j++ {
			hight2 := height[j]
			hight := If(hight1, hight2, hight1 > hight2)
			wight := int(math.Abs(float64(v - keyHeight[j])))
			fmt.Println("wight:", wight)
			fmt.Println("hight:", hight)
			ret := wight * hight
			if max < ret {
				max = ret
			}
		}
	}
	return 0
}
func isInArray(ints []int, int2 int) int {
	for k, v := range ints {
		if v == int2 {
			return k
		}
	}
	return -1
}
func If(a int, b int, bool2 bool) int {
	if bool2 {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 9, 1, 1, 8}))
}
