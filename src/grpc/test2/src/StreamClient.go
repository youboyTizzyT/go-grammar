/**
 * @program: work
 *
 * @description:
 *
 * @author: Mr.weicong
 *
 * @create: 2020-06-20 11:16
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
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(100)*time.Second)
	conn, err := grpc.DialContext(ctx, "127.0.0.1:8029", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 函数结束时关闭连接
	defer conn.Close()
	// 创建Waiter服务的客户端
	t := pb.NewWaiterClient(conn)
	//_, _ := context.WithTimeout(context.Background(), time.Duration(6)*time.Second)
	stream, err := t.Stream(ctx)
	if err != nil {
		log.Printf("could not greet: %v", grpc.ErrorDesc(err))
	}
	//go func() {
	//	defer func() {
	//		fmt.Println("读结束", time.Now().UnixNano()/1e6)
	//	}()
	//	for i := 0; i >= 0; i++ {
	//		_, err := stream.Recv()
	//		if i == 0 {
	//			fmt.Println("读开始", time.Now().UnixNano()/1e6)
	//		}
	//		if err != nil {
	//			return
	//		}
	//		if i == 99999 {
	//			return
	//		}
	//	}
	//}()
	fmt.Println("写开始", time.Now().UnixNano()/1e6)
	for i := 0; i < 100000; i++ {
		err := stream.Send(&pb.Req{JsonStr: "test123"})
		if err != nil {
			log.Printf("could not greet: %v", grpc.ErrorDesc(err))
			break
		}
		_, err = stream.Recv()
	}
	fmt.Println("写结束", time.Now().UnixNano()/1e6)
	for true {
		time.Sleep(time.Second)
	}
}
