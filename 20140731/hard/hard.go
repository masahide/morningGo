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
			case *net.IPAddr:
				if !ip.IP.IsLoopback() {
					pretty.Printf("%s:%s\n", i.Name, ip)
				}
				/*
					//↓の書き方はコンパイルエラーに
					case *net.IPNet, *net.IPAddr:
						if !ip.IP.IsLoopback() {  // > ip.IP undefined (type net.Addr has no field or method IP)
							pretty.Printf("%s:%s\n", i.Name, ip)
						}
				*/
			}
		}
	}
}
