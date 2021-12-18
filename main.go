package main

import "zinx/znet"

func main() {

	server := znet.Server{IP: "", Port: 8080}

	server.Start()

	defer server.Stop()

}
