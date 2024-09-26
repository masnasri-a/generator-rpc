package parser

import (
	"fmt"
	"strings"
)

func GetService(data string) string {
	var serviceName string
	for _, line := range strings.Split(data, "\n") {
		if strings.Contains(line, "service ") {
			serviceName = strings.Split(line, " ")[1]
		}
		if strings.Contains(line, "rpc ") {
			rpcParser(line)
		}
	}
	fmt.Println("service name : ", serviceName)
	return serviceName
}

func rpcParser(data string) {
	data = strings.Replace(data, "{}", "", -1)
	data = strings.Replace(data, ";", "", -1)
	data = strings.Replace(data, "rpc", "", -1)
	data = strings.Replace(data, "returns", "", -1)
	data = strings.TrimSpace(data)
	splitter := strings.Split(data, ")  (")
	prefix := strings.Split(splitter[0], "(")[0]
	requestParams := strings.Split(splitter[0], "(")[1]
	requestParams = strings.Replace(requestParams, ")", "", -1)
	responseParam := strings.Split(splitter[1], ")")[0]
	fmt.Println("rpc prefix : ", prefix)
	fmt.Println("rpc request params : ", requestParams)
	fmt.Println("rpc response params : ", responseParam)
	fmt.Println("========================================")

}
