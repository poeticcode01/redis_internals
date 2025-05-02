package main

import (
	"io"
	"log"
	"net"
)

func readCommand(c net.Conn) (string, error) {
	var buf []byte = make([]byte, 512)
	n, err := c.Read(buf[:])
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}
func respond(cmd string, c net.Conn) error {
	if _, err := c.Write([]byte(cmd)); err != nil {
		return err
	}
	return nil
}
func main() {
	log.Println("Spawning tcp echo server:")
	listener, err := net.Listen("tcp", "0.0.0.0:7379")
	if err != nil {
		panic(err)

	}
	var con_clients int = 0
	for {
		c, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		con_clients += 1
		log.Println("concurrent clients", con_clients)
		for {
			cmd, err := readCommand(c)
			if err != nil {
				c.Close()
				con_clients -= 1
				log.Println("client disconnected", c.RemoteAddr(), "concurrent clients", con_clients)
				if err == io.EOF {
					break
				}
			}
			log.Println("command", cmd)
			if err = respond(cmd, c); err != nil {
				log.Print("err write:", err)
			}
		}

	}
}
