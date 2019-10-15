/**
 * @program: go-grammar
 *
 * @description:
 *
 * @author: Mr.weicong
 *
 * @create: 2019-10-15 17:03
 **/
package main

import (
	"html/template"
	"os"
)

func main() {
	//{{ $v1 := "helloworld" }}   //在Go语言的模板中定义一个字符串变量没问题
	//
	//{{ $v2 := [5]int{1,2,3,4,5} }} //但是定义一个数组变量就会报错。。
	//
	//请问可以在Go语言的模板中定义一个数组变量吗?如果可以，那应该怎么定义？

	//https://stackoverflow.com/questions/25012467/golang-templates-how-to-define-array-in-a-variable
	tmpl := `
{{ $slice := mkSlice "a" 5 "b" }}
{{ range $slice }}
     {{ . }}
{{ end }}
`
	funcMap := map[string]interface{}{"mkSlice": mkSlice}
	t := template.New("demo").Funcs(template.FuncMap(funcMap))
	template.Must(t.Parse(tmpl))
	t.ExecuteTemplate(os.Stdout, "demo", nil)
}

func mkSlice(args ...interface{}) []interface{} {
	return args
}
