
def create_repo(data: dict):
    print("Creating repository...")
    service_name: str = data.get("service")
    service_data = data.get("service_data")
    repo_name = service_name.replace("Service", "Repository")
    template = """
package repository

import (
    "gorm.io/gorm"
)

    """
    struct = """
type {repo_name} struct {{
    db *gorm.DB,
}}

func Init{repo_name}(db *gorm.DB) *{repo_name} {{
    return &{repo_name}{{
        DB: db
    }}
}}
    """.format(repo_name=repo_name)
    template = template + struct
    for item in service_data:
        name = item.get("name")
        request = item.get("request")
        response = item.get("response")
        
        func = """
func (r *{repo_name}) {name}(param *pb.{request_name}) (*pb.{response_name}, error) {{
    // Implement your code here
    return &pb.{response_name}{{
        // Implement your code here
    }}, nil
}}


        """.format(repo_name=repo_name, name=name, request_name=request, response_name=response)
        template = template + func
    
        # create func
        
    with open(f"result/rpc/{name}.repository.go", "w") as f:
        f.write(template)
    # print(struct)
    
    # print(create_func)
    
    return repo_name