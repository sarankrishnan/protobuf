package main

import (
  "fmt"
  "log"
  "github.com/golang/protobuf/proto"
  pb "example"
)

func decode(in []byte) {
  msg := &pb.ExchangeData{}
  if err := proto.Unmarshal(in, msg); err != nil {
     log.Fatalln("Failed to decode message:", err)
  }
  fmt.Println(msg)
}

func encode(Index uint32, Status string) []byte {
   msg := &pb.ExchangeData{}
   msg.Index = Index
   msg.Status = Status
   out, err := proto.Marshal(msg)
   if err != nil {
     log.Fatalln("Failed to encode message:", err)
   }
   return out
}

