package aliyun

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

//阿里云push高级接口(方法首字母必须大写，否则ide可能无法识别)
func PushAdvanced(linkStr string)  {

	baseUrl := "http://cloudpush.aliyuncs.com/?";
	realUrl := baseUrl + linkStr;
	fmt.Println(">>> START >>> request realUrl = " + realUrl)
	resp , err := http.Get(realUrl);
	if(err != nil){
		fmt.Println("push failed( GET ) >>> url + " + realUrl);
	}
	handle_result_pushAdvanced(resp);
	fmt.Println(">>> END >>> request realUrl = " + realUrl)
}

func handle_result_pushAdvanced(resp *http.Response)  {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("called [ handle_result_pushAdvanced() ]>>> Error: ", err)
	}
	fmt.Println(string(body))
}



