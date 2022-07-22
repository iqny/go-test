package parser_address

import (
	"regexp"
	"strings"
)

func ParserAddress(addrStr string) (address map[string]string) {
	if addrStr == "" {
		return nil
	}
	provinceReg := `(.*?(省|自治区|北京市|北京|天津市|天津|重庆市|重庆|上海市|上海|特别行政区))`
	compile := regexp.MustCompile(provinceReg)
	province := compile.FindString(addrStr)
	if province != "" {
		addrStr = addrStr[len(province):]
		if strings.Index(province, "市") == -1 &&
			strings.Index(province, "区") == -1 &&
			strings.Index(province, "省") == -1 {
			province = province + "市"
		}
	}
	var cityReg string
	if province == "香港特别行政区" {
		cityReg = `(.*?(香港岛|九龙|新界))`
	} else {
		cityReg = `(.*?(市|自治州|地区|区划|县))`
	}
	compile=regexp.MustCompile(cityReg)
	city := compile.FindString(addrStr)
	if city!="" {
		addrStr = addrStr[len(city):]
	}else{
		city = province
	}

	compile = regexp.MustCompile(`(.+?(区|县|市))`)
	area :=compile.FindString(addrStr)
	if area !=""{
		addrStr = addrStr[len(area):]
	}
	return map[string]string{"province":province,"city":city,"area":area,"address":addrStr}
}
