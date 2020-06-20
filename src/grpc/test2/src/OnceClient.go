/**
 * @program: work
 *
 * @description:
 *
 * @author: Mr.weicong
 *
 * @create: 2020-06-20 11:15
 **/
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"haidao/backend/src/grpc/test2/src/pb"
	"log"
	"time"
)

func main() {
	// 建立连接到gRPC服务
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(1)*time.Second)
	conn, err := grpc.DialContext(ctx, "127.0.0.1:8028", grpc.WithInsecure())
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
	// 模拟请求数据
	res := "test123"
	req := &pb.Req{JsonStr: res}
	start := time.Now().UnixNano() / 1e6
	for i := 0; i < 100000; i++ {
		// 调用gRPC接口
		ctx1, _ := context.WithTimeout(context.Background(), time.Duration(6)*time.Second)
		_, err := t.Once(ctx1, req)
		if err != nil {
			log.Printf("could not greet: %v", grpc.ErrorDesc(err))
		}
	}
	fmt.Println(time.Now().UnixNano()/1e6 - start)
}
