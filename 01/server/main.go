package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 开启rpc监听 -》端口和ip
	listen, _ := net.Listen("tcp", ":1234")
	defer listen.Close()
	// 建立连接
	conn, _ := listen.Accept()
	// 注册服务   =》》 hash表
	// RegisterName （name,recv）
	// （name 服务标识    recv 具体的服务
	rpc.RegisterName("goods", new(Goods))

	jsonrpc.ServeConn(conn)
}

type Goods struct {
	Id   int
	Name string
}
type Params struct {
	Id   int
	Name string
}

func (g *Goods) FindById(args *Params, reply *Goods) error {
	fmt.Println("接收到请求信息 ：", args)
	*reply = *g
	return nil
}

func (g *Goods) GetByIdGoodsName(args *Params, reply *string) error {
	fmt.Println("接收到请求信息 ：", args)
	*reply = g.Name
	return nil
}
