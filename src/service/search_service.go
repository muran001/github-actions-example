package service

import (
  "context"
  "errors"
  pb "./../../pb"
)

type SearchService struct {}

func (s *SearchService) Search(ctx context.Context, request *pb.SearchRequest) (*pb.SearchResponse, error) {
  if request.Query == "" {
    return nil, errors.New("No query is passed")
  }
  return &pb.SearchResponse {
    Result: request.Query + " is passed",
  }, nil
}
