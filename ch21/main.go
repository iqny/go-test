package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var j = `{"trade_fullinfo_get_response":{"trade":{"adjust_fee":"0.00","alipay_id":2088802545036797,"alipay_no":"2016102521001001795007279709","alipay_point":0,"available_confirm_fee":29800,"buyer_alipay_no":"13932626505","buyer_area":"广东省广州市敦和路189号海珠科技园2期1号楼","buyer_cod_fee":"0.00","buyer_email":"","buyer_nick":"刘莉","buyer_obtain_point_fee":0,"buyer_rate":false,"cod_fee":"0.00","cod_status":"NEW_CREATED","commission_fee":"0.00","created":"2021-09-29 11:03:25","discount_fee":"0.00","has_post_fee":true,"is_3D":false,"is_brand_sale":false,"is_daixiao":false,"is_force_wlb":false,"is_lgtype":false,"is_part_consign":false,"is_wt":false,"modified":"2021-09-29 11:03:25","num":200,"num_iid":544160959115,"orders":{"order":[{"adjust_fee":"0.00","buyer_rate":false,"cid":578,"discount_fee":"0","is_oversold":false,"num":100,"num_iid":67964013,"oid":"163288460521879179","order_from":"WAP,WAP","outer_iid":"","outer_sku_id":"4973210993669","part_mjz_discount":"0.00","payment":14900,"pic_path":"https:\/\/img.alicdn.com\/imgextra\/i3\/3676232520\/O1CN01dI9Rmn1UUCYmJVRJM_!!0-item_pic.jpg_430x430q90.jpg","price":"149.00","refund_status":"NO_REFUND","seller_rate":false,"seller_type":"B","sku_id":"4973210993669","sku_properties_name":"","snapshot_url":"l=>163288460521882636_1","status":"WAIT_SELLER_SEND_GOODS","title":"","total_fee":14900},{"adjust_fee":"0.00","buyer_rate":false,"cid":82773138,"discount_fee":"0","is_oversold":false,"num":100,"num_iid":27,"oid":"163288460521986755","order_from":"WAP,WAP","outer_iid":"","outer_sku_id":"4973210993676","part_mjz_discount":"0.00","payment":14900,"pic_path":"https:\/\/img.alicdn.com\/imgextra\/i3\/3676232520\/O1CN01dI9Rmn1UUCYmJVRJM_!!0-item_pic.jpg_430x430q90.jpg","price":"149.00","refund_status":"NO_REFUND","seller_rate":false,"seller_type":"B","sku_id":"4973210993676","sku_properties_name":"","snapshot_url":"l=>163288460521882636_1","status":"WAIT_SELLER_SEND_GOODS","title":"","total_fee":14900}]},"pay_time":"2021-09-29 11:03:25","payment":29800,"pcc_af":0,"pic_path":"","point_fee":0,"post_fee":"0.00","promotion_details":{"promotion_detail":[{"discount_fee":"0.00","id":"126950071993312757","promotion_desc":"聚划算:省0.00元","promotion_id":"Tmall-3153345006_33729171907","promotion_name":"聚划算"}]},"price":29800,"real_point_fee":0,"received_payment":"0.00","receiver_address":"敦和路189号海珠科技园2期1号楼","receiver_city":"广州市","receiver_district":"海珠区","receiver_mobile":"15102088122","receiver_name":"刘莉","receiver_state":"广东省","receiver_zip":"510900","seller_alipay_no":"fgu@ecco.com","seller_can_rate":false,"seller_cod_fee":"0.00","seller_email":"e-commerce-cn@ecco.com","seller_flag":0,"seller_memo":"","seller_mobile":"15921534824","seller_name":"skxqm_yn","seller_phone":"0571-28181301","seller_rate":false,"service_tags":{"logistics_tag":[{"logistic_service_tag_list":{"logistic_service_tag":[{"service_tag":"lgType=-4","service_type":"FAST"}]},"order_id":"163288460521882636"}]},"shipping_type":"express","snapshot_url":"l:163288460521882636_1","status":"WAIT_SELLER_SEND_GOODS","tid":"163288460521882636","title":"skxqm_yn","total_fee":29800,"trade_from":"WAP,WAP","type":"step","credit_card_fee":29800}}}`
	var content = make(map[string]interface{})
	json.Unmarshal([]byte(j), &content)
	//var promotionDetails = make(map[string]interface{}) //content["promotion_details"].(map[string]interface{})["promotion_detail"])
	if promotionDetail, ok := content["promotion_details"]; ok {
		details := promotionDetail.(map[string]interface{})["promotion_detail"].([]interface{})
		for _, detail := range details {
			row := detail.(map[string]interface{})
			fmt.Println(row["promotion_desc"])
		}
	}
	tradeFullinfoGetResponse := content["trade_fullinfo_get_response"].(map[string]interface{})
	trade := tradeFullinfoGetResponse["trade"].(map[string]interface{})
	if orders, ok := trade["orders"]; ok {
		order := orders.(map[string]interface{})["order"].([]interface{})
		for _, goods := range order {
			row := goods.(map[string]interface{})
			fmt.Println(row["adjust_fee"])
		}
	}
	MyPrintln("abc")
	fmt.Println(07)
}

func MyPrintln[T any](a T) {
	fmt.Println(a)
}
