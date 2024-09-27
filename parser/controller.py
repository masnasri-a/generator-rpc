def create_controller(data: dict):
    print("Creating controller...")
    service_name: str = data.get("service")
    service_data = data.get("service_data")
    names: str = service_name.replace("Service", "")
    print(f"Creating controller for '{service_name}'...")
    template = """
package controllers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type {names}Controller struct {{
	Client pb.{names}ServiceClient
	Ctx    context.Context
}}


func Init{names}Controller(client pb.{names}ServiceClient, ctx context.Context) *{names}Controller {{
	return &{names}Controller{{
		Client: client,
		Ctx:    ctx,
	}}
}}

""".format(names=names)
    print(template)

    for item in service_data:
        name = item.get("name")
        request = item.get("request")
        response = item.get("response")
        
        open_api = """
        
        
// PingExample godoc
// @Schemes
// @Tags {names}
// @Accept json
// @Produce json

// @Success 200 {{object}} models.BaseResponseModel
"""
        
        if 'Create' in name:
            open_api += "// @Router /"+names.lower()+"/create [post]"
        elif 'Update' in name:
            open_api += "// @Router /"+names.lower()+"/update [put]"
        elif 'Delete' in name:
            open_api += "// @Router /"+names.lower()+"/delete [delete]"
        else:
            open_api += "// @Router /"+names.lower()+"/** [get]"
        
        func = """
func (s *SavingController) {name}Controller(r *gin.Context) {{
    baseResponse := models.BaseResponseModel{{}}
    svc := services.Init{names}Service(s.Client, s.Ctx)
    param := pb.{request}{{}}
    res, err := svc.{name}(u.Ctx, &param)
    if err != nil {{
		baseResponse.Metadata = models.MetadataModel{{
			Status:  "error",
			Code:    500,
			Message: "error",
		}}
		baseResponse.Data = nil
		r.AbortWithStatusJSON(500, baseResponse)
		return
	}}
	baseResponse.Metadata = models.MetadataModel{{
		Status:  "OK",
		Code:    200,
		Message: "Success",
	}}
	baseResponse.Data = res

	r.JSON(http.StatusOK, baseResponse)
}}""".format(names=names, name=name, request=request, response=response)
        template = template + open_api + func
        
    with open(f"result/svc/{names}.controller.go", "w") as f:
        f.write(template)