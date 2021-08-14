package main

import (
	"fmt"
	alidns "github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"

)

func main() {
	client, err := alidns.NewClientWithAccessKey("cn-shanghai", "<accessKeyId>", "<accessSecret>")

	request := alidns.CreateUpdateDomainRecordRequest()
	request.Scheme = "https"

	response, err := client.UpdateDomainRecord(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}