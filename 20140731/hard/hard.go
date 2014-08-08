package main

import (
	"net"

	"github.com/kr/pretty"
)

func main() {
	/* Hard:自分のマシンのIPアドレスを取得し画面表示 */
	interfaces, _ := net.Interfaces()
	//fmt.Printf("%#v\n", interfaces)
	for _, i := range interfaces {
		ips, _ := i.Addrs()
		for _, ip := range ips {
			//pretty.Printf("%# v\n", ip)
			switch ip := ip.(type) {
			default:
				pretty.Printf("unexpected type %t\n", ip)
			case *net.IPNet:
				if !ip.IP.IsLoopback() {
					pretty.Printf("%s:%s\n", i.Name, ip)
				}
				/*
					// Windowsはこうかな？とおもったけど 違うみたいコンパイルエラーに
					// impossible type switch case: ip (type net.Addr) cannot have dynamic type *net.IP (missing Network method)
					case *net.IP:

					if !ip.IsLoopback() {
						pretty.Printf("%s:%s\n", i.Name, ip)
					}
				*/
			}
		}
	}
}
