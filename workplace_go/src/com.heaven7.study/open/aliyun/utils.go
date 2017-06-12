package aliyun

import (
	"fmt"
	"strings"
	"net/url"
	"time"
	"com.heaven7.study/iso6801"
	"encoding/json"
)
func CurrentIso6801Time()string{
	return Iso6801Time(time.Now().UTC());
}
func Iso6801Time(t time.Time)string{
	now := iso6801.NewISO6801Time(t)
	data, err := json.Marshal(now)
	if err != nil {
		fmt.Println(err)
	}
	var now2 iso6801.ISO6801Time
	err = json.Unmarshal(data, &now2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("called Iso6801Time(): time = " + now2.String());
	return now2.String();
}

func FormatTimeStamp(t time.Time)string{
	           //2016-02-25T12:00:00Z
	template := "2006-01-02T15:04:05Z";
	//template := "2016-02-25T12:00:00Z";
	local_utc, err1 := time.LoadLocation("UTC") //输入参数"UTC"，等同于""
	if err1 != nil {
		fmt.Println(err1)
	}
	return t.In(local_utc).Format(template);
}
//concat with &
func Concat(stringArr []string) string{
	resultStr := "";
	for _,val := range stringArr{
		var link string = "";
		if(!strings.EqualFold(resultStr, "")){
			link="&";
		}
		resultStr = resultStr + link + val;
	}
	return resultStr;
}

func UrlEncode(target string)string  {
	target = url.QueryEscape(target);
	//+
	target = strings.Replace(target,"+","20%",-1);
	//*
	target = strings.Replace(target,"*","%2A",-1);
	//%7E
	target = strings.Replace(target,"%7E","~",-1);
	return target;
}

func PercentReplace(str string) string {
	str = strings.Replace(str, "+", "%20", -1)
	str = strings.Replace(str, "*", "%2A", -1)
	str = strings.Replace(str, "%7E", "~", -1)

	return str
}