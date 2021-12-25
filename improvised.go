package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/pires/go-proxyproto"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

var ctx = context.Background()

func main() {
	fmt.Println("Starting Improvised !")

	port := 8080
	if len(os.Args) >= 2 {
		number, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Panic("Parameter 1 is not a number", err)
		}
		port = number
	}

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	checkError(err)
	log.Printf("Listening to " + strconv.Itoa(port))

	source := Redis{
		Options: &redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	}
	err = source.init()
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().Unix())

	if err != nil {
		panic("connection error:" + err.Error())
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept Error:", err)
			continue
		}

		dstServer := source.getRandom()
		if dstServer == nil {
			closeConnection(conn)
			continue
		}

		dst, err := net.ResolveTCPAddr("tcp", *dstServer)
		if err != nil {
			fmt.Println("Resolve Error:", err)
			continue
		}
		copyConn(conn, dst)
	}
}

func copyConn(src net.Conn, dstAddr *net.TCPAddr) {
	dst, err := net.DialTCP("tcp", nil, dstAddr)
	if err != nil {
		panic("Dial Error:" + err.Error())
	}

	srcRemoteAddr := src.RemoteAddr().(*net.TCPAddr)
	destRemoteAddr := dst.RemoteAddr().(*net.TCPAddr)

	header := &proxyproto.Header{
		Version:           2,
		Command:           proxyproto.PROXY,
		TransportProtocol: proxyproto.TCPv4,
		SourceAddr:        srcRemoteAddr,
		DestinationAddr:   destRemoteAddr,
	}

	_, err = header.WriteTo(dst)
	checkError(err)

	done := make(chan struct{})

	go func() {
		defer closeConnection(dst)
		defer closeConnection(src)

		_, _ = io.Copy(dst, src)

		done <- struct{}{}
	}()

	go func() {
		defer closeConnection(dst)
		defer closeConnection(src)

		_, _ = io.Copy(src, dst)
		done <- struct{}{}
	}()

	<-done
	<-done
}

var ErrNetClosing = errors.New("use of closed network connection")

func IsErrNetClosing(err error) bool {
	if e, ok := err.(*net.OpError); ok {
		return e.Err.Error() == ErrNetClosing.Error()
	}
	return false
}

func closeConnection(conn net.Conn) {
	err := conn.Close()

	if err != nil {
		if !IsErrNetClosing(err) {
			fmt.Println(err)
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
