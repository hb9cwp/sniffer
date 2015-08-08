// from
//  https://github.com/david415/HoneyBadger/issues/71
// it works with icmp but not tcp.

package main

import (
    "fmt"
    "net"
)

func main() {
    protocol := "icmp"
    netaddr, _ := net.ResolveIPAddr("ip4", "127.0.0.1")
    conn, _ := net.ListenIP("ip4:"+protocol, netaddr)

    buf := make([]byte, 1024)
    numRead, _, _ := conn.ReadFrom(buf)
    fmt.Printf("% X\n", buf[:numRead])
}

