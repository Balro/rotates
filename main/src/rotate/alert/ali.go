package alert

import (
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"log"
	"rotate/conf"
)

type AliAlert struct {
	Name          string
	AliClient     *sdk.Client
	AliRequest    *requests.CommonRequest
	ShowNumber    int
	CalledNumbers []string
	TtsCode       string
	Params        map[string]string
}

func NewAliAlert(config conf.Ali) *AliAlert {
	aa := AliAlert{}
	aa.Name = config.Name
	client, err := sdk.NewClientWithAccessKey("default", config.Key, config.Secret)
	if err != nil {
		panic(err)
	}
	aa.AliClient = client

	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https"
	request.Domain = "dyvmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SingleCallByTts"
	request.QueryParams["RegionId"] = "default"
	request.QueryParams["CalledShowNumber"] = config.ShowNumber
	request.QueryParams["TtsCode"] = config.TtsCode
	aa.AliRequest = request

	aa.CalledNumbers = config.CalledNumbers
	aa.Params = config.Params
	return &aa
}

func (aa *AliAlert) Send(info string) {
	log.Printf("AliAlert to %+v, %+v\n", aa.CalledNumbers, info)
	for _, n := range aa.CalledNumbers {
		aa.AliRequest.QueryParams["CalledNumber"] = n

		ps := make(map[string]string)
		for k, v := range aa.Params {
			ps[k] = v
		}

		js, err := json.Marshal(ps)
		if err != nil {
			log.Println(err)
		}

		aa.AliRequest.QueryParams["TtsParam"] = string(js)

		response, err := aa.AliClient.ProcessCommonRequest(aa.AliRequest)
		if err != nil {
			panic(err)
		}
		log.Println(response.GetHttpContentString())
	}
}
