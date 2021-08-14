package main

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"io"
	"io/ioutil"
	"net/http"
)

func getIP() string {
	resp, err := http.Get("http://whatismyip.akamai.com/")
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
	request.Value = getIP()

	response, err := client.UpdateDomainRecord(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println("DNS解析已修改为:", getIP())
	fmt.Printf("response is %#v\n", response)
}