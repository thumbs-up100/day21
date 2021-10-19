package main

import (
	"fmt"
	"net/rpc"
)

type Goods struct {
	Id   int
	Name string
}

type Params struct {
	Id   int
	Name string
}

func main() {
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	// resgoods := &Goods{}
	resgoods := ""
	reqParam := Params{
		Id:   1,
		Name: "请求",
	}
	// 发起请求，调用服务端的方法
	// err := client.Call("Goods.FindById", reqParam, &resgoods)
	err := client.Call("Goods.GetByIdGoodsName", reqParam, &resgoods)
	if err != nil {
		fmt.Println("err : ", err)
	}
	fmt.Println("resgoods : ", resgoods)
}
