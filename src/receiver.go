package main

import (
  "fmt"
  "log"
  "net"
  "os"
)

func handleRequest(conn net.Conn){
   var buf [20]byte
   var rcvl int
   rcvl, err := conn.Read(buf[0:])
   if err != nil {
     log.Fatalln("Failed to read from socket:", err)
   }
   //It's important that we just unmarshal only the byte stream
   //received. Any padded 0's will make the unmarshal fail
   decode(buf[0:rcvl])
}

func main() {
   if len(os.Args) != 2 {
     fmt.Println("Syntax %s host:port", os.Args[0])
     os.Exit(1)
   }

   service := os.Args[1]

   sock, _ := net.Listen("tcp4", service)
   defer sock.Close()
   for {
     conn, _ := sock.Accept()
     handleRequest(conn) 
   }
}
