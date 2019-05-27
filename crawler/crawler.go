package crawler

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/kataras/iris/context"
)

const (
	CONTENT_TYPE_JSON = "application/json"
	CONTENT_TYPE_FORM = "application/x-www-form-urlencoded"
	//CONTENT_TYPE_OCTET = "application/octet-stream"
	//DEFAULT_USER_AGENT = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36"
)

type H map[string]string

type App struct {
	*http.Client
}

func NewApp() *App {
	jar, _ := cookiejar.New(nil)
	return &App{
		Client: &http.Client{
			Jar: jar,
		},
	}
}

func (app *App) DoGet(requrl string, headers H) (data []byte, err error) {
	var (
		req  *http.Request
		resp *http.Response
	)
	req, err = http.NewRequest(http.MethodGet, requrl, nil)
	if err != nil {
		return
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err = app.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, err = ioutil.ReadAll(resp.Body)
	return
}

func (app *App) DoPostJson(requrl string, headers H, body interface{}) (data []byte, err error) {
	data, err = json.Marshal(body)
	if err != nil {
		return
	}
	if nil == headers {
		headers = make(H)
	}
	headers[context.ContentTypeHeaderKey] = CONTENT_TYPE_JSON
	data, err = app.DoPost(requrl, headers, bytes.NewReader(data))
	return
}

func (app *App) DoPostForm(requrl string, headers, params H) (data []byte, err error) {
	var (
		u = url.Values{}
	)
	for k, v := range params {
		u.Set(k, v)
	}
	if nil == headers {
		headers = make(H)
	}
	headers[context.ContentTypeHeaderKey] = CONTENT_TYPE_FORM
	data, err = app.DoPost(requrl, headers, strings.NewReader(u.Encode()))
	return
}

func (app *App) DoPost(requrl string, headers H, body io.Reader) (data []byte, err error) {
	var (
		req  *http.Request
		resp *http.Response
	)
	req, err = http.NewRequest(http.MethodPost, requrl, body)
	if err != nil {
		return
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err = app.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, err = ioutil.ReadAll(resp.Body)
	return
}
