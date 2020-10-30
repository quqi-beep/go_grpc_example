package main

import (
	"client/pb/number"
	"context"
	"fmt"
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

	//3.创建要发送的结构体
	request := &number.AddNumberRequest{
		First:  10,
		Second: 11,
	}
	//3.调用方法
	res, err := c.AddNumberAsync(context.Background(), request)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	fmt.Println("服务器端返回的结果为：" + strconv.FormatInt(int64(res.Num), 10))
}
