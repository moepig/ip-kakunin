package main

import (
	"fmt"
	"log"
	"net"
	"os"
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
	port := getEnvWithDefault("IP_KAKUNIN_TCP_PORT", "8000")
	tcpListener, err := net.Listen("tcp", ":"+port)
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
	port := getEnvWithDefault("IP_KAKUNIN_TCP_PORT", "8001")
	udpAddr, err := net.ResolveUDPAddr("udp", ":"+port)
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

func getEnvWithDefault(key string, defaultValue string) string {
	ret := os.Getenv(key)
	if ret == "" {
		ret = defaultValue
	}

	return ret
}
