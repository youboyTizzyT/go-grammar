/**
 * @program: go-grammar
 *
 * @description:
 *
 * @author: Mr.weicong
 *
 * @create: 2019-10-15 16:35
 **/
package main

import "fmt"

/**
map 怎么确定 key 是否存在，如果访问了不存在的 key 会出现什么问题？
*/

func main() {
	intMap := make(map[int]int)
	stringMap := make(map[int]string)
	structMap := make(map[int]struct{})
	floatMap := make(map[int]float64)
	funcMap := make(map[int]func())
	key := 2
	if value, ok := intMap[key]; ok {
		fmt.Println("key存在")
	} else {
		fmt.Println("key不存在, value为空值:", value)
	}
	fmt.Println("直接访问intMap：", intMap[key])
	fmt.Println("直接访问stringMap：", stringMap[key])
	fmt.Println("直接访问structMap：", structMap[key])
	fmt.Println("直接访问floatMap：", floatMap[key])
	fu := funcMap[key]
	fmt.Println("直接访问funcMap：", fu)
	/**
	panic: runtime error: invalid memory address or nil pointer dereference

	空指针
	*/
	//fu()
}

// todo 空 map 的大小是怎么样的?
// todo 空 map（或者 map）在内存中的布局是怎么样的？
// todo 为什么对不存在的 key 的访问不会报错，为什么这么设计？

//https://golang.org/ref/spec#Index_expressions
