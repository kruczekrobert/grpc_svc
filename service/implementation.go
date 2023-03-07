package service

import (
	context "context"
)

type Service struct {
	UnimplementedServiceServer
}

func (s *Service) DoSomething(ctx context.Context, in *SomeRequest) (*SomeResponse, error) {
	return &SomeResponse{
		Message: in.Message,
	}, nil
}
