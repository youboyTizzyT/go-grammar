/**
 * @program: go-grammar
 *
 * @description:
 *
 * @author: Mr.weicong
 *
 * @create: 2019-11-19 17:53
 **/
package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/afocus/captcha"
	"image/color"
	"image/png"
	"testing"
)

var cap *captcha.Captcha

func BenchmarkXXX(b *testing.B) {
	cap = captcha.New()
	// 可以设置多个字体 或使用cap.AddFont("xx.ttf")追加
	if err := cap.SetFont("D:/work/go-grammar/src/team2/main/comic.ttf"); err != nil {
		panic(err.Error())
	}
	// 设置验证码大小
	cap.SetSize(128, 64)
	// 设置干扰强度
	cap.SetDisturbance(8)
	// 设置前景色 可以多个 随机替换文字颜色 默认黑色
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	// 设置背景色 可以多个 随机替换背景色 默认白色
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})

	img, _ := cap.Create(4, captcha.NUM)
	//f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	emptyBuff := bytes.NewBuffer(nil) //开辟一个新的空buff
	png.Encode(emptyBuff, img)        //img写入到buff
	emptyBuff.Bytes()
	fmt.Println(base64.StdEncoding.EncodeToString(emptyBuff.Bytes())) //buff转成base64
}
