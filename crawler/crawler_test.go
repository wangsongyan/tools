package crawler

import "testing"

func TestApp_DoGet(t *testing.T) {
	app := NewApp()
	_, err := app.DoGet("http://example.com/", nil)
	if err != nil {
		t.Errorf("DoGet failed: %v", err)
	}
}

func TestApp_DoPostForm(t *testing.T) {
	app := NewApp()
	_, err := app.DoPostForm("http://example.com/", nil, H{
		"foo": "bar",
	})
	if err != nil {
		t.Errorf("DoPostForm failed: %v", err)
	}
}

func TestApp_DoPostJson(t *testing.T) {
	app := NewApp()
	_, err := app.DoPostJson("http://example.com/", nil, H{
		"foo": "bar",
	})
	if err != nil {
		t.Errorf("DoPostJson failed: %v", err)
	}
}
