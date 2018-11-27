package idcard

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	YYYYMMDD = "20060102"
)

var (
	verifyCodeMap = [...]string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}
	wi            = [...]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	startTime, _  = time.Parse(YYYYMMDD, "19000101")

	LENGTH_17_ERR = errors.New("the length of string must be >= 17")
	LENGTH_15_ERR = errors.New("the length of string must equal to 15")
	BIRTHDAY_ERR  = errors.New("error birthday")
)

//func main() {
//	fmt.Println(Validate("110101189903076751"))
//}

// 校验18位身份证号是否合法
func Validate(id string) bool {
	var (
		code string
		err  error
	)
	//id = strings.Join(strings.Fields(id), "")
	id = strings.ToUpper(id)
	if len(id) != 18 {
		return false
	}
	if !validateDate(id[6:14]) {
		return false
	}
	code, err = sumVerifyCode(id[:17])
	if err != nil {
		return false
	}
	return code == id[17:]
}

// 15位身份证号转18位
func Convert15To18(id string) (idstr string, err error) {
	var (
		code string
	)
	//id = strings.Join(strings.Fields(id), "")
	id = strings.ToUpper(id)
	if len(id) != 15 {
		err = LENGTH_15_ERR
		return
	}
	if !validateDate("19" + id[6:12]) {
		err = BIRTHDAY_ERR
		return
	}
	id = id[:6] + "19" + id[6:]
	code, err = sumVerifyCode(id)
	if err != nil {
		return
	}
	return id + code, nil
}

func validateDate(birth string) bool {
	var (
		birthdate time.Time
		err       error
	)
	birthdate, err = time.Parse(YYYYMMDD, birth)
	if err != nil {
		return false
	}
	fmt.Println(birthdate)
	return birthdate.After(startTime) && birthdate.Before(time.Now())
}

func sumVerifyCode(id string) (code string, err error) {
	if len(id) < 17 {
		err = LENGTH_17_ERR
		return
	}
	arr := make([]int, 17)
	for i := 0; i < 17; i++ {
		arr[i], err = strconv.Atoi(string(id[i]))
		if err != nil {
			return
		}
	}
	var res int
	for i := 0; i < 17; i++ {
		res += arr[i] * wi[i]
	}
	return verifyCodeMap[(res % 11)], nil
}
