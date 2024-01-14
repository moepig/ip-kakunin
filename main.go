package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go tcpListener()
	go udpListener()

	wg.Wait()
}

func tcpListener() {
	tcpListener, err := net.Listen("tcp", "localhost:8088")
	if err != nil {
		log.Fatal(err)
	}

	defer tcpListener.Close()

	fmt.Println("TCP Server listening on ", tcpListener.Addr().String())

	for {
		tcpConn, err := tcpListener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		tcpHandler(tcpConn)
	}
}

func tcpHandler(conn net.Conn) error {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return err
	}

	fmt.Printf("[TCP] data from %s: %s\n", conn.RemoteAddr().String(), buffer[:n])

	return nil
}

func udpListener() {
	udpAddr, err := net.ResolveUDPAddr("udp", ":8089")
	if err != nil {
		log.Fatal(err)
	}

	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer udpConn.Close()

	fmt.Println("UDP Server listening on", udpConn.LocalAddr().String())

	for {
		udpHandler(udpConn)
	}
}

func udpHandler(conn *net.UDPConn) error {
	buffer := make([]byte, 1024)
	n, addr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		return err
	}

	fmt.Printf("[UDP] data from %s: %s\n", addr.String(), buffer[:n])

	return nil
}
