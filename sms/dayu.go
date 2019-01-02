package dayu

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"hash"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

func SendSms(accessKeyId, accessSecret, signName, phoneNumbers, templateParam, templateCode, format, outId string) (bodyBytes []byte, err error) {
	var (
		resp       *http.Response
		requestUrl string
	)
	requestUrl = createRequestUrl(accessKeyId, accessSecret, signName, phoneNumbers, templateParam, templateCode, format, outId)
	resp, err = http.Get(requestUrl)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bodyBytes, err = ioutil.ReadAll(resp.Body)
	return
}

func createRequestUrl(accessKeyId, accessSecret, signName, phoneNumbers, templateParam, templateCode, format, outId string) (reqUrl string) {
	var (
		paras             = map[string]string{}
		keys              = make([]string, 0)
		sortedQueryString string
		stringToSign      string
		signature         string
	)
	rand.Seed(time.Now().UnixNano())
	paras["SignatureMethod"] = "HMAC-SHA1"
	paras["SignatureNonce"] = fmt.Sprintf("%d", rand.Int63())
	paras["AccessKeyId"] = accessKeyId
	paras["SignatureVersion"] = "1.0"
	paras["Timestamp"] = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	paras["Format"] = format

	paras["Action"] = "SendSms"
	paras["Version"] = "2017-05-25"
	paras["RegionId"] = "cn-hangzhou"
	paras["PhoneNumbers"] = phoneNumbers
	paras["SignName"] = signName
	paras["TemplateParam"] = templateParam
	paras["TemplateCode"] = templateCode
	paras["OutId"] = outId

	for key, _ := range paras {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		sortedQueryString += "&" + specialUrlEncode(key) + "=" + specialUrlEncode(paras[key])
	}
	stringToSign += "GET&" + specialUrlEncode("/") + "&" + specialUrlEncode(sortedQueryString[1:])
	signature = sign(accessSecret+"&", stringToSign)
	reqUrl = "http://dysmsapi.aliyuncs.com/?Signature=" + specialUrlEncode(signature) + sortedQueryString
	return
}

func specialUrlEncode(value string) (result string) {
	result = url.QueryEscape(value)
	result = strings.Replace(result, "+", "%20", -1)
	result = strings.Replace(result, "*", "%2A", -1)
	result = strings.Replace(result, "%7E", "~", -1)
	return
}

func sign(accessSecret, stringToSign string) (signature string) {
	var (
		key []byte
		mac hash.Hash
	)
	key = []byte(accessSecret)
	mac = hmac.New(sha1.New, key)
	mac.Write([]byte(stringToSign))
	signature = base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return
}
