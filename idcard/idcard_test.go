package idcard

import "testing"

// 正确
func TestValidate(t *testing.T) {
	if !Validate("110101199003078216") {
		t.Error("test fail")
	}
}

// 超过18位
func TestValidate01(t *testing.T) {
	if Validate("1101011990030782161") {
		t.Error("test fail")
	}
}

// 少于18位
func TestValidate02(t *testing.T) {
	if Validate("11010119900307821") {
		t.Error("test fail")
	}
}

// 出生日期
func TestValidate03(t *testing.T) {
	if Validate("110101189903076751") {
		t.Error("test fail")
	}
}

// 出生日期
func TestValidate04(t *testing.T) {
	if Validate("110101204803070312") {
		t.Error("test fail")
	}
}

// 15位转18位
func TestConvert15To18(t *testing.T) {
	id, err := Convert15To18("110101900307821")
	if err != nil || id != "110101199003078216" {
		t.Error("test fail")
	}
}

// 多于15位
func TestConvert15To1801(t *testing.T) {
	_, err := Convert15To18("1101019003078211")
	if err == nil {
		t.Error("test fail")
	}
}

// 少于15位
func TestConvert15To1802(t *testing.T) {
	_, err := Convert15To18("11010190030782")
	if err == nil {
		t.Error("test fail")
	}
}
