/**
 * @program: work
 *
 * @description:
 *
 * @author: Mr.weicong
 *
 * @create: 2020-06-01 10:56
 **/
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"haidao/backend/src/grpc/test1/src/pb"
	"log"
	"time"
)

func main() {
	conn, err := grpc.DialContext(context.Background(), "127.0.0.1:8028", grpc.WithInsecure())
	fmt.Println(conn.GetState())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	fmt.Println(conn.GetState())
	// 函数结束时关闭连接
	defer conn.Close()
	fmt.Println(conn.GetState())
	// 创建Waiter服务的客户端
	t := pb.NewWaiterClient(conn)
	fmt.Println(conn.GetState())
	for true {
		fmt.Println(conn.GetState())
		// 模拟请求数据
		res := "test123"
		// 调用gRPC接口
		ctx1, _ := context.WithTimeout(context.Background(), time.Duration(6)*time.Second)
		tr, err := t.DoMD5(ctx1, &pb.Req{JsonStr: res})

		if err != nil {
			log.Printf("could not greet: %v", grpc.ErrorDesc(err))
		}
		log.Printf("服务端响应: %s\n\n", tr)
		time.Sleep(1 * time.Second)
	}
}
