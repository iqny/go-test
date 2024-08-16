package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
)

func main() {
	defer func() {
		if errRecover := recover(); errRecover != nil {
			fmt.Println(errRecover)
			fmt.Println(fileInfo(4))
		}
	}()
	var j = `{"orderSellerPrice":"1.00","orderType":"22","venderRemark":"","flag":"","orderSign":"00000000200000000005000002001000030000100000000000600000000001030000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002100000000000203020000000000000000000000000000000000000010000000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","couponDetailList":[[{"abc":"iiii"}]],"orderId":"249786309732","sendpayMap":"{\"34\":\"3\",\"26\":\"2\",\"39\":\"1\",\"29\":\"1\",\"190\":\"2\",\"173\":\"2\",\"174\":\"1\",\"186\":\"2\",\"342\":\"1\",\"188\":\"3\",\"239\":\"2\",\"9\":\"2\",\"229\":\"1\",\"51\":\"6\",\"62\":\"1\",\"20\":\"5\",\"64\":\"3\"}","venderId":"10622317","paymentConfirmTime":"2022-07-16 16:38:32","orderStateRemark":"\u7b49\u5f85\u51fa\u5e93","orderState":"WAIT_SELLER_STOCK_OUT","itemInfoList":[{"skuName":"\uff08\u6d4b\u8bd5\u52ff\u62cd\uff0c\u4e0d\u53d1\u8d27\uff09chuchu\u557e\u557e\u65e5\u672c\u539f\u88c5\u8fdb\u53e3 \u5b9d\u5b9d\u6c90\u6d74\u6413\u6fa1\u624b\u5957 \u65b0\u751f\u5a74\u513f\u513f\u7528\u54c1\u67d4\u8f6f\u62a4\u624b \u767d\u8272","itemTotal":"1","wareId":"10020665903516","outerSkuId":"testjd4206151136003","jdPrice":"1.00","id":228886533499,"invoiceContentId":"","itemExt":"{\"skuUuid\":\"1012_FKq1n1105171626116390912\",\"isVMI\":0,\"newStoreId\":\"0\",\"cid\":\"13288\"}","giftPoint":"0","productNo":"testjd","skuId":"10054887745652","newStoreId":"0"}],"payType":"4-\u5728\u7ebf\u652f\u4ed8","pin":"jd_ituyy715jh7pssw","parentOrderId":"245752632275","modified":"2022-07-16 16:38:32","commission":{"amount":0,"centFactor":100,"cent":0,"currency":"CNY","currencyCode":"CNY"},"freightPrice":"0.00","vatInfo":[],"balanceUsed":"0.00","directParentOrderId":"245752632275","invoiceInfo":"\u4e0d\u9700\u8981\u5f00\u5177\u53d1\u7968","orderSource":"\u79fb\u52a8\u7aef\u8ba2\u5355","dataVersion":"1657960712000","orderTotalPrice":"1.00","deliveryType":"\u4efb\u610f\u65f6\u95f4","storeId":"0","mdbStoreId":"","idSopShipmenttype":67,"orderPayment":"1.00","idShipmenttype":"70","outBoundDate":"0001-01-01 00:00:00","consigneeInfo":{"town":"\u8d64\u5c97\u8857\u9053","city":"\u5e7f\u5dde\u5e02","county":"\u6d77\u73e0\u533a","mobile":"136*******5","telephone":"136*******5","cityId":"1601","townId":"63199","provinceId":"19","province":"\u5e7f\u4e1c","countyId":"3634","fullAddress":"AAS\/ZFJFpmfveSmeZ0P9i+iPFTw2VsdIkCtv0YhpKHQZAdj5ssES7b4bjL2KwIPn\/iK+ZcCvaBRb++Gpxh8mZde6Fk2cgEYiLGWLMbjzeytj42zpmYlhKmszUHsaNjhNBawgQxYpKKgAKIchK6Si5Ji+tCAAmW6E5fwNtA3IaxGcgx2IDi5LC0d9HmgWVZ1asFI=","fullname":"AAS\/ZFJFpmfveSmeZ0P9i+iPiAVtOMmowMRtStdbSOA32r6wKNWTMeyV2CU1umLIhl0=","desen_telephone":"656233eb3c9d6035653f5a6d7f5965ac0c53815f3e9b785c17d35366fbbb72f1b52994","desen_mobile":"656233eb3c9d6035653f5a6d7f5965ac0c53815f3e9b785c17d35366fbbb72f1b52994"},"sellerDiscount":"0.00","returnOrder":"0","invoiceEasyInfo":{"invoiceState":"3","invoiceType":"0","invoiceContentId":"1","invoiceTitle":"AAS\/ZFJFpmfveSmeZ0P9i+iP9iBeF5j0Za+BaS+i3tEKfUdoE6rd9HfWubLzXhJiVbk="},"storeOrder":"","orderStartTime":"2022-07-16 16:37:38"}`
	var content = make(map[string]interface{})
	json.Unmarshal([]byte(j), &content)
	//var promotionDetails = make(map[string]interface{}) //content["promotion_details"].(map[string]interface{})["promotion_detail"])
	/*if promotionDetail, ok := content["promotion_details"]; ok {
		details := promotionDetail.(map[string]interface{})["promotion_detail"].([]interface{})
		for _, detail := range details {
			row := detail.(map[string]interface{})
			fmt.Println(row["promotion_desc"])
		}
	}*/
	/*itemInfoList := content["itemInfoList"].([]interface{})

	for _, goods := range itemInfoList {
		row := goods.(map[string]interface{})
		fmt.Println(row["itemTotal"])
	}*/
	if couponDetailList, ok := content["couponDetailList"]; ok {
		details := couponDetailList.([]interface{})
		for _, coupon := range details {
			for _, row := range coupon.([]interface{}) {
				r := row.(map[string]interface{})
				fmt.Println(r["abc"].(float64))
			}

		}
	}
}
func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}
