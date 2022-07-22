package parser_address

import (
	"testing"
)

func TestParserAddress(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		province string
		city     string
		area     string
		address  string
	}{
		{"1", "上海浦东新区王桥路786号普洛斯金桥南园区1号厂房", "上海市", "上海市", "浦东新区", "王桥路786号普洛斯金桥南园区1号厂房"},
		{"2", "香港特别行政区香港岛元朗区小黄村", "香港特别行政区", "香港岛", "元朗区", "小黄村"},

		  {"3","广东省广州市天河区北京路88号", "广东省", "广州市", "天河区", "北京路88号"},
		  {"4","广东省深圳市龙华新区大浪街道同胜科技大厦", "广东省", "深圳市", "龙华新区", "大浪街道同胜科技大厦"},
		  {"5","深圳市龙华新区大浪街道同胜科技大厦", "", "深圳市", "龙华新区", "大浪街道同胜科技大厦"},
		  {"6","广东省湛江市霞山区人民大道南81号新宇大厦15楼1501房", "广东省", "湛江市", "霞山区", "人民大道南81号新宇大厦15楼1501房"},
		  {"7","广东省湛江市霞山区人民大道南81号新宇大厦15楼1501房--12", "广东省", "湛江市", "霞山区", "人民大道南81号新宇大厦15楼1501房--12"},
		  {"8","山东省青岛市市南区北京路88号", "山东省", "青岛市", "市南区", "北京路88号"},
		  {"9","山东省青岛市区南区北京路88号", "山东省", "青岛市", "区南区", "北京路88号"},
		  {"10","山东省青岛市县南区北京路88号", "山东省", "青岛市", "县南区", "北京路88号"},
		  {"11","山东省青岛市县南县北京县路88号", "山东省", "青岛市", "县南县", "北京县路88号"},
		  {"12","山东省青岛市县南县经济开发区88号", "山东省", "青岛市", "县南县", "经济开发区88号"},
		{"13","广东省广州市宝安区", "广东省", "广州市", "宝安区", ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response := ParserAddress(test.input)
			if province := response["province"]; province != test.province {
				t.Errorf("ParserAddress(%q)[province]%v = %v", test.input,province, test.province)
			}
			if city := response["city"]; city != test.city {
				t.Errorf("ParserAddress(%q)[city]%v  = %v", test.input,city, test.city)
			}
			if area := response["area"]; area != test.area {
				t.Errorf("ParserAddress(%q)[area]%v  = %v", test.input,area, test.area)
			}
			if address := response["address"]; address != test.address {
				t.Errorf("ParserAddress(%q)[address]%v = %v", test.input,address, test.address)
			}
		})
	}
	//ParserAddress("上海浦东新区王桥路786号上海普洛斯金桥南园区1号厂房")
	//ParserAddress("香港特别行政区香港岛元朗区小黄村")
	//ParserAddress("山东省青岛市市南区北京路88号")
}
