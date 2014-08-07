package main

import "fmt"
import "net"

func main() {
	/* Hard:自分のマシンのIPアドレスを取得し画面表示 */
	interfaces, _ := net.Interfaces()
	//fmt.Printf("%#v\n", interfaces)
	for _, i := range interfaces {
		ips, _ := i.Addrs()
		for _, ip := range ips {
			switch ip.String() {
			case "127.0.0.1/8":
			case "::1/128":
			default:
				fmt.Printf("%s:%s\n", i.Name, ip)
			}
		}
	}
}
