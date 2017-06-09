package main

import (
	fm "fmt"
	"net/url"
	rt "runtime"

	"crypto/md5"
	"flag"
	"time"
	"strings"
)

func main() {
	//encode1()
	encode2()

	keyvalue := make(map[string]string)
	keyvalue["Format"] = "XML"
	keyvalue["AccessKeyId"] = "testid"
	keyvalue["Action"] = "GetDeviceInfos"
	keyvalue["SignatureMethod"] = "HMAC-SHA1"
	keyvalue["RegionId"] = "cn-hangzhou"
	fm.Println(url.Parse("ABCDEFGHIJKLMNOPQURSTUVWXYZ-_.~"))
     //加号和=号
	query :=url.Values{};
	query.Add(" =", "\"__,__'__:__ ") // '='='%3D'   ' '='+'
	fm.Println(query.Encode())
	fm.Println(url.QueryEscape("\"__,__'__:__ ")) //编码任意字符串, ok
	//string util
	fm.Println(strings.Replace("hhhh+", "+", "%20", -1));

	result1, err1 := url.Parse("\"__,__'__:__ ");
	if(result1 != nil){
		fm.Println(result1.Query().Encode(), err1) //只能针对完整的url编码
	}
}
func encode2()  {
	fm.Println("./timetoken -t 3600 -key key  -url url")

	var t int64
	var key string
	var resUrl string
	flag.Int64Var(&t, "t", 0, "expire timestamp")
	flag.StringVar(&key, "key", "", "encrypt key")
	flag.StringVar(&resUrl, "url", "", "resource url")
	flag.Parse()

	if t == 0 || key == "" || resUrl == "" {
		return
	}

	expireTime := fm.Sprintf("%x", time.Now().Unix()+t)

	resUri, pErr := url.Parse(resUrl)
	if pErr != nil {
		return
	}
	fm.Println(resUri)
	path := resUri.EscapedPath()
	fm.Println(path)
	rawStr := fm.Sprintf("%s%s%s", key, path, expireTime)
	fm.Println(rawStr)
	md5H := md5.New()
	md5H.Write([]byte(rawStr))

	sign := fm.Sprintf("%x", md5H.Sum(nil))
	//sign := hex.EncodeToString(md5H.Sum(nil))

	var newUrl string
	if strings.Contains(resUrl, "?") {
		newUrl = fm.Sprintf("%s&sign=%s&t=%s", resUrl, sign, expireTime)
	} else {
		newUrl = fm.Sprintf("%s?sign=%s&t=%s", resUrl, sign, expireTime)
	}

	fm.Println(newUrl)
}

func encode1()  {
	var urlStr string = "http://baidu.com/index.php/?abc=1_羽毛"
	l, err := url.ParseQuery(urlStr)
	fm.Println(l, err)
	l2, err2 := url.ParseRequestURI(urlStr)
	fm.Println(l2, err2)
	l3, err3 := url.Parse(urlStr)
	fm.Println(l3, err3)

	fm.Println(l3.Path)
	fm.Println(l3.RawQuery)
	fm.Println(l3.Query())
	fm.Println(l3.Query().Encode())

	fm.Println(l3.RequestURI())
	fm.Printf("Hello World! version : %s", rt.Version())
	fm.Println(url.QueryEscape(urlStr))
}
