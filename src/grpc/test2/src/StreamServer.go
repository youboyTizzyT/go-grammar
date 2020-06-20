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
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"haidao/backend/src/grpc/test2/src/pb"
	"haidao/backend/src/grpc/test2/src/server"
	"log"
	"net"
	"time"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:8029") //监听所有网卡8028端口的TCP连接
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}
	s := grpc.NewServer(grpc.ConnectionTimeout(50 * time.Second)) //创建gRPC服务

	/**注册接口服务
	 * 以定义proto时的service为单位注册，服务中可以有多个方法
	 * (proto编译时会为每个service生成Register***Server方法)
	 * 包.注册服务方法(gRpc服务实例，包含接口方法的结构体[指针])
	 */
	pb.RegisterWaiterServer(s, &server.Server{})
	/**如果有可以注册多个接口服务,结构体要实现对应的接口方法
	 * user.RegisterLoginServer(s, &server{})
	 * minMovie.RegisterFbiServer(s, &server{})
	 */
	// 在gRPC服务器上注册反射服务
	reflection.Register(s)
	// 将监听交给gRPC服务处理

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
