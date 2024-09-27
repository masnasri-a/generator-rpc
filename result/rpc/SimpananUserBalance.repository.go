
package repository

import (
    "gorm.io/gorm"
)

    
type SimpananRepository struct {
    db *gorm.DB,
}

func InitSimpananRepository(db *gorm.DB) *SimpananRepository {
    return &SimpananRepository{
        DB: db
    }
}
    
func (r *SimpananRepository) SimpananCreate(param *pb.SimpananCreateRequest) (*pb.DefaultSimpananResponse, error) {
    // Implement your code here
    return &pb.DefaultSimpananResponse{
        // Implement your code here
    }, nil
}


        
func (r *SimpananRepository) SimpananDetail(param *pb.SimpananDetailRequest) (*pb.SimpananDetailResponse, error) {
    // Implement your code here
    return &pb.SimpananDetailResponse{
        // Implement your code here
    }, nil
}


        
func (r *SimpananRepository) SimpananList(param *pb.SimpananListFilterRequest) (*pb.SimpananListResponse, error) {
    // Implement your code here
    return &pb.SimpananListResponse{
        // Implement your code here
    }, nil
}


        
func (r *SimpananRepository) SimpananUpdate(param *pb.SimpananUpdateRequest) (*pb.DefaultSimpananResponse, error) {
    // Implement your code here
    return &pb.DefaultSimpananResponse{
        // Implement your code here
    }, nil
}


        
func (r *SimpananRepository) SimpananDelete(param *pb.SimpananDeleteRequest) (*pb.DefaultSimpananResponse, error) {
    // Implement your code here
    return &pb.DefaultSimpananResponse{
        // Implement your code here
    }, nil
}


        
func (r *SimpananRepository) SimpananKasBalance(param *pb.SimpananKasBalanceRequest) (*pb.SimpananKasBalanceResponse, error) {
    // Implement your code here
    return &pb.SimpananKasBalanceResponse{
        // Implement your code here
    }, nil
}


        
func (r *SimpananRepository) SimpananUserBalance(param *pb.SimpananUserBalanceRequest) (*pb.SimpananUserBalanceResponse, error) {
    // Implement your code here
    return &pb.SimpananUserBalanceResponse{
        // Implement your code here
    }, nil
}


        