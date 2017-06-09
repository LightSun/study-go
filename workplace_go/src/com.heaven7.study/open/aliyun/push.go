package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"net/url"
	"strings"
	"math/rand"
	"time"
	"strconv"
)

var (
	Key_Value map[string]string = map[string]string{
		"Format":"JSON",
		"RegionId":"cn-hangzhou",
		"Version":"2016-08-01",
		"AccessKeyId":"",
		"SignatureMethod":"HMAC-SHA1",
		"Timestamp":"",
		"SignatureVersion":"1.0",
		"SignatureNonce":"",
          //推送目标(destination) //TODO populate value
		"Action":"Push",
		"AppKey":"",
		"Target":"ALL",
		"TargetValue":"ALL",
		"DeviceType":"ALL",
		//推送配置 //TODO
		"PushType":"NOTICE",
		"Title":"title_go_1",
		"Body":"body_go_1",
		//保留参数: AndroidMusic, AndroidNotifyType, AndroidActivity(full name of activity), AndroidOpenUrl,
		// AndroidNotificationBarType,AndroidNotificationBarPriority,AndroidExtParameters, PushTime
		"AndroidOpenType":"APPLICATION", //or ACTIVITY, URL , NONE
		"AndroidExtParameters":"{}", //json map "{\"k1\":\"ios\",\"k2\":\"v2\"}"
		"StoreOffline":"true",

	}
)
//高级接口： https://help.aliyun.com/document_detail/48089.html?spm=5176.doc48085.2.1.fDxlAn
//签名机制: https://help.aliyun.com/document_detail/48047.html?spm=5176.doc48045.2.2.cb9cX0
func main() {
	//need replace when upload to github
	Key_Value["AccessKeyId"]=""
	Key_Value["AppKey"]=""

	//cur := time.Now()
	Key_Value["Timestamp"]=strconv.FormatInt(time.Now().UnixNano(),10)
	Key_Value["SignatureNonce"]=strconv.FormatInt(rand.Int63n(64664465),10);

	tempMap := make(map[string]string, 20);
	for k := range tempMap {
		fmt.Println("before encode:  key = %s , value = %s", k, Key_Value[k])
		key := urlEncode(k);
		value := urlEncode(Key_Value[k]);
		fmt.Println("after encode:  key = %s , value = %s", key, value)
		tempMap[key]=value
	}

}

func urlEncode(target string)string  {
	var r string = "";
	//+
	r = strings.Replace(url.QueryEscape(target),"+","20%",-1);
	//*
	r = strings.Replace(url.QueryEscape(r),"*","%2A",-1);
	//%7E
	r = strings.Replace(url.QueryEscape(r),"%7E","~",-1);
	return r;
}

/**
1, rq 排序 . exclude  Signature
2, key-value utf-8 的 url编码
     A-Z、a-z、0-9以及字符“-”、“_”、“.”、“~”不编码；
     对于其他字符编码成“%XY”的格式，其中XY是字符对应ASCII码的16进制表示。比如英文的双引号（”）对应的编码就是%22
     对于扩展的UTF-8字符，编码成“%XY%ZA…”的格式
     需要说明的是英文空格（ ）要被编码是%20，而不是加号（+）。
 */