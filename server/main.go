package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"server/pb/number"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	Address string = "localhost:6001"
)

type server struct{}

//简单rpc模式
func (s *server) AddNumberAsync(ctx context.Context, request *number.AddNumberRequest) (*number.AddNumberResponse, error) {
	num1 := request.First + request.Second
	res := &number.AddNumberResponse{
		Num: num1,
	}
	return res, nil
}

//双向流rpc模式
func (s *server) TestStreamSendRequestAsync(srv number.OprationNumberService_TestStreamSendRequestAsyncServer) error {
	n := 1
	for {
		//接受客户端的消息
		request, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			fmt.Println(err)
		}
		//发送信息
		err = srv.Send(&number.StreamResponse{
			Answer: "服服服服服服",
		})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Server_" + request.Question + strconv.Itoa(n))
		n++

	}
	return nil
}

func main() {
	//1.监听本地端口
	listener, err := net.Listen("tcp", Address)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Grpc服务端启动......")
	//2.创建Grpc
	ser := grpc.NewServer()
	//3.在Grpc服务端注册自定义的服务
	number.RegisterOprationNumberServiceServer(ser, &server{})
	//4.在Grpc服务器注册服务器反射服务
	reflection.Register(ser)

	/***
	**  5.Serve方法接收监听的端口,每到一个连接创建一个ServerTransport和server的grroutine
	**	这个goroutine读取GRPC请求,调用已注册的处理程序进行响应
	***/
	err = ser.Serve(listener)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Grpc服务端启动......")
}
