/*
 from
  https://github.com/david415/HoneyBadger/issues/71

 to run:
  $ sudo go run tinysniffer.go
  $ ping localhost
  $ nc -u localhost 222
    then type anything & press <return>

 from Comments: it works with ICMP & UDP, but not with TCP!?
  $ telnet localhost 19 

*/

package main

import (
    "fmt"
    "net"
)

func main() {
    protocol := "icmp"
//    protocol := "tcp"
    netaddr, _ := net.ResolveIPAddr("ip4", "127.0.0.1")
    conn, _ := net.ListenIP("ip4:"+protocol, netaddr)

    buf := make([]byte, 1024)
    numRead, _, _ := conn.ReadFrom(buf)
    fmt.Printf("% X\n", buf[:numRead])
}

