package templates

import (
	"bytes"
	"html/template"
)

func Execute(tplContent string, data interface{}) (string, error) {
	t := template.New("")
	tpl, err := t.Parse(tplContent)
	if err != nil {
		return "", err
	}
	buff := new(bytes.Buffer)
	err = tpl.Execute(buff, data)
	return buff.String(), err
}
