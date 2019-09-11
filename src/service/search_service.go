package service

import (
  "context"
  "errors"
  "fmt"
  pb "./../../pb"
)

type SearchService struct {}

func (s *SearchService) Search(ctx context.Context, request *pb.SearchRequest) (*pb.SearchResponse, error) {
  if request.Query == "" {
    return nil, errors.New("No query is passed")
  }
  fmt.Printf("result: %#v \n", request)
  return &pb.SearchResponse {
    Result: request.Query + " is passed",
  }, nil
}
