package main

import (
	"fmt"
	"log"
	"net"
	"node"
	"os"
)

const COPY_SOCKET = "/tmp/copy.sock"

func main() {
	logf, err := os.OpenFile("/home/arm/copy.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer logf.Close()

	log.SetFlags(log.LstdFlags)
	log.SetOutput(logf)
	root, ok := os.LookupEnv("HOME")
	if ok {
		log.Printf("root path to %s\n", root)
	}
	fmt.Println("server start")
	ln, err := net.Listen("unix", COPY_SOCKET)
	if err != nil {
		if err = os.Remove(COPY_SOCKET); err != nil {
			log.Fatal(err)
		}
		// try again
		ln, err = net.Listen("unix", COPY_SOCKET)
		if err != nil {
			log.Fatal(err)
		}
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		//defer conn.Close()
		// sub function is responsable for close
		if node.Host(conn, root) {
			break
		}
	}
}
