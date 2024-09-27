package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SimpananController struct {
	Client pb.SimpananServiceClient
	Ctx    context.Context
}

func InitSimpananController(client pb.SimpananServiceClient, ctx context.Context) *SimpananController {
	return &SimpananController{
		Client: client,
		Ctx:    ctx,
	}
}

// PingExample godoc
// @Schemes
// @Tags {names}
// @Accept json
// @Produce json

// @Success 200 {{object}} models.BaseResponseModel
// @Router /simpanan/create [post]
func (s *SavingController) SimpananCreateController(r *gin.Context) {
	baseResponse := models.BaseResponseModel{}
	svc := services.InitSimpananService(s.Client, s.Ctx)
	param := pb.SimpananCreateRequest{}
	res, err := svc.SimpananCreate(u.Ctx, &param)
	if err != nil {
		baseResponse.Metadata = models.MetadataModel{
			Status:  "error",
			Code:    500,
			Message: "error",
		}
		baseResponse.Data = nil
		r.AbortWithStatusJSON(500, baseResponse)
		return
	}
	baseResponse.Metadata = models.MetadataModel{
		Status:  "OK",
		Code:    200,
		Message: "Success",
	}
	baseResponse.Data = res

	r.JSON(http.StatusOK, baseResponse)
}

// PingExample godoc
// @Schemes
// @Tags {names}
// @Accept json
// @Produce json

// @Success 200 {{object}} models.BaseResponseModel
// @Router /simpanan/** [get]
func (s *SavingController) SimpananDetailController(r *gin.Context) {
	baseResponse := models.BaseResponseModel{}
	svc := services.InitSimpananService(s.Client, s.Ctx)
	param := pb.SimpananDetailRequest{}
	res, err := svc.SimpananDetail(u.Ctx, &param)
	if err != nil {
		baseResponse.Metadata = models.MetadataModel{
			Status:  "error",
			Code:    500,
			Message: "error",
		}
		baseResponse.Data = nil
		r.AbortWithStatusJSON(500, baseResponse)
		return
	}
	baseResponse.Metadata = models.MetadataModel{
		Status:  "OK",
		Code:    200,
		Message: "Success",
	}
	baseResponse.Data = res

	r.JSON(http.StatusOK, baseResponse)
}

// PingExample godoc
// @Schemes
// @Tags {names}
// @Accept json
// @Produce json

// @Success 200 {{object}} models.BaseResponseModel
// @Router /simpanan/** [get]
func (s *SavingController) SimpananListController(r *gin.Context) {
	baseResponse := models.BaseResponseModel{}
	svc := services.InitSimpananService(s.Client, s.Ctx)
	param := pb.SimpananListFilterRequest{}
	res, err := svc.SimpananList(u.Ctx, &param)
	if err != nil {
		baseResponse.Metadata = models.MetadataModel{
			Status:  "error",
			Code:    500,
			Message: "error",
		}
		baseResponse.Data = nil
		r.AbortWithStatusJSON(500, baseResponse)
		return
	}
	baseResponse.Metadata = models.MetadataModel{
		Status:  "OK",
		Code:    200,
		Message: "Success",
	}
	baseResponse.Data = res

	r.JSON(http.StatusOK, baseResponse)
}

// PingExample godoc
// @Schemes
// @Tags {names}
// @Accept json
// @Produce json

// @Success 200 {{object}} models.BaseResponseModel
// @Router /simpanan/update [put]
func (s *SavingController) SimpananUpdateController(r *gin.Context) {
	baseResponse := models.BaseResponseModel{}
	svc := services.InitSimpananService(s.Client, s.Ctx)
	param := pb.SimpananUpdateRequest{}
	res, err := svc.SimpananUpdate(u.Ctx, &param)
	if err != nil {
		baseResponse.Metadata = models.MetadataModel{
			Status:  "error",
			Code:    500,
			Message: "error",
		}
		baseResponse.Data = nil
		r.AbortWithStatusJSON(500, baseResponse)
		return
	}
	baseResponse.Metadata = models.MetadataModel{
		Status:  "OK",
		Code:    200,
		Message: "Success",
	}
	baseResponse.Data = res

	r.JSON(http.StatusOK, baseResponse)
}

// PingExample godoc
// @Schemes
// @Tags {names}
// @Accept json
// @Produce json

// @Success 200 {{object}} models.BaseResponseModel
// @Router /simpanan/delete [delete]
func (s *SavingController) SimpananDeleteController(r *gin.Context) {
	baseResponse := models.BaseResponseModel{}
	svc := services.InitSimpananService(s.Client, s.Ctx)
	param := pb.SimpananDeleteRequest{}
	res, err := svc.SimpananDelete(u.Ctx, &param)
	if err != nil {
		baseResponse.Metadata = models.MetadataModel{
			Status:  "error",
			Code:    500,
			Message: "error",
		}
		baseResponse.Data = nil
		r.AbortWithStatusJSON(500, baseResponse)
		return
	}
	baseResponse.Metadata = models.MetadataModel{
		Status:  "OK",
		Code:    200,
		Message: "Success",
	}
	baseResponse.Data = res

	r.JSON(http.StatusOK, baseResponse)
}

// PingExample godoc
// @Schemes
// @Tags {names}
// @Accept json
// @Produce json

// @Success 200 {{object}} models.BaseResponseModel
// @Router /simpanan/** [get]
func (s *SavingController) SimpananKasBalanceController(r *gin.Context) {
	baseResponse := models.BaseResponseModel{}
	svc := services.InitSimpananService(s.Client, s.Ctx)
	param := pb.SimpananKasBalanceRequest{}
	res, err := svc.SimpananKasBalance(u.Ctx, &param)
	if err != nil {
		baseResponse.Metadata = models.MetadataModel{
			Status:  "error",
			Code:    500,
			Message: "error",
		}
		baseResponse.Data = nil
		r.AbortWithStatusJSON(500, baseResponse)
		return
	}
	baseResponse.Metadata = models.MetadataModel{
		Status:  "OK",
		Code:    200,
		Message: "Success",
	}
	baseResponse.Data = res

	r.JSON(http.StatusOK, baseResponse)
}

// PingExample godoc
// @Schemes
// @Tags {names}
// @Accept json
// @Produce json

// @Success 200 {{object}} models.BaseResponseModel
// @Router /simpanan/** [get]
func (s *SavingController) SimpananUserBalanceController(r *gin.Context) {
	baseResponse := models.BaseResponseModel{}
	svc := services.InitSimpananService(s.Client, s.Ctx)
	param := pb.SimpananUserBalanceRequest{}
	res, err := svc.SimpananUserBalance(u.Ctx, &param)
	if err != nil {
		baseResponse.Metadata = models.MetadataModel{
			Status:  "error",
			Code:    500,
			Message: "error",
		}
		baseResponse.Data = nil
		r.AbortWithStatusJSON(500, baseResponse)
		return
	}
	baseResponse.Metadata = models.MetadataModel{
		Status:  "OK",
		Code:    200,
		Message: "Success",
	}
	baseResponse.Data = res

	r.JSON(http.StatusOK, baseResponse)
}
