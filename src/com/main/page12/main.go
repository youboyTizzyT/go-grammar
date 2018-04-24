package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}
// map 映射键到值。

// 定义map
var m map[string]Vertex
var n = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	// 如果顶级的类型只有类型名的话，可以在文法的元素中省略键名。
	"Google":{
		37.42202, -122.08408,
	},
}
func main() {
	// map 在使用之前必须用 make 而不是 new 来创建；
	// 值为 nil 的 map 是空的，并且不能赋值。
	m=make(map[string]Vertex)
	m["Bell Labs"]=Vertex{
		40.684333,-15.15431,
	}
	fmt.Println(m["Bell Labs"])
	// 因为n在定义是就有值,那么就不用make
	// 当然也可以继续往里面加
	n["123"]=Vertex{
		123,516.151513,
	}
	// 也可以删除
	delete(n,"Bell Labs")
	fmt.Println(n,len(n))
	// 获取元素.和php有点像~txtx
	fmt.Println(n["123"])
	// 修改元素
	n["123"]=Vertex{
		151.15613,484.1131,
	}

	// 通过双赋值检测某个键存在：  很新颖,也是go语言特性.
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}