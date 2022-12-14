package util

import (
	"encoding/json"
	"fmt"
	"gin/pkg/setting"
	"log"
	"os"
	"reflect"
	"strings"
)

func LanguageMapping(language string) string {

	language = strings.ToLower(language)
	var res string = "en"

	if strings.HasPrefix(language, "pt") {
		res = "pt"
	} else if strings.HasPrefix(language, "fr") {
		res = "fr"
	} else if strings.HasPrefix(language, "de") {
		res = "de"
	} else if strings.HasPrefix(language, "ru") {
		res = "ru"
	} else if strings.HasPrefix(language, "zh-cn") {
		res = "cn"
	} else if strings.HasPrefix(language, "zh-tw") {
		res = "tw"
	} else if strings.HasPrefix(language, "ru") {
		res = "ru"
	} else if strings.HasPrefix(language, "ja-jp") {
		res = "jp"
	} else if strings.HasPrefix(language, "ko") {
		res = "kr"
	} else if strings.HasPrefix(language, "pt") {
		res = "pt"
	} else if strings.HasPrefix(language, "vi") {
		res = "vi"
	}

	return res
}

func GetNotifiArn(app string, version int) string {
	var arn string
	switch app {
	case "poto":
		arn = "arn:aws:sns:us-east-1:561020269087:app/APNS/poto"
		break
	case "manly":
		arn = "arn:aws:sns:us-east-1:561020269087:app/APNS/manly"
		break
	case "facey":
		if version == 90909 {
			arn = "arn:aws:sns:us-east-1:561020269087:app/APNS_SANDBOX/faceytest"
		} else {
			arn = "arn:aws:sns:us-east-1:561020269087:app/APNS/facey2"
		}
		break
	case "Everlook":
		if version == 90909 {
			arn = "arn:aws:sns:us-east-1:561020269087:app/APNS_SANDBOX/everlookDev"
		} else {
			arn = "arn:aws:sns:us-east-1:561020269087:app/APNS/evacam"
		}
		break
	case "evalook2":
		if version == 90909 {
			arn = "arn:aws:sns:us-east-1:561020269087:app/APNS_SANDBOX/everlook2_dev"
		} else {
			arn = "arn:aws:sns:us-east-1:561020269087:app/APNS/everlook2_dis"
		}
		break
	case "Meepo2":
		if version == 90909 {
			arn = "arn:aws:sns:us-east-1:561020269087:app/APNS_SANDBOX/Meepo_test"
		} else {
			arn = "arn:aws:sns:us-east-1:561020269087:app/APNS/Meepo"
		}
		break
	case "famey2":
		if version == 90909 {
			arn = "aarn:aws:sns:us-east-1:561020269087:app/APNS_SANDBOX/famey2dev"
		} else {
			arn = "arn:aws:sns:us-east-1:561020269087:app/APNS/famey2prod"
		}
		break
	}

	return arn
}

func IsChina(country string) bool {
	var res bool = false

	country = strings.ToLower(country)

	if country == "cn" {
		res = true
	}

	return res
}

func IsAsia(country string) bool {

	var res bool
	asia := [...]string{"cn", "CN", "jp", "JP", "KR", "kr", "my", "MY", "sg", "SG", "tw", "TW", "HK", "hk", "mo", "MO"}

	if country != "" {
		for i := 0; i < len(asia); i++ {
			if asia[i] == country {
				res = true
				break
			}
		}
	}

	return res
}

// ????????????
func Empty(params interface{}) bool {
	//???????????????
	var (
		flag          bool = true
		default_value reflect.Value
	)

	r := reflect.ValueOf(params)

	//???????????????????????????
	default_value = reflect.Zero(r.Type())
	//??????params ???????????? ??????default_value???????????????????????????????????? ?????????????????????????????? ???????????????false
	if !reflect.DeepEqual(r.Interface(), default_value.Interface()) {
		flag = false
	}

	return flag
}

func InArrayNotNil(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return false
			}
		}
	case reflect.Map:
		v := targetValue.MapIndex(reflect.ValueOf(obj))
		if v.IsValid() && !Empty(v) {
			return false
		}
	}
	return true
}

func GetUrl(url string, isChina bool) (resUrl string) {

	if url == "" {
		return ""
	}

	if strings.HasPrefix(url, "http") {
		resUrl = url
	} else {

		if isChina == true {
			resUrl = "https://s3.cn-north-1.amazonaws.com.cn/poto-php/" + url
		} else {
			resUrl = "http://d273s9in8abcdi.cloudfront.net/" + url
		}
	}

	return resUrl
}

//json????????????map
func JsonStrToMap(param string) interface{} {

	if strings.HasPrefix(param, "[") && strings.HasPrefix(param, "{") {
		var res []map[string]interface{}
		json.Unmarshal([]byte(fmt.Sprintf("%+v", param)), &res)
		return res

	} else if strings.HasPrefix(param, "[") && strings.HasPrefix(param, "{") == false {
		var res []interface{}
		json.Unmarshal([]byte(fmt.Sprintf("%+v", param)), &res)
		return res
	} else {
		var res map[string]interface{}
		json.Unmarshal([]byte(param), &res)
		return res
	}

}

//map???json?????????
func MapToJsonStr(param interface{}) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}

func InArray(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}

	return false
}

//struct???map
func StructToMapViaReflect(person interface{}) map[string]interface{} {
	res := make(map[string]interface{}, 0)

	t := reflect.TypeOf(person)
	v := reflect.ValueOf(person)

	if t.Kind() != reflect.Struct {
		panic("????????? struct")
	}

	for i := 0; i < t.NumField(); i++ {
		res[string(t.Field(i).Tag.Get("json"))] = v.Field(i).Interface()
	}

	return res
}

//??????map????????? ??????????????????????????????
func GetMapValueForKey(arr map[string]interface{}, key string, default_value interface{}) interface{} {
	res := default_value
	if _, value := arr[key]; value {
		res = value
	}
	return res
}

//????????????????????????
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

//????????????????????????
func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func GetEnv(key string) (r string) {
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		r = os.Getenv(key)
	} else {
		sec, err := setting.Cfg.GetSection("app")
		if err != nil {
			log.Fatal(2, "Fail to get section 'database': %v", err)
		}

		r = sec.Key(key).String()
	}
	return r
}
