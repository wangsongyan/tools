package date

import (
	"testing"
	"time"
)

func TestGetMondayOfWeek(t *testing.T) {
	t1, _ := time.Parse("2006-01-02", "2018-02-24")
	t2, _ := time.Parse("2006-01-02", "2018-02-19")
	if !t2.Equal(GetMondayOfWeek(t1)) {
		t.Errorf("not same date")
	}
}

func TestGetMondayOfWeek_01(t *testing.T) {
	t1, _ := time.Parse("2006-01-02", "2018-02-25")
	t2, _ := time.Parse("2006-01-02", "2018-02-19")
	if !t2.Equal(GetMondayOfWeek(t1)) {
		t.Errorf("not same date")
	}
}

func TestGetFridayOfWeek(t *testing.T) {
	t1, _ := time.Parse("2006-01-02", "2018-02-24")
	t2, _ := time.Parse("2006-01-02", "2018-02-23")
	if !t2.Equal(GetFridayOfWeek(t1)) {
		t.Errorf("not same date")
	}
}

func TestGetFridayOfWeek_01(t *testing.T) {
	t1, _ := time.Parse("2006-01-02", "2018-02-25")
	t2, _ := time.Parse("2006-01-02", "2018-02-23")
	if !t2.Equal(GetFridayOfWeek(t1)) {
		t.Errorf("not same date")
	}
}
