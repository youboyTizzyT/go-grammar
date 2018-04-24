package main


import "fmt"

func main() {
	// slice, 切片
	// 一个 slice 会指向一个数组，并且包含了长度信息。
	// []T 是一个元素类型为 T 的 slice。

	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n",
			i, p[i])
	}
	// slice 可以重新切片，创建一个新的 slice 值指向相同的数组。
	// 表达式

	// s[lo:hi]
	// 表示从 lo 到 hi-1 的 slice 元素，含两端。因此

	// s[lo:lo]
	// 是空的，而

	// s[lo:lo+1]
	// 有一个元素。
	fmt.Println("p[1:4] ==", p[1:4])

	// 省略下标代表从 0 开始
	fmt.Println("p[:3] ==", p[:3])

	// 省略上标代表到 len(s) 结束
	fmt.Println("p[4:] ==", p[4:])
}