package main

import (
	"log"
	"net"
	"os"

	"github.com/EDDYCJY/redis-protocol-example/protocol"
)

const (
	//Address 默认地址
	Address = "127.0.0.1:6379"
	// Network 默认连接方式
	Network = "tcp"
)

//Conn 是连接redis的方法
func Conn(network, address string) (net.Conn, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func main() {
	args := os.Args[1:]
	if len(args) <= 0 {
		log.Fatalf("Os.Args <= 0")
	}
	//获取请求协议
	reqCommand := protocol.GetRequest(args)
	redisConn, err := Conn(Network, Address)
	if err != nil {
		log.Fatalf("Conn err : %v", err)
	}
	defer redisConn.Close()
	// 写入请求内容
	_, err = redisConn.Write(reqCommand)
	if err != nil {
		log.Fatalf("Conn Read err :%v", err)
	}
	// 读取回复
	command := make([]byte, 1024)
	n, err := redisConn.Read(command)
	if err != nil {
		log.Fatalf("Conn Read err:%v", err)
	}
	//处理请求
	reply, err := protocol.GetReply(command[:n])
	if err != nil {
		log.Fatalf("protocol.GetReply err : %v", err)
	}
	//处理后得到的回复内容
	log.Printf("Reply:%v", reply)
	log.Printf("Command: %v", string(command[:n]))
}
