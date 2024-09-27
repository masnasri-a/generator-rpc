
package service

import (
    "context"
)

type SimpananService struct {
    Client pb.SimpananServiceClient
    Ctx    context.Context
}

func InitSimpananService(client pb.SimpananServiceClient, ctx context.Context) *SimpananService {
    return &SimpananService{
        Client: client,
        Ctx:    ctx,
    }
}


func (s *SimpananService) SimpananCreate(ctx context.Context, in *pb.SimpananCreateRequest) (*pb.DefaultSimpananResponse, error) {
    resp, err := s.Client.SimpananCreate(s.Ctx, in)
    if err != nil {
        return nil, err
    }
    return resp, nil
}

        
func (s *SimpananService) SimpananDetail(ctx context.Context, in *pb.SimpananDetailRequest) (*pb.SimpananDetailResponse, error) {
    resp, err := s.Client.SimpananDetail(s.Ctx, in)
    if err != nil {
        return nil, err
    }
    return resp, nil
}

        
func (s *SimpananService) SimpananList(ctx context.Context, in *pb.SimpananListFilterRequest) (*pb.SimpananListResponse, error) {
    resp, err := s.Client.SimpananList(s.Ctx, in)
    if err != nil {
        return nil, err
    }
    return resp, nil
}

        
func (s *SimpananService) SimpananUpdate(ctx context.Context, in *pb.SimpananUpdateRequest) (*pb.DefaultSimpananResponse, error) {
    resp, err := s.Client.SimpananUpdate(s.Ctx, in)
    if err != nil {
        return nil, err
    }
    return resp, nil
}

        
func (s *SimpananService) SimpananDelete(ctx context.Context, in *pb.SimpananDeleteRequest) (*pb.DefaultSimpananResponse, error) {
    resp, err := s.Client.SimpananDelete(s.Ctx, in)
    if err != nil {
        return nil, err
    }
    return resp, nil
}

        
func (s *SimpananService) SimpananKasBalance(ctx context.Context, in *pb.SimpananKasBalanceRequest) (*pb.SimpananKasBalanceResponse, error) {
    resp, err := s.Client.SimpananKasBalance(s.Ctx, in)
    if err != nil {
        return nil, err
    }
    return resp, nil
}

        
func (s *SimpananService) SimpananUserBalance(ctx context.Context, in *pb.SimpananUserBalanceRequest) (*pb.SimpananUserBalanceResponse, error) {
    resp, err := s.Client.SimpananUserBalance(s.Ctx, in)
    if err != nil {
        return nil, err
    }
    return resp, nil
}

        