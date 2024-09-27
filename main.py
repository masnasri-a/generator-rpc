import os
from parser import repo, service, svcservice, controller


def check_source_file():
    print("Checking source file...")
    source_file = os.path.join("source", "test.proto")
    if os.path.isfile(source_file):
        print(f"File '{source_file}' exists.")
        return source_file
    else:
        print(f"File '{source_file}' does not exist.")
        return None

def load_source_file(source_file):
    print("Loading source file...")
    with open(source_file, "r") as f:
        lines = f.readlines()
        return get_service_data(lines)
    print(f"Loaded {len(lines)} lines.")
    return lines

def get_service_data(lines : list[str]):
    print("Getting service data...")
    service_data = []
    service = None
    for line in lines:
        if line.startswith("service"):
            service = line.split(" ")[1].strip()
        if line.strip().startswith("rpc"):
            rpc = line.strip()
            spliter = rpc.split("returns")
            first = spliter[0].strip().replace("rpc ", "").strip()
            content = first.split("(") 
            nameContent = content[0]
            requestContent = content[1].replace(")", "").strip()
            responseContent = spliter[1].strip().replace(";", "").replace("{}","").replace("(", "").replace(")", "").strip()
            temp = {
                "name": nameContent,
                "request": requestContent,
                "response": responseContent
            }
            service_data.append(temp)
    return {
        "service": service,
        "service_data": service_data
    }

if __name__ == "__main__":
    source = check_source_file()
    if not source:
        print("Source file does not exist.")
        exit(1)
    lines = load_source_file(source)
    # repo_name = repo.create_repo(lines)
    # service = service.create_service(repo_name, lines)
    # svcservice.create_svc_service(lines)
    controller.create_controller(lines)