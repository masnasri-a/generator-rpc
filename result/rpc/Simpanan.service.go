
package service

import (
    "context"
    "fmt"
 
    "gorm.io/gorm"
)

type SimpananService struct {
    pb.UnimplementedLoanServiceServer
	Db *gorm.DB
}

func (s *SimpananService) SimpananCreate(ctx context.Context, param *pb.SimpananCreateRequest) (*pb.DefaultSimpananResponse, error) {
    repo := repository.InitSimpananRepository(s.Db)
    response, err := repo.SimpananCreate(param)
    if err != nil {
        return nil, err
    }
    return resp, nil
}
        
func (s *SimpananService) SimpananDetail(ctx context.Context, param *pb.SimpananDetailRequest) (*pb.SimpananDetailResponse, error) {
    repo := repository.InitSimpananRepository(s.Db)
    response, err := repo.SimpananDetail(param)
    if err != nil {
        return nil, err
    }
    return resp, nil
}
        
func (s *SimpananService) SimpananList(ctx context.Context, param *pb.SimpananListFilterRequest) (*pb.SimpananListResponse, error) {
    repo := repository.InitSimpananRepository(s.Db)
    response, err := repo.SimpananList(param)
    if err != nil {
        return nil, err
    }
    return resp, nil
}
        
func (s *SimpananService) SimpananUpdate(ctx context.Context, param *pb.SimpananUpdateRequest) (*pb.DefaultSimpananResponse, error) {
    repo := repository.InitSimpananRepository(s.Db)
    response, err := repo.SimpananUpdate(param)
    if err != nil {
        return nil, err
    }
    return resp, nil
}
        
func (s *SimpananService) SimpananDelete(ctx context.Context, param *pb.SimpananDeleteRequest) (*pb.DefaultSimpananResponse, error) {
    repo := repository.InitSimpananRepository(s.Db)
    response, err := repo.SimpananDelete(param)
    if err != nil {
        return nil, err
    }
    return resp, nil
}
        
func (s *SimpananService) SimpananKasBalance(ctx context.Context, param *pb.SimpananKasBalanceRequest) (*pb.SimpananKasBalanceResponse, error) {
    repo := repository.InitSimpananRepository(s.Db)
    response, err := repo.SimpananKasBalance(param)
    if err != nil {
        return nil, err
    }
    return resp, nil
}
        
func (s *SimpananService) SimpananUserBalance(ctx context.Context, param *pb.SimpananUserBalanceRequest) (*pb.SimpananUserBalanceResponse, error) {
    repo := repository.InitSimpananRepository(s.Db)
    response, err := repo.SimpananUserBalance(param)
    if err != nil {
        return nil, err
    }
    return resp, nil
}
        