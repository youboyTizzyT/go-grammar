/**
 * @program: go-grammar
 *
 * @description:
 *
 * @author: Mr.weicong
 *
 * @create: 2019-10-15 16:32
 **/
package main

/**
golang time Unmarshal 必须包含时区吗？



必须包含，文档里面有https://golang.org/pkg/time/#Time.UnmarshalJSON，时间必须符合RFC 3339格式。

要自定义Unmarshal，需要实现json.Unmarshaler接口。看这里https://stackoverflow.com/questions/25087960/json-unmarshal-time-that-isnt-in-rfc-3339-format
*/
