package main

import (
  "log"
  "net"
  pb "../../pb"
  "../service"
  "google.golang.org/grpc"
)

func main() {
  listenPort, err := net.Listen("tcp", ":19003")
  if err != nil {
    log.Fatalln(err)
  }
  server := grpc.NewServer()
  searchService := &service.SearchService {}
  pb.RegisterSearchServiceServer(server, searchService)
  server.Serve(listenPort)
}
