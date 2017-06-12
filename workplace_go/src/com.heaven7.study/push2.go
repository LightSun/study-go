package main

import (
	"github.com/denverdino/aliyungo/push"
)

func main() {
	//TODO
	var AccessKeyId string=""
	var AccessSecretKey string=""
	var AppKey int64 =0

	client := push.NewClient(AccessKeyId, AccessSecretKey)
	client.SetDebug(true)

	args := new(push.PushArgs);
	args.AppKey = AppKey
	args.Target = "ALL"
	args.TargetValue = "ALL"
	//args.DeviceType = "ALL"  //all 要求android 和ios必须都配置才行。控制台。
	args.DeviceType = "ANDROID"
	args.PushType = "NOTICE"
	args.Title = "title_1_test"
	args.Body = "desc_1_test"
	args.AndroidOpenType = "APPLICATION"
	args.AndroidExtParameters = "{}"
	args.StoreOffline = "true"
	client.Push(args)
}
