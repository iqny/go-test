package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strconv"
	"time"
)

type User struct {
	Name string
	Age  int
}

type Employee struct {
	Name string
	Age  int
	Role string
}
type BaseChannelOrder struct {
	ChannelId    int64                `json:"channel_id"`
	TotalFee     float64              `json:"total_fee"`
	PayId        int64                `json:"pay_id"`
	PayNo        string               `json:"pay_no"`
	Tid          string               `json:"tid"`
	BuyerNick    string               `json:"buyer_nick"`
	SellerNick   string               `json:"seller_nick"`
	Status       string               `json:"status"`
	Origin       string               `json:"origin"`
	PayTime      string               `json:"pay_time"`
	Modified     string               `json:"modified"`
	Created      string               `json:"created"`
	Goods        []*ChannelOrderGoods `json:"goods"`
	ReceiverInfo *ChannelReceiverInfo `json:"receiver_info"`
}
type ChannelOrderGoods struct {
	DiscountFee float64 `json:"discount_fee"`
	Num         int64   `json:"num"`
	NumIid      int64   `json:"num_iid"`
	Oid         string  `json:"oid"`
	OrderFrom   string  `json:"order_from"`
	SkuSn       string  `json:"sku_sn"`
	Payment     float64 `json:"payment"`
	Price       float64 `json:"price"`
	SkuId       string  `json:"sku_id"`
	Status      string  `json:"status"`
	TotalFee    float64 `json:"total_fee"`
}
type ChannelReceiverInfo struct {
	Province string `json:"province"`
	Address  string `json:"address"`
	City     string `json:"city"`
	District string `json:"district"`
	Mobile   string `json:"mobile"`
	Name     string `json:"name"`
	Zip      string `json:"zip"`
}

func main() {
	/*user := User{Name: "dj", Age: 18}
	employee := Employee{}

	copier.Copy(&employee, &user)
	fmt.Printf("%#v\n", employee)
	fmt.Printf("%%%s%%\n", "abc")
	m := make(map[string]interface{})
	ms := make(map[string]interface{})
	ms = map[string]interface{}{"ms": "value"}
	m = map[string]interface{}{"teset": ms}
	mapt(m)*/
	ss := `{"channel_id":1,"total_fee":29800,"pay_id":2088802545036797,"pay_no":"2016102521001001795007279709","tid":"163288460521882636","buyer_nick":"相蝗","seller_nick":"test-s","goods":[{"discount_fee":0,"num":100,"num_iid":0,"oid":"163288460521879179","order_from":"WAP,WAP","sku_sn":"4973210993669","payment":14900,"price":149,"sku_id":"4973210993669","status":"WAIT_SELLER_SEND_GOODS","total_fee":14900},{"discount_fee":0,"num":100,"num_iid":0,"oid":"163288460521986755","order_from":"WAP,WAP","sku_sn":"4973210993676","payment":14900,"price":149,"sku_id":"4973210993676","status":"WAIT_SELLER_SEND_GOODS","total_fee":14900}],"receiver_info":{"province":"广东省","address":"敦和路189号海珠科技园2期1号楼","city":"广州市","district":"海珠区","mobile":"15102088122","name":"刘莉","zip":"510900"},"promotion_detail":[{"discount_fee":0,"id":"126950071993312757","promotion_desc":"聚划算:省0.00元","promotion_id":"Tmall-3153345006_33729171907","promotion_name":"聚划算"}],"status":"WAIT_SELLER_SEND_GOODS","origin":"","pay_time":"","modified":"2022-06-08 15:23:25","created":"2021-09-29 11:03:25"}`
	//jm := make(map[string]interface{})
	jm := new(BaseChannelOrder)
	json.Unmarshal([]byte(ss), &jm)
	tid, _ := strconv.ParseInt(jm.Tid, 10, 64)
	fmt.Println(tid)
	/*goods := jm["goods"].([]interface{})[0]
	g := goods.(map[string]interface{})
	num, _ := strconv.ParseInt(strconv.FormatFloat(g["num"].(float64), 'f', 0, 64), 10, 64)
	fmt.Printf("%T", num)
	fmt.Println(g["price"].(float64))*/
	st := []*goods{{s: "abx"}}

	name1(st)
	fmt.Println(st[0])
	fmt.Println(GetTimeDifference())

	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum:", cpuNum)
	fmt.Printf("%d\n", 0xEB86D391)
	fmt.Printf("%d\n", 0xFFEFF47D)
}

func GetTimeDifference() int64 {
	nowTime := time.Now()
	nowTimeStamp := nowTime.Unix()
	nowTimeStr := nowTime.Format("2006-01-02")
	t2, _ := time.ParseInLocation("2006-01-02", nowTimeStr, time.Local)
	towTimeStamp := t2.AddDate(0, 0, 1).Unix()
	fmt.Println(towTimeStamp)
	fmt.Println(nowTimeStamp)
	return towTimeStamp - nowTimeStamp
}

type info struct {
	ss []*goods `json:"ss"`
}
type goods struct {
	s string
}

func mapt(m map[string]interface{}) {
	switch m["teset"].(type) {
	case map[string]interface{}:
		ms := m["teset"].(map[string]interface{})
		//fmt.Println(ms["ms"])
		m["teset"] = ms
	}
	fmt.Println(m["teset"])
}

func name1(gg []*goods) {
	for _, g := range gg {
		g.s = "oooooo"
	}
}
