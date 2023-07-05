// server.go

package main

import (
	"fmt"
	"github.com/songgao/packets/ethernet"
	"github.com/songgao/water"
	"log"
	"net"
)

type Server struct {
	listener net.Listener
	config   water.Config
	ifce     *water.Interface
}

func NewServer(address string) (*Server, error) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	config := water.Config{
		DeviceType: water.TAP,
	}
	config.Name = "O_O"

	ifce, err := water.New(config)
	if err != nil {
		log.Fatal(err)
	}

	return &Server{
		listener: listener,
		config:   config,
		ifce:     ifce,
	}, nil
}

func (s *Server) Start() {
	defer s.listener.Close()

	fmt.Println("等待客户端连接...")

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("无法接受客户端连接:", err)
			continue
		}

		fmt.Println("客户端已连接:", conn.RemoteAddr().String())

		go s.readFromClient(conn) // 启动处理客户端连接的 goroutine
	}
}

//func (s *Server) handleClient(conn net.Conn) {
//	defer conn.Close()
//
//	go s.readFromClient(conn) // 启动从客户端读取数据的 goroutine
//
//	scanner := bufio.NewScanner(conn)
//	for scanner.Scan() {
//		message := scanner.Text()
//		fmt.Println("从标准输入读取到数据:", message)
//	}
//}

func (s *Server) readFromClient(conn net.Conn) {
	buffer := make([]byte, 1500)
	for {
		n, errR := conn.Read(buffer)
		if errR != nil {
			fmt.Println("读取数据失败:", err)
			break
		}
		frame := ethernet.Frame(buffer[:n])
		n, errE := ifce.Write([]byte(frame))
		// n, errE := ifce.Write(buffer[:n])
		if errE != nil {
			log.Fatal(err)
		}
		log.Printf("Dst: %s\n", frame.Destination())
		log.Printf("Src: %s\n", frame.Source())
		log.Printf("Ethertype: % x\n", frame.Ethertype())
		log.Printf("Payload: % x\n", frame.Payload())
		// fmt.Println("收到客户端数据:", message)
	}
}
