package main

import (
  "context"
  "fmt"
  "log"
  pb "../../pb"
  "google.golang.org/grpc"
)

func main() {
  conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
  if err != nil {
    log.Fatal("client connection error:", err)
  }
  defer conn.Close()
  client := pb.NewSearchServiceClient(conn)
  request := &pb.SearchRequest { Query: "sample.query" }
  res, err := client.Search(context.TODO(), request)
  fmt.Printf("result: %#v \n", res)
  fmt.Printf("error: %#v \n", err)
}
