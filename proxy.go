package main

import (
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()
	b := make([]byte,512)
	for {
		size, err := conn.Read(b[0:])
		if err!=nil {
			log.Println("unexpected error")
			break
		}
		log.Printf("recvd %d bytes: %s\n",size,string(b))
		log.Println("writing data...")
		if _, err:=conn.Write(b[0:size]); err !=nil {
			log.Fatalln("unable to write data")
		}
	}
}

func main() {
	listener, err:=net.Listen("tcp",":20080")
	if err!=nil {
		log.Fatalln("unable to bind")
	}
	log.Println("listening on 0.0.0.0:20080")
	for {
		conn, err := listener.Accept()
		log.Println("recvd conn")
		if err !=nil {
			log.Fatalln("unable to accept")
		}
		go echo(conn)
	}
}