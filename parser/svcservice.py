def create_svc_service(data: dict):
    print("Creating service...")
    service_name: str = data.get("service")
    service_data = data.get("service_data")
    service_name = service_name.replace("Service", "")
    template = """
package service

import (
    "context"
)

type {service_name}Service struct {{
    Client pb.{service_name}ServiceClient
    Ctx    context.Context
}}

func Init{service_name}Service(client pb.{service_name}ServiceClient, ctx context.Context) *{service_name}Service {{
    return &{service_name}Service{{
        Client: client,
        Ctx:    ctx,
    }}
}}

""".format(service_name=service_name)
    print(template)
    
    for item in service_data:
        name = item.get("name")
        request = item.get("request")
        response = item.get("response")
        
        func = """
func (s *{service_name}Service) {name}(ctx context.Context, in *pb.{request_name}) (*pb.{response_name}, error) {{
    resp, err := s.Client.{name}(s.Ctx, in)
    if err != nil {{
        return nil, err
    }}
    return resp, nil
}}

        """.format(service_name=service_name, name=name, request_name=request, response_name=response)
        template = template + func
        
    with open(f"result/svc/{service_name}.service.go", "w") as f:
        f.write(template)