def create_service(repo_name : str ,data: dict):
    print("Creating service...")
    service_name: str = data.get("service")
    service_data = data.get("service_data")
    service_name = service_name.replace("Service", "")
    template = """
package service

import (
    "context"
    "fmt"
 
    "gorm.io/gorm"
)

type {service_name}Service struct {{
    pb.UnimplementedLoanServiceServer
	Db *gorm.DB
}}
""".format(service_name=service_name)

    for item in service_data:
        name = item.get("name")
        request = item.get("request")
        response = item.get("response")
        
        func = """
func (s *{service_name}Service) {name}(ctx context.Context, param *pb.{request_name}) (*pb.{response_name}, error) {{
    repo := repository.Init{service_name}Repository(s.Db)
    response, err := repo.{name}(param)
    if err != nil {{
        return nil, err
    }}
    return resp, nil
}}
        """.format(service_name=service_name, name=name, request_name=request, response_name=response)
        template = template + func
    
    with open(f"result/rpc/{service_name}.service.go", "w") as f:
        f.write(template)
    print(f"Service '{service_name}' created.")
