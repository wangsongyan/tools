package templates

import (
	"testing"
)

func TestExecute(t *testing.T) {

	var templateContent = `hello {{.}}`

	output, err := Execute(templateContent, "world")
	if err != nil || output != "hello world" {
		t.Errorf("render failed")
	}

}

func TestExecute_01(t *testing.T) {

	var templateContent = `my name is {{.Name}},I'm {{.Age}} years old.`

	data := struct {
		Name string
		Age  int
	}{
		Name: "bob",
		Age:  18,
	}

	output, err := Execute(templateContent, data)
	if err != nil || output != "my name is bob,I'm 18 years old." {
		t.Errorf("render failed")
	}

}
