package main

import (
	"encoding/json"
	"fmt"
	"geerpc"
	// "github.com/go-basic/uuid"
	"geerpc/codec"
	"log"
	"net"
	"time"
)

func startServer(addr chan string) {
	// pick a free port
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	addr <- l.Addr().String()
	geerpc.Accept(l)
}

func main() {
	log.SetFlags(0)
	addr := make(chan string)
	go startServer(addr)

	// in fact, following code is like a simple geerpc client
	conn, _ := net.Dial("tcp", <-addr)
	defer func() { _ = conn.Close() }()

	time.Sleep(time.Second)
	// send options
	_ = json.NewEncoder(conn).Encode(geerpc.DefaultOption)
	cc := codec.NewGobCodec(conn)
	// send request & receive response
	for i := 0; i < 5; i++ {
		h := &codec.Header{
			ServiceMethod: "Foo.Sum",
			Seq:           uint64(i),
		}
		// go server.handleRequest(cc, req, sending, wg)
		_ = cc.Write(h, fmt.Sprintf("geerpc req %v", time.Now().Nanosecond()))
		_ = cc.Write(h, fmt.Sprintf("geerpc req %v", time.Now().Nanosecond()))

		_ = cc.Write(h, fmt.Sprintf("geerpc req %v", time.Now().Nanosecond()))
		_ = cc.Write(h, fmt.Sprintf("geerpc req %v", time.Now().Nanosecond()))

		_ = cc.ReadHeader(h)
		var reply string
		_ = cc.ReadBody(&reply)
		log.Println("reply:", reply)

		_ = cc.ReadHeader(h)
		_ = cc.ReadBody(&reply)
		log.Println("reply:", reply)

		_ = cc.ReadHeader(h)
		_ = cc.ReadBody(&reply)
		log.Println("reply:", reply)

		_ = cc.ReadHeader(h)
		_ = cc.ReadBody(&reply)
		log.Println("reply:11", reply)

		
	}
	// for {
	// 	var reply string
	// 	h := &codec.Header{}
	// 	err := cc.ReadHeader(h)
	// 	fmt.Println(time.Now())
	// 	if err != nil {
	// 		break
	// 	}
	// 	err = cc.ReadBody(&reply)
	// 	if err != nil {
	// 		break
	// 	}
	// 	log.Println("reply:", reply, ",答复次数")
	// }
}
