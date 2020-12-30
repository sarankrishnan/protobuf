package main

import (
  "fmt"
  "net"
  "os"
)

func main() {
   if len(os.Args) != 2 {
     fmt.Println("Syntax %s host:port", os.Args[0])
     os.Exit(1)
   }

   //var buf [512]byte

   service := os.Args[1]

   tcpAddr, _ := net.ResolveTCPAddr("tcp4", service)
   conn, _ := net.DialTCP("tcp", nil, tcpAddr)
   conn.Write(encode(10, "Hello"))
   //conn.Read(buf[0:])
   //decode(buf[0:])
}
