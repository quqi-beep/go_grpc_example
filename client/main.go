package main

import (
	"client/pb/number"
	"context"
	"fmt"
	"io"
	"strconv"

	"google.golang.org/grpc"
)

const Address string = "localhost:6001"

func main() {
	//1.连接服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	//2.连接Grpc
	c := number.NewOprationNumberServiceClient(conn)

	//3.调用方法(简单rpc模式)
	addNumberAsync(c)

	//4.双向流模式
	testStreamSendRequestAsync(c)
}

//简单rpc模式
func addNumberAsync(c number.OprationNumberServiceClient) {
	//创建要发送的结构体
	request := &number.AddNumberRequest{
		First:  10,
		Second: 11,
	}
	res, err := c.AddNumberAsync(context.Background(), request)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	fmt.Println("服务器端返回的结果为：" + strconv.FormatInt(int64(res.Num), 10))
}

//双向流模式
func testStreamSendRequestAsync(c number.OprationNumberServiceClient) {
	stream, err := c.TestStreamSendRequestAsync(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 30; i++ {
		err = stream.Send(&number.StreamRequest{
			Question: "客客客客客",
		})
		if err != nil {
			fmt.Println(err)
		}
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Client_" + res.Answer + "_" + strconv.Itoa(i+1))
	}
	err = stream.CloseSend()
	if err != nil {
		fmt.Println(err)
	}
}
