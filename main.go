package main

import "fmt"

func main() {
	// server启动
	server, errs := NewServer(":10000")
	if errs != nil {
		fmt.Println("Failed to start server:", err)
		return
	}
	server.Start()

	// client启动
	//client, errs := NewClient("host:10000")
	//if errs != nil {
	//	fmt.Println("Failed to start client:", err)
	//	return
	//}
	//client.Start()

}
