package dingtalk

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Robot struct {
	accessToken string
}

type RespMsg struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (r *Robot) SendMsg(msg interface{}) (respMsg RespMsg, err error) {
	var (
		byteData []byte
	)
	byteData, err = json.Marshal(msg)
	if err != nil {
		return
	}
	respMsg, err = r.postServer(byteData)
	return
}

func NewRobot(accessToken string) (r *Robot) {
	return &Robot{
		accessToken: accessToken,
	}
}

func (r *Robot) postServer(data []byte) (respMsg RespMsg, err error) {
	var (
		resp     *http.Response
		respData []byte
	)
	resp, err = http.Post("https://oapi.dingtalk.com/robot/send?access_token="+r.accessToken, "application/json", bytes.NewReader(data))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	respData, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(respData, &respMsg)
	return
}
