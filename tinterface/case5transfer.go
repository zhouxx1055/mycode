package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
)

func main() {

	var interf interface{}
	cookie := make(map[string]string)
	cookie["did"] = "000000000011111"
	cookie["pid"] = "000000000011122"
	cookie["devName"] = "Lamp"

	log.Println("1===", Strval(cookie))
	log.Println("2===", Strval(interf))

	interf = cookie
	log.Println("3===", Strval(interf))
	// log.Println("11===",interf.(string)) //panic: interface conversion: interface {} is map[string]string, not string

	interf1 := interf.(map[string]string)
	log.Println("4===", Strval(interf1))

	var appliance Appliance
	var appliance1 Appliance1
	applianceStr := `{
    "applianceId":"applianceId",
    "attributeName":"attributeName",
    "additionalApplianceDetails":{
      "did":"did112222222222222",
       "pid":"pid112222222222222",
"range":"{\"colortemp\":{\"max\":4000,\"min\":2000}}"
     }
    }`
	err := json.Unmarshal([]byte(applianceStr), &appliance)
	log.Println(err)

	log.Println(appliance)
	log.Println(appliance.AdditionalApplianceDetails)

	err = json.Unmarshal([]byte(applianceStr), &appliance1)
	log.Println(err)
	log.Println(appliance1)
	log.Println(appliance1.AdditionalApplianceDetails)

	cookiep, ok := appliance.AdditionalApplianceDetails.(CookieStu)
	log.Println(ok)
	log.Println(cookiep)
	if ok {
		if cookiep.Range != "" {
			var rangedata struct {
				Colortemp struct {
					Min float64 `json:"min"`
					Max float64 `json:"max"`
				} `json:"colortemp"`
			}
			err := json.Unmarshal([]byte(cookiep.Range), &rangedata)
			if err != nil {
				log.Println(err.Error())
			} else {
				log.Println(fmt.Sprintf("[%v,%v]", rangedata.Colortemp.Min, rangedata.Colortemp.Max))
			}
		}
	}

	cookiep12, ok := appliance.AdditionalApplianceDetails.(map[string]interface{})
	log.Println(ok)
	log.Println(cookiep12)
	//ok=true
	if ok {
		range1, ok := cookiep12["range"]
		if ok {
			if range1 != nil {
				var rangedata struct {
					Colortemp struct {
						Min float64 `json:"min"`
						Max float64 `json:"max"`
					} `json:"colortemp"`
				}
				err := json.Unmarshal([]byte(range1.(string)), &rangedata)
				if err != nil {
					log.Println(err.Error())
				} else {
					log.Println(fmt.Sprintf("==============[%v,%v]", rangedata.Colortemp.Min, rangedata.Colortemp.Max))
				}
			}
		}

	}

	cookiep = appliance1.AdditionalApplianceDetails
	log.Println(ok)
	ok = true
	log.Println(cookiep)
	if ok {
		if cookiep.Range != "" {
			var rangedata struct {
				Colortemp struct {
					Min float64 `json:"min"`
					Max float64 `json:"max"`
				} `json:"colortemp"`
			}
			err := json.Unmarshal([]byte(cookiep.Range), &rangedata)
			if err != nil {
				log.Println(err.Error())
			} else {
				log.Println(fmt.Sprintf("[%v,%v]", rangedata.Colortemp.Min, rangedata.Colortemp.Max))
			}
		}
	}
}

type Appliance struct {
	AdditionalApplianceDetails interface{} `json:"additionalApplianceDetails,omitempty"`
	ApplianceID                string      `json:"applianceId"`
	AttributeName              string      `json:"attributeName,omitempty"`
}

type Appliance1 struct {
	AdditionalApplianceDetails CookieStu `json:"additionalApplianceDetails,omitempty"`
	ApplianceID                string    `json:"applianceId"`
	AttributeName              string    `json:"attributeName,omitempty"`
}
type CookieStu struct {
	Did           string `json:"did,omitempty"`
	Pid           string `json:"pid,omitempty"`
	DevPid        string `json:"devpid,omitempty"`
	DevDid        string `json:"devdid,omitempty"`
	Sdid          string `json:"sdid,omitempty"`
	Spid          string `json:"spid,omitempty"`
	UserId        string `json:"userid,omitempty"`
	DevName       string `json:"devname,omitempty"`
	ModuleId      string `json:"moduleid,omitempty"`
	ModuleType    string `json:"moduletype,omitempty"`
	TokenInfo     string `json:"tokeninfo,omitempty"`
	Mac           string `json:"mac,omitempty"`
	SubDeviceInfo string `json:"subDeviceInfo,omitempty"`
	DeviceInfo    string `json:"deviceInfo,omitempty"`
	Range         string `json:"range,omitempty"` //"range":"{"colortemp":{"max":4000,"min":2000}}"
}

// Strval 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func Strval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}
	log.Println(reflect.TypeOf(value))
	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	case map[string]string:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}
