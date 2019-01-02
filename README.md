# tools
golang common tools

[![Build Status](https://www.travis-ci.org/wangsongyan/tools.svg?branch=master)](http://travis-ci.org/wangsongyan/tools)

## install
```
go get -u github.com/wangsongyan/tools/...
```
### date
日期相关函数

### dayu
阿里大于短信发送
```golang
package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/wangsongyan/tools/dayu"
)

func main() {
	var (
		accessKeyId   = "testId"
		accessSecret  = "testSecret"
		signName      = "阿里云短信测试专用"
		phoneNumbers  = "15300000001"
		templateParam = "{\"customer\":\"test\"}"
		templateCode  = "SMS_71390007"
		format        = "JSON"
		outId         = 123
		err           error
		bodyBytes     []byte
	)
	bodyBytes, err = dayu.SendSms(accessKeyId, accessSecret, signName, phoneNumbers, templateParam, templateCode, format, outId)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(bodyBytes))
}
```

### idcard
身份证号验证及转换

### mail
smtp邮件发送

### templates
字符串模板渲染
