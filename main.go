package main

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"io"
	"io/ioutil"
	"net/http"
	"time"
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

func changeDNS() {
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
	fmt.Printf("response is %#v\n", response)
}

func main() {
	ticker := time.NewTicker(1 * time.Minute)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				fmt.Println(getIP())
			}
		}
	}()

	time.Sleep(876000 * time.Hour)
	ticker.Stop()
	done <- true
	fmt.Println("----------到点熄灯----------")
}