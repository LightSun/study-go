package main

import (
	"strconv"
	"time"
	"fmt"
	"sort"
	"math/rand"

	"crypto/hmac"
	"crypto/sha1"

	//"io"
	"encoding/base64"
	"com.heaven7.study/open/aliyun"
)

var (
	Key_Value map[string]string = map[string]string{
		"Format":"JSON",
		"RegionId":"cn-hangzhou",
		"Version":"2016-08-01",
		"AccessKeyId":"",
		"SignatureMethod":"HMAC-SHA1",
		"Timestamp":"", //
		"SignatureVersion":"1.0",
		"SignatureNonce":"", //
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
		//Signature 签名的

	}
)
//高级接口： https://help.aliyun.com/document_detail/48089.html?spm=5176.doc48085.2.1.fDxlAn
//签名机制: https://help.aliyun.com/document_detail/48047.html?spm=5176.doc48045.2.2.cb9cX0
func main() {
	//need replace when upload to github
	accessKeySecret := "fUY2be6g7V7zJVWJ4L1ZlTQnVE4Zea";
	Key_Value["AccessKeyId"]="LTAI2ujuwJjEucva"
	Key_Value["AppKey"]="24011705"

	Key_Value["Timestamp"]= aliyun.FormatTimeStamp(time.Now().UTC())
	//Key_Value["Timestamp"]= aliyun.CurrentIso6801Time()
	Key_Value["SignatureNonce"]=strconv.FormatInt(rand.Int63n(time.Now().UnixNano()),10);
	fmt.Println("time stamp : = " + Key_Value["Timestamp"])
	fmt.Println("SignatureNonce : = " + Key_Value["SignatureNonce"])

	step1_str,stringArr := step1();
	stringToSign := step2(step1_str)
	//计算hmac-sha1 and Signature
	signValue := step3And4(accessKeySecret, stringToSign);

	finalStr := aliyun.Concat(strep5(signValue, stringArr));
	aliyun.PushAdvanced(finalStr);

}

//添加 Signature and value //TODO 签名算法有问题
func strep5(signValue string, stringArr []string) []string {
	//size := len(stringArr);
	arr2 := append(stringArr, "Signature=" + aliyun.UrlEncode(signValue))
	sort.Strings(arr2);
	fmt.Printf("step 5 >>> after sort: %v\n" , arr2)
	return arr2;
}
//计算hmac-sha1 and Signature
func step3And4(accessKeySecret string,stringToSign string)string{
	sign := hmac.New(sha1.New, []byte(accessKeySecret))
	//hmac_sha1 := sign.Sum([]byte(stringToSign));
	//fmt.Printf("sign_hmac-sha1 value = %v\n" , hmac_sha1)
	//io.WriteString(sign, stringToSign)
	sign.Write([]byte(stringToSign))
	sum := sign.Sum(nil);
	result := fmt.Sprintf("%x",sum);
	fmt.Println("HMAC-SHA1 = " + result)

	signature := base64.StdEncoding.EncodeToString(sum)
	return signature;
}
//string to sign
func step2(str string)string{
	stringToSign:= "GET" + "&" + "%2F"+"&" + str
	fmt.Println("stringToSign: " + stringToSign)
	return stringToSign;
}
//sort and encode
func step1() (string,[]string) {
	//step: 1 key自然排序,
	// 2 key-value编码，
	// 3 对编码后的参数名称和值使用英文等号（=）进行连接。
	// 4,再把英文等号连接得到的字符串按参数名称的字典顺序依次使用&符号连接，即得到规范化请求字符串。

	//1
	stringArr := make([]string, len(Key_Value));
	index := 0;
	for k,_ := range Key_Value {
		stringArr[index] = k;
		index ++;
	}
	fmt.Printf("step 1.1 >>> before sort: %v\n" ,stringArr)
	sort.Strings(stringArr);
	fmt.Printf("step 1.1 >>>  after sort: %v\n" ,stringArr)

	//2
	tempMap := make(map[string]string, 20);
	for _,val := range stringArr{
		//fmt.Printf("step 1.2 >>> before encode:  key = %s , value = %s\n", val, Key_Value[val])
		encodeKey := aliyun.UrlEncode(val);
		encodeValue := aliyun.UrlEncode(Key_Value[val]);
		tempMap[encodeKey] = encodeValue;
		//fmt.Printf("step 1.2 >>> after encode:  key = %s , value = %s\n", encodeKey, encodeValue)
	}

	//3
	index = 0;
	for k,v := range tempMap {
		stringArr[index] = k + "=" + v;
		fmt.Printf("step 1.3 >>> string concat (=): raw key = %s, value = %s ---> result = %s\n", k, v, stringArr[index]);
		index ++;
	}
	//4
	fmt.Printf("step 1.4 >>> before sort: %v\n" ,stringArr)
	sort.Strings(stringArr);
	fmt.Printf("step 1.4 >>> after sort: %v\n" , stringArr)

	resultStr := aliyun.Concat(stringArr);
	fmt.Println("step 1.4 >>> after concat(&): "+ resultStr)
	return resultStr,stringArr;
}

/**
1, rq 排序 . exclude  Signature
2, key-value utf-8 的 url编码
     A-Z、a-z、0-9以及字符“-”、“_”、“.”、“~”不编码；
     对于其他字符编码成“%XY”的格式，其中XY是字符对应ASCII码的16进制表示。比如英文的双引号（”）对应的编码就是%22
     对于扩展的UTF-8字符，编码成“%XY%ZA…”的格式
     需要说明的是英文空格（ ）要被编码是%20，而不是加号（+）。
 */