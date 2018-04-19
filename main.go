package main

import (
	"net"
	"fmt"
	"os"
)




func main() {
	if len(os.Args) < 3 {
		fmt.Println("not enough arguments")
		return
	}
	addr := os.Args[1]
	port := os.Args[2]
	fmt.Println(fmt.Sprintf("%s:%s", addr, port))
	udpAddr, e := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%s", addr, port))
	if e != nil {
		panic(e)
		return
	}
	ln, e := net.ListenUDP("udp", udpAddr)
	if e != nil {
		panic(e)
	}
	go func () {
		for {
			var buf = make([]byte, 1000)
			n, cAddr, e := ln.ReadFrom(buf)
			if e != nil {
				fmt.Println(e)
				break
			}
			if n > 0 {
				fmt.Println("from:", cAddr.String(), string(buf[:n]))
			}
		}
	}()
	for {
		var r, dst string
		n, e := fmt.Scanf("%s %s", &dst, &r)
		if e != nil {
			fmt.Println(e)
			continue
		}
		if n == 0 {
			continue
		}
		//fmt.Scanln(&dst, &r)
		remoteAddr, e := net.ResolveUDPAddr("udp", dst)
		if e != nil {
			panic(e)
		}
		n, e = ln.WriteToUDP([]byte(r), remoteAddr)
		if e != nil {
			panic(e)
		}
		if n > 0 {
			fmt.Println("write data:", n, r)
		}
	}
}