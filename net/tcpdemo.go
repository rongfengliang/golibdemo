package net

import (
	"io"
	"log"
	"net"
	"time"
)

func TimeServer() {
	server, err := net.Listen("tcp", "0.0.0.0:1234")
	if err != nil {
		panic("server start error")
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Print("error")
			continue
		}
		handler(conn)
	}
}
func handler(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
