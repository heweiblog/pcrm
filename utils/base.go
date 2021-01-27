package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"pcrm/models"
	"reflect"
	"time"
)

/*
type Content struct {
	Mid     uint   `json:"id" binding:"required"`
	Bt      string `json:"bt" binding:"required"`
	Sbt     string `json:"sbt" binding:"required"`
	Source  string `json:"source" binding:"required"`
	Service string `json:"service" binding:"required"`
	Op      string `json:"op" binding:"required"`
	Reason  string `json:"reason"`
	Data    interface{}
	Jdata   datatypes.JSON `json:"data" binding:"required"`
}
*/

type BindRes struct {
	Rcode       uint   `json:"rcode"`
	Description string `json:"description"`
}

type BaseMap map[string]func(c *models.Content) string

const (
	DnsService = "dns"
	ZrpService = "zrp"
	MsSource   = "ms"
)

var (
	OpMap         map[string]bool
	CheckMethods  BaseMap
	HandleMethods BaseMap
)

func init() {
	OpMap = make(map[string]bool)
	OpMap["add"] = true
	OpMap["delete"] = true
	OpMap["update"] = true
	OpMap["clear"] = true
	OpMap["query"] = true
	CheckMethods = make(BaseMap)
	HandleMethods = make(BaseMap)
	//CheckMethods["dns"+"view"+"rules"] = PviewCheck
	CheckMethods["dns"+"view"+"qpslimit"] = PviewLimitCheck
	//HandleMethods["dns"+"view"+"qpslimit"] = ViewLimitHandle
	CheckMethods["zrp"+"zrpaccesscontrol"+"switch"] = SwitchCheck
	CheckMethods["zrp"+"zrpaccesscontrol"+"rules"] = ZidCheck
	CheckMethods["zrp"+"backend"+"forwardserver"] = BackendForwardCheck
}

//发送json数据 可满足所有请求方法
func ClientDo(url, method string, data interface{}) string {
	var req *http.Request
	if data != nil {
		requestBody := new(bytes.Buffer)
		json.NewEncoder(requestBody).Encode(data)
		req, _ = http.NewRequest(method, url, requestBody)
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, _ = http.NewRequest(method, url, nil)
	}

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var res BindRes
	err = json.Unmarshal(body, &res)
	if err != nil {
		return err.Error()
	}
	return res.Description
}

func HandleData(c []models.Content) {
	for i := 0; i < len(c); i++ {
		if c[i].Source != MsSource {
			c[i].Reason = "source value error"
		}
		if c[i].Mid < 0 {
			c[i].Reason = "id value error"
		}
		if _, ok := OpMap[c[i].Op]; ok == false {
			c[i].Reason = "op value error"
		}
		key := c[i].Service + c[i].Bt + c[i].Sbt
		if f, ok := CheckMethods[key]; ok {
			//校验模块
			res := f(&c[i])
			if res == "" {
				//校验通过 入队(通道) 直接启动协程
				//go testPost(d)
				fmt.Println("数据校验通过", c)
				fmt.Println(reflect.TypeOf(c[i].Data))
			} else if res == "done" {
				fmt.Println("数据无需配置", res)
			} else {
				//校验失败，直接记录oplog，或者发到通道，在另一个协程中收取写入
				fmt.Println("数据校验失败", res)
			}
		} else {
			//service bt sbt或对应的函数有问题，直接记录oplog失败，或者发到通道，在另一个协程中收取写入
			fmt.Println("unsupported config")
			c[i].Reason = "unsupported config"
		}
	}
}
