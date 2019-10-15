/**
 * @program: go-grammar
 *
 * @description:
 *
 * @author: Mr.weicong
 *
 * @create: 2019-10-15 17:05
 **/
package main

import "errors"

/**
golang中这两种方法的如何取舍选择，之间有什么区别func test(&post) error{} 和func test() (post Post, err error){}
学习golang期间看到有些内部方法是通过形参形式给参数赋值的，比如json.Unmarshal定义为
*/
func Unmarshal(data []byte, v interface{}) error {
	// Check for well-formedness.
	// Avoids filling out half a data structure
	// before discovering a JSON syntax error.
	//var d decodeState
	//err := checkValid(data, &d.scan)
	//if err != nil {
	//	return err
	//}
	//
	//d.init(data)
	return errors.New("")
}

/**
为何不定义成如下这种方式
*/
func Unmarshal1(data []byte) (v interface{}, err error) {
	//	……
	return nil, nil
}

/**
func test(&post) error给与了一个具体的对象，通常用于需改等操作，例如 func UpdatePostInfo(&post) error
func test() (post Post, err error)没有给与具体的对象，通常用于new，例如func NewPostInfo() (post Post, err error)
*/
