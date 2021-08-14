package main

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"io"
	"io/ioutil"
	"net/http"
)

func getIP() string {
	resp, err := http.Get("http://ifconfig.me/ip")
	if err != nil {
		return ""
	}

	content, _ := ioutil.ReadAll(resp.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)
	return string(content)
}

func main() {
	client, err := alidns.NewClientWithAccessKey("cn-shanghai", "<accessKeyId>", "<accessSecret>")

	request := alidns.CreateUpdateDomainRecordRequest()
	request.Scheme = "https"

	request.RecordId = "<RecordId>"
	request.RR = "<RR>"
	request.Type = "<Type>"
	request.Value = "<Value>"

	response, err := client.UpdateDomainRecord(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}