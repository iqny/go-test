package main

import (
	amqp_open20191212 "github.com/alibabacloud-go/amqp-open-20191212/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	"github.com/alibabacloud-go/tea/tea"
	"os"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *amqp_open20191212.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("amqp-open.cn-qingdao.aliyuncs.com")
	_result = &amqp_open20191212.Client{}
	_result, _err = amqp_open20191212.NewClient(config)
	return _result, _err
}

func _main(args []*string) (_err error) {
	client, _err := CreateClient(tea.String("accessKeyId"), tea.String("accessKeySecret"))
	if _err != nil {
		return _err
	}

	listQueuesRequest := &amqp_open20191212.ListQueuesRequest{}
	// 复制代码运行请自行打印 API 的返回值
	_, _err = client.ListQueues(listQueuesRequest)
	if _err != nil {
		return _err
	}
	return _err
}

func main() {
	err := _main(tea.StringSlice(os.Args[1:]))
	if err != nil {
		panic(err)
	}
}
