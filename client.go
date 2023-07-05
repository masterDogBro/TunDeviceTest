// client.go

package main

import (
	"fmt"
	"github.com/songgao/packets/ethernet"
	"github.com/songgao/water"
	"log"
	"net"
)

type Client struct {
	conn   net.Conn
	config water.Config
	ifce   *water.Interface
}

func NewClient(address string) (*Client, error) {
	conn, err := net.Dial("tcp", address)
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

	return &Client{
		conn:   conn,
		config: config,
		ifce:   ifce,
	}, nil
}

func (c *Client) Start() {
	defer c.conn.Close()

	fmt.Println("已连接到服务器")

	// go c.readFromServer() // 启动从服务器读取数据的 goroutine
	c.writeToServer() // 从标准输入读取用户输入，并发送到服务器
}

//func (c *Client) readFromServer() {
//	buffer := make([]byte, 1024)
//	for {
//		n, err := c.conn.Read(buffer)
//		if err != nil {
//			fmt.Println("读取数据失败:", err)
//			break
//		}
//		message := string(buffer[:n])
//		fmt.Println("收到服务器数据:", message)
//	}
//}

func (c *Client) writeToServer() {

	var frame ethernet.Frame

	for {
		frame.Resize(1500)
		n, errE := ifce.Read([]byte(frame))
		if errE != nil {
			log.Fatal(err)
		}
		frame = frame[:n]
		_, errW := c.conn.Write([]byte(message + "\n"))
		if errW != nil {
			fmt.Println("发送数据失败:", err)
			break
		}
		log.Printf("Dst: %s\n", frame.Destination())
		log.Printf("Src: %s\n", frame.Source())
		log.Printf("Ethertype: % x\n", frame.Ethertype())
		log.Printf("Payload: % x\n", frame.Payload())
	}

	//scanner := bufio.NewScanner(os.Stdin)
	//for scanner.Scan() {
	//	message := scanner.Text()
	//	_, err := c.conn.Write([]byte(message + "\n"))
	//	if err != nil {
	//		fmt.Println("发送数据失败:", err)
	//		break
	//	}
	//}
}
