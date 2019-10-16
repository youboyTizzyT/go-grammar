/**
 * @program: go-grammar
 *
 * @description:
 *
 * @author: Mr.weicong
 *
 * @create: 2019-10-15 17:34
 **/
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	lenNums := len(nums)
	if lenNums < 2 {
		if lenNums == 1 {
			if nums[0] == target {
				return []int{0}
			}
		} else {
			return []int{}
		}
	}
	for i := 0; i < lenNums-1; i++ {
		for j := i + 1; j < lenNums; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSum2(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for k, v := range nums {
		numsMap[v] = k
	}
	for k, v := range nums {
		otherNum := target - v
		index, ok := numsMap[otherNum]
		if ok {
			if k == index {
				continue
			}
			return []int{k, index}
		}
	}
	return []int{}
}
func twoSum3(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for k, v := range nums {
		otherNum := target - v
		if index, ok := numsMap[otherNum]; ok {
			return []int{k, index}
		}
		numsMap[v] = k
	}
	return []int{}
}
func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum2([]int{3, 2, 4}, 6))
	fmt.Println(twoSum3([]int{3, 2, 4}, 6))
}
