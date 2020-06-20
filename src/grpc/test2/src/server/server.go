/**
 * @program: work
 *
 * @description:
 *
 * @author: Mr.weicong
 *
 * @create: 2020-06-20 11:30
 **/
package server

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"google.golang.org/grpc/status"
	"haidao/backend/src/grpc/test2/src/pb"
	"time"
)

// 业务实现方法的容器
type Server struct{}

func (s *Server) Stream(stream pb.Waiter_StreamServer) error {
	for true {
		req, err := stream.Recv()
		if err != nil {
			fmt.Println(err, status.Convert(err))
			return errors.New("没拿到请求")
		}
		fmt.Println(time.Now())
		//panic("qweqwe")
		_ = stream.Send(&pb.Res{BackJson: "MD5 :" + fmt.Sprintf("%x", md5.Sum([]byte(req.JsonStr)))})
	}
	return nil
}

func (s *Server) Once(ctx context.Context, req *pb.Req) (*pb.Res, error) {
	fmt.Println(time.Now())
	//panic("qweqwe")
	return &pb.Res{BackJson: "MD5 :" + fmt.Sprintf("%x", md5.Sum([]byte(req.JsonStr)))}, nil
}
