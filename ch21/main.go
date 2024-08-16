package main

import (
	"encoding/json"
	"fmt"
	"time"
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
	/*diffConvertCnts := []DiffConvertCnt{{Type: 1, ConvertCntGap: -1789, DayBalance: -36, AdId: 1751091337040030},
		{Type: 1, ConvertCntGap: -139, DayBalance: -33, AdId: 1751338258342916},
		{Type: 1, ConvertCntGap: 6, DayBalance: -33, AdId: 1751338285403165},
		{Type: 2, ConvertCntGap: -310, DayBalance: -113, AdId: 1744260831036475},
		{Type: 2, ConvertCntGap: 10, DayBalance: -113, AdId: 1744260536385568},
		{Type: 2, ConvertCntGap: 10, DayBalance: -118, AdId: 1743854516042783},
		{Type: 2, ConvertCntGap: 10, DayBalance: -118, AdId: 1743848655659118},
		{Type: 2, ConvertCntGap: 10, DayBalance: -118, AdId: 1743842346412076},
		{Type: 2, ConvertCntGap: 10, DayBalance: -118, AdId: 1743841866915884},
		{Type: 2, ConvertCntGap: 10, DayBalance: -118, AdId: 1743840378872846},
		{Type: 2, ConvertCntGap: 10, DayBalance: -118, AdId: 1743840234974255},
		{Type: 2, ConvertCntGap: 10, DayBalance: -126, AdId: 1743132095045636},
		{Type: 2, ConvertCntGap: 10, DayBalance: -126, AdId: 1743131414058027},
		{Type: 2, ConvertCntGap: 10, DayBalance: -126, AdId: 1743129862592548},
		{Type: 2, ConvertCntGap: 10, DayBalance: -129, AdId: 1742867107510376},
		{Type: 2, ConvertCntGap: 10, DayBalance: -129, AdId: 1742864049404959},
		{Type: 2, ConvertCntGap: 10, DayBalance: -129, AdId: 1742846425851944},
		{Type: 2, ConvertCntGap: 10, DayBalance: -113, AdId: 1744291786848283},
		{Type: 2, ConvertCntGap: 10, DayBalance: -113, AdId: 1744284303336471},
		{Type: 2, ConvertCntGap: 10, DayBalance: -114, AdId: 1744194212378699},
		{Type: 2, ConvertCntGap: 10, DayBalance: -117, AdId: 1743926354773083},
		{Type: 2, ConvertCntGap: 10, DayBalance: -117, AdId: 1743914311490639},
		{Type: 2, ConvertCntGap: 10, DayBalance: -118, AdId: 1743810123922436},
		{Type: 2, ConvertCntGap: 10, DayBalance: -118, AdId: 1743809987145731},
		{Type: 2, ConvertCntGap: 10, DayBalance: -119, AdId: 1743733514555421},
		{Type: 2, ConvertCntGap: 10, DayBalance: -120, AdId: 1743646228056124},
		{Type: 2, ConvertCntGap: 10, DayBalance: -120, AdId: 1743632338190380},
		{Type: 2, ConvertCntGap: 10, DayBalance: -121, AdId: 1743548973911088},
		{Type: 2, ConvertCntGap: 10, DayBalance: -123, AdId: 1743394299708499},
		{Type: 2, ConvertCntGap: 10, DayBalance: -123, AdId: 1743394197056595},
		{Type: 2, ConvertCntGap: 10, DayBalance: -124, AdId: 1743312571527196},
		{Type: 2, ConvertCntGap: 10, DayBalance: -125, AdId: 1743206455251022},
		{Type: 2, ConvertCntGap: 10, DayBalance: -125, AdId: 1743205095733280},
		{Type: 2, ConvertCntGap: 10, DayBalance: -127, AdId: 1743052597924955},
		{Type: 2, ConvertCntGap: 10, DayBalance: -128, AdId: 1742919533917191},
		{Type: 2, ConvertCntGap: 10, DayBalance: -128, AdId: 1742902982588462},
		{Type: 2, ConvertCntGap: 10, DayBalance: -128, AdId: 1742901456816183},
		{Type: 2, ConvertCntGap: 10, DayBalance: -129, AdId: 1742861067915300},
		{Type: 2, ConvertCntGap: -1352, DayBalance: -129, AdId: 1742861036596292},
		{Type: 2, ConvertCntGap: 10, DayBalance: -129, AdId: 1742843842029647},
		{Type: 2, ConvertCntGap: 10, DayBalance: -129, AdId: 1742843470773268},
		{Type: 2, ConvertCntGap: -696, DayBalance: -129, AdId: 1742842990166051},
		{Type: 2, ConvertCntGap: 10, DayBalance: -129, AdId: 1742842798897176},
		{Type: 2, ConvertCntGap: 10, DayBalance: -129, AdId: 1742842655961092},
		{Type: 2, ConvertCntGap: -182, DayBalance: -109, AdId: 1744672759390243},
		{Type: 2, ConvertCntGap: -352, DayBalance: -108, AdId: 1744761702891556},
		{Type: 2, ConvertCntGap: -248, DayBalance: -108, AdId: 1744762203233291},
		{Type: 2, ConvertCntGap: -60, DayBalance: -107, AdId: 1744836480316416},
		{Type: 2, ConvertCntGap: -132, DayBalance: -107, AdId: 1744836777162803},
		{Type: 2, ConvertCntGap: -375, DayBalance: -107, AdId: 1744852899499060},
		{Type: 2, ConvertCntGap: -272, DayBalance: -107, AdId: 1744852870521859},
		{Type: 2, ConvertCntGap: -1931, DayBalance: -107, AdId: 1744856813690899},
		{Type: 2, ConvertCntGap: 10, DayBalance: -106, AdId: 1744911449405476},
		{Type: 2, ConvertCntGap: -5, DayBalance: -106, AdId: 1744911517231144},
		{Type: 2, ConvertCntGap: 10, DayBalance: -106, AdId: 1744911738559495},
		{Type: 2, ConvertCntGap: 10, DayBalance: -106, AdId: 1744911783917630},
		{Type: 2, ConvertCntGap: -318, DayBalance: -106, AdId: 1744947890070557},
		{Type: 2, ConvertCntGap: -103, DayBalance: -105, AdId: 1744993382829117},
		{Type: 2, ConvertCntGap: -3102, DayBalance: -105, AdId: 1744994436637704},
		{Type: 2, ConvertCntGap: -3864, DayBalance: -105, AdId: 1745038859401291},
		{Type: 2, ConvertCntGap: -350, DayBalance: -104, AdId: 1745077499534371},
		{Type: 2, ConvertCntGap: -390, DayBalance: -104, AdId: 1745129784175647},
		{Type: 2, ConvertCntGap: -1968, DayBalance: -104, AdId: 1745129901055004},
		{Type: 2, ConvertCntGap: -26, DayBalance: -104, AdId: 1745129860868204},
		{Type: 2, ConvertCntGap: -145, DayBalance: -104, AdId: 1745130065933375},
		{Type: 2, ConvertCntGap: -61, DayBalance: -97, AdId: 1745741314668551},
		{Type: 2, ConvertCntGap: -397, DayBalance: -95, AdId: 1745942997157997},
		{Type: 2, ConvertCntGap: -2270, DayBalance: -95, AdId: 1745943126670372},
		{Type: 2, ConvertCntGap: -116, DayBalance: -93, AdId: 1746075479999518},
		{Type: 2, ConvertCntGap: 5, DayBalance: -93, AdId: 1746101386113085},
		{Type: 2, ConvertCntGap: -150, DayBalance: -91, AdId: 1746284709062679},
		{Type: 2, ConvertCntGap: -186, DayBalance: -91, AdId: 1746285086931048},
		{Type: 2, ConvertCntGap: -520, DayBalance: -91, AdId: 1746297099202607},
		{Type: 2, ConvertCntGap: -261, DayBalance: -90, AdId: 1746398638600195},
		{Type: 2, ConvertCntGap: -47, DayBalance: -90, AdId: 1746398625872932},
		{Type: 2, ConvertCntGap: -301, DayBalance: -87, AdId: 1746622081780795},
		{Type: 2, ConvertCntGap: -38, DayBalance: -88, AdId: 1746584637765683},
		{Type: 2, ConvertCntGap: -485, DayBalance: -87, AdId: 1746642784569348},
		{Type: 2, ConvertCntGap: -9, DayBalance: -87, AdId: 1746643475429435},
		{Type: 2, ConvertCntGap: -925, DayBalance: -87, AdId: 1746669573431357},
		{Type: 2, ConvertCntGap: -424, DayBalance: -85, AdId: 1746823377899539},
		{Type: 2, ConvertCntGap: -212, DayBalance: -85, AdId: 1746823405146171},
		{Type: 2, ConvertCntGap: -2401, DayBalance: -84, AdId: 1746889575478334},
		{Type: 2, ConvertCntGap: -189, DayBalance: -84, AdId: 1746892016018507},
		{Type: 2, ConvertCntGap: -81, DayBalance: -84, AdId: 1746933260430397},
		{Type: 2, ConvertCntGap: -611, DayBalance: -84, AdId: 1746934485520435},
		{Type: 2, ConvertCntGap: -122, DayBalance: -84, AdId: 1746935823875091},
		{Type: 2, ConvertCntGap: -561, DayBalance: -84, AdId: 1746936008432643},
		{Type: 2, ConvertCntGap: -48, DayBalance: -84, AdId: 1746936166974468},
		{Type: 2, ConvertCntGap: -624, DayBalance: -83, AdId: 1746992351819806},
		{Type: 2, ConvertCntGap: -993, DayBalance: -83, AdId: 1747029908462644},
		{Type: 2, ConvertCntGap: -229, DayBalance: -82, AdId: 1747105188102211},
		{Type: 2, ConvertCntGap: -246, DayBalance: -82, AdId: 1747105286947875},
		{Type: 2, ConvertCntGap: -588, DayBalance: -82, AdId: 1747114048411677},
		{Type: 2, ConvertCntGap: -61, DayBalance: -82, AdId: 1747119473452078},
		{Type: 2, ConvertCntGap: -52, DayBalance: -82, AdId: 1747124226781208},
		{Type: 2, ConvertCntGap: 10, DayBalance: -81, AdId: 1747161195815028},
		{Type: 2, ConvertCntGap: -315, DayBalance: -81, AdId: 1747217397074995},
		{Type: 2, ConvertCntGap: -101, DayBalance: -81, AdId: 1747217363895339},
		{Type: 2, ConvertCntGap: 10, DayBalance: -81, AdId: 1747217472756756},
		{Type: 2, ConvertCntGap: -20, DayBalance: -81, AdId: 1747217434224707},
		{Type: 2, ConvertCntGap: -43, DayBalance: -81, AdId: 1747218955872260},
		{Type: 2, ConvertCntGap: -25, DayBalance: -81, AdId: 1747219143790622},
		{Type: 2, ConvertCntGap: 10, DayBalance: -80, AdId: 1747264914554915},
		{Type: 2, ConvertCntGap: 10, DayBalance: -80, AdId: 1747279515238420},
		{Type: 2, ConvertCntGap: 10, DayBalance: -80, AdId: 1747279686495300},
		{Type: 2, ConvertCntGap: 10, DayBalance: -80, AdId: 1747300856661035},
		{Type: 2, ConvertCntGap: 10, DayBalance: -77, AdId: 1747534212452387},
		{Type: 2, ConvertCntGap: 10, DayBalance: -77, AdId: 1747534272940059},
		{Type: 2, ConvertCntGap: -450, DayBalance: -68, AdId: 1748338408043619},
		{Type: 2, ConvertCntGap: -78, DayBalance: -68, AdId: 1748346148070411},
		{Type: 2, ConvertCntGap: 10, DayBalance: -68, AdId: 1748351212886071},
		{Type: 2, ConvertCntGap: 10, DayBalance: -68, AdId: 1748362753045659},
		{Type: 2, ConvertCntGap: 10, DayBalance: -68, AdId: 1748363295234152},
		{Type: 2, ConvertCntGap: 10, DayBalance: -68, AdId: 1748390520623171},
		{Type: 2, ConvertCntGap: -184, DayBalance: -67, AdId: 1748438378684456},
		{Type: 2, ConvertCntGap: -78, DayBalance: -67, AdId: 1748439363497086},
		{Type: 2, ConvertCntGap: -247, DayBalance: -67, AdId: 1748442264483883},
		{Type: 2, ConvertCntGap: 10, DayBalance: -66, AdId: 1748523078798375},
		{Type: 2, ConvertCntGap: -309, DayBalance: -63, AdId: 1748827817473171},
		{Type: 2, ConvertCntGap: -497, DayBalance: -63, AdId: 1748842288462919},
		{Type: 2, ConvertCntGap: -344, DayBalance: -63, AdId: 1748842437553179},
		{Type: 2, ConvertCntGap: -368, DayBalance: -63, AdId: 1748842573173860},
		{Type: 2, ConvertCntGap: 10, DayBalance: -63, AdId: 1748849604917309},
		{Type: 2, ConvertCntGap: 10, DayBalance: -62, AdId: 1748885522151443},
		{Type: 2, ConvertCntGap: 10, DayBalance: -62, AdId: 1748885564828763},
		{Type: 2, ConvertCntGap: 10, DayBalance: -62, AdId: 1748890017160215},
		{Type: 2, ConvertCntGap: 10, DayBalance: -62, AdId: 1748890502405147},
		{Type: 2, ConvertCntGap: 10, DayBalance: -61, AdId: 1749003018699806},
		{Type: 2, ConvertCntGap: 2, DayBalance: -57, AdId: 1749379666640955},
		{Type: 2, ConvertCntGap: -806, DayBalance: -57, AdId: 1749379735914509},
		{Type: 2, ConvertCntGap: -314, DayBalance: -57, AdId: 1749380910444595},
		{Type: 2, ConvertCntGap: -201, DayBalance: -55, AdId: 1749567064722459},
		{Type: 2, ConvertCntGap: -526, DayBalance: -54, AdId: 1749624227612797},
		{Type: 2, ConvertCntGap: -243, DayBalance: -53, AdId: 1749700351243299},
		{Type: 2, ConvertCntGap: -1006, DayBalance: -52, AdId: 1749841939392580},
		{Type: 2, ConvertCntGap: -87, DayBalance: -52, AdId: 1749842037241864},
		{Type: 2, ConvertCntGap: -430, DayBalance: -52, AdId: 1749843817417795},
		{Type: 2, ConvertCntGap: -418, DayBalance: -51, AdId: 1749926656356398},
		{Type: 2, ConvertCntGap: -1071, DayBalance: -49, AdId: 1750066471479412},
		{Type: 2, ConvertCntGap: -529, DayBalance: -49, AdId: 1750067977260100},
		{Type: 2, ConvertCntGap: -509, DayBalance: -48, AdId: 1750155471381511},
		{Type: 2, ConvertCntGap: -1203, DayBalance: -47, AdId: 1750249071754243},
		{Type: 2, ConvertCntGap: -1467, DayBalance: -45, AdId: 1750472818414627},
		{Type: 2, ConvertCntGap: -310, DayBalance: -44, AdId: 1750537667301412},
		{Type: 2, ConvertCntGap: -105, DayBalance: -42, AdId: 1750705647848532},
		{Type: 2, ConvertCntGap: -339, DayBalance: -42, AdId: 1750705886739460},
		{Type: 2, ConvertCntGap: -69, DayBalance: -39, AdId: 1750972516676660},
		{Type: 2, ConvertCntGap: -145, DayBalance: -39, AdId: 1750974397358087},
		{Type: 2, ConvertCntGap: -864, DayBalance: -39, AdId: 1750987572919316},
		{Type: 2, ConvertCntGap: -251, DayBalance: -39, AdId: 1750987638712408},
		{Type: 2, ConvertCntGap: -107, DayBalance: -38, AdId: 1751091540230174},
		{Type: 2, ConvertCntGap: -116, DayBalance: -38, AdId: 1751112062027835},
		{Type: 2, ConvertCntGap: -3685, DayBalance: -37, AdId: 1751159800082443},
		{Type: 2, ConvertCntGap: -675, DayBalance: -37, AdId: 1751179726575620},
		{Type: 2, ConvertCntGap: -1942, DayBalance: -36, AdId: 1751245612228723},
		{Type: 2, ConvertCntGap: -5290, DayBalance: -36, AdId: 1751253887368231},
		{Type: 2, ConvertCntGap: -53, DayBalance: -36, AdId: 1751254249020475},
		{Type: 2, ConvertCntGap: 10, DayBalance: -35, AdId: 1751338333284420},
		{Type: 2, ConvertCntGap: -161, DayBalance: -35, AdId: 1751338496143406},
		{Type: 2, ConvertCntGap: -44, DayBalance: -35, AdId: 1751381278145627},
		{Type: 2, ConvertCntGap: -336, DayBalance: -35, AdId: 1751381542970404},
		{Type: 2, ConvertCntGap: -82, DayBalance: -33, AdId: 1751512840554523},
		{Type: 2, ConvertCntGap: -658, DayBalance: -33, AdId: 1751513140884531},
		{Type: 2, ConvertCntGap: -522, DayBalance: -33, AdId: 1751516211455070},
		{Type: 2, ConvertCntGap: -10, DayBalance: -32, AdId: 1751601602722887},
		{Type: 2, ConvertCntGap: -107, DayBalance: -32, AdId: 1751602023969803},
		{Type: 2, ConvertCntGap: -157, DayBalance: -31, AdId: 1751697290934331},
		{Type: 2, ConvertCntGap: 0, DayBalance: -31, AdId: 1751698517482595},
		{Type: 2, ConvertCntGap: -1015, DayBalance: -29, AdId: 1751911310260243},
		{Type: 2, ConvertCntGap: -77, DayBalance: -29, AdId: 1751911620256798},
		{Type: 2, ConvertCntGap: -63, DayBalance: -29, AdId: 1751913149202462},
		{Type: 2, ConvertCntGap: -55, DayBalance: -29, AdId: 1751913390214243},
		{Type: 2, ConvertCntGap: -16, DayBalance: -29, AdId: 1751913573801019},
		{Type: 2, ConvertCntGap: -28, DayBalance: -29, AdId: 1751913717298199},
		{Type: 2, ConvertCntGap: -98, DayBalance: -29, AdId: 1751913912292430},
		{Type: 2, ConvertCntGap: -34, DayBalance: -29, AdId: 1751914078150707},
		{Type: 2, ConvertCntGap: -66, DayBalance: -29, AdId: 1751914359906315},
		{Type: 2, ConvertCntGap: -46, DayBalance: -28, AdId: 1752013580937275},
		{Type: 2, ConvertCntGap: 0, DayBalance: -28, AdId: 1752013646631000},
		{Type: 2, ConvertCntGap: 1, DayBalance: -28, AdId: 1752013614832718},
		{Type: 2, ConvertCntGap: 0, DayBalance: -28, AdId: 1752015533311019},
		{Type: 2, ConvertCntGap: -39, DayBalance: -28, AdId: 1752015641888792},
		{Type: 2, ConvertCntGap: 8, DayBalance: -28, AdId: 1752015687000075},
		{Type: 2, ConvertCntGap: -1434, DayBalance: -28, AdId: 1752015782585380},
		{Type: 2, ConvertCntGap: -769, DayBalance: -28, AdId: 1752015891944491},
		{Type: 2, ConvertCntGap: -45, DayBalance: -28, AdId: 1752018468996189},
		{Type: 2, ConvertCntGap: 1, DayBalance: -28, AdId: 1752018549661739},
		{Type: 2, ConvertCntGap: -106, DayBalance: -28, AdId: 1752018531407931},
		{Type: 2, ConvertCntGap: 8, DayBalance: -27, AdId: 1752063785954388},
		{Type: 2, ConvertCntGap: 2, DayBalance: -27, AdId: 1752063971855427},
		{Type: 2, ConvertCntGap: -290, DayBalance: -27, AdId: 1752070005877806},
		{Type: 2, ConvertCntGap: -27, DayBalance: -27, AdId: 1752069961249853},
		{Type: 2, ConvertCntGap: -5, DayBalance: -26, AdId: 1752177139259400},
		{Type: 2, ConvertCntGap: -3, DayBalance: -26, AdId: 1752177216953384},
		{Type: 2, ConvertCntGap: 10, DayBalance: -21, AdId: 1752602568478734},
		{Type: 2, ConvertCntGap: -12, DayBalance: -21, AdId: 1752602916161598},
		{Type: 2, ConvertCntGap: -213, DayBalance: -21, AdId: 1752604639288340},
		{Type: 2, ConvertCntGap: -69, DayBalance: -21, AdId: 1752605066700827},
		{Type: 2, ConvertCntGap: -234, DayBalance: -21, AdId: 1752638481074315},
		{Type: 2, ConvertCntGap: -355, DayBalance: -21, AdId: 1752638642636936},
		{Type: 2, ConvertCntGap: -419, DayBalance: -21, AdId: 1752638715961363},
		{Type: 2, ConvertCntGap: -276, DayBalance: -21, AdId: 1752639077727267},
		{Type: 2, ConvertCntGap: -97, DayBalance: -21, AdId: 1752640725598259},
		{Type: 2, ConvertCntGap: -54, DayBalance: -21, AdId: 1752640945169476},
		{Type: 2, ConvertCntGap: 10, DayBalance: -20, AdId: 1752706505155646},
		{Type: 2, ConvertCntGap: -301, DayBalance: -20, AdId: 1752719498286107},
		{Type: 2, ConvertCntGap: -10, DayBalance: -19, AdId: 1752790838151236},
		{Type: 2, ConvertCntGap: -408, DayBalance: -19, AdId: 1752790889213955},
		{Type: 2, ConvertCntGap: -660, DayBalance: -19, AdId: 1752790904217611},
		{Type: 2, ConvertCntGap: -60, DayBalance: -19, AdId: 1752811606804493},
		{Type: 2, ConvertCntGap: -36, DayBalance: -19, AdId: 1752811575132283},
		{Type: 2, ConvertCntGap: -165, DayBalance: -19, AdId: 1752811683753012},
		{Type: 2, ConvertCntGap: 10, DayBalance: -16, AdId: 1753081749720084},
		{Type: 2, ConvertCntGap: 10, DayBalance: -16, AdId: 1753081666931747},
		{Type: 2, ConvertCntGap: -266, DayBalance: -16, AdId: 1753079449658392},
		{Type: 2, ConvertCntGap: 10, DayBalance: -15, AdId: 1753154129995779},
		{Type: 2, ConvertCntGap: -8, DayBalance: -16, AdId: 1753081335536692},
		{Type: 2, ConvertCntGap: -29, DayBalance: -12, AdId: 1753466423985176},
		{Type: 2, ConvertCntGap: 8, DayBalance: -11, AdId: 1753557144764483},
		{Type: 2, ConvertCntGap: 8, DayBalance: -11, AdId: 1753558717866211},
		{Type: 2, ConvertCntGap: -48, DayBalance: -9, AdId: 1753695366118455},
		{Type: 2, ConvertCntGap: 10, DayBalance: -7, AdId: 1753922193480711},
		{Type: 2, ConvertCntGap: 8, DayBalance: -6, AdId: 1753981887312907},
		{Type: 2, ConvertCntGap: -8, DayBalance: -5, AdId: 1754021799828571},
		{Type: 2, ConvertCntGap: 9, DayBalance: -4, AdId: 1754142858857595},
		{Type: 2, ConvertCntGap: 10, DayBalance: -4, AdId: 1754145086086205},
		{Type: 2, ConvertCntGap: 2, DayBalance: -4, AdId: 1754150862186515},
		{Type: 2, ConvertCntGap: -110, DayBalance: -4, AdId: 1754159405093899},
		{Type: 2, ConvertCntGap: 0, DayBalance: -3, AdId: 1754242283150376},
		{Type: 2, ConvertCntGap: 6, DayBalance: -2, AdId: 1754318128782403},
		{Type: 2, ConvertCntGap: -5, DayBalance: -2, AdId: 1754318497518619},
		{Type: 2, ConvertCntGap: 8, DayBalance: -2, AdId: 1754325985531917},
		{Type: 2, ConvertCntGap: 9, DayBalance: -2, AdId: 1754337522407454},
		{Type: 2, ConvertCntGap: -317, DayBalance: -1, AdId: 1754461010862123},
		{Type: 2, ConvertCntGap: -291, DayBalance: -1, AdId: 1754462676066323}}
	for _, v := range diffConvertCnts {
		fmt.Printf("%d,", v.AdId)
	}*/
	now := time.Now().Format("2006-01-02 15:04:00")
	fmt.Printf("%s", now)
}

type DiffConvertCnt struct {
	Type          int64 //类型 1=直播 2=支付ROI
	ConvertCntGap int64 //转换数缺口
	DayBalance    int64 //结算剩余天数
	AdId          int64
}

func SliceToGroup(slice []int64, lens int) [][]int64 {
	totalLen := len(slice)
	slices := make([][]int64, 0)
	s := make([]int64, 0)
	if totalLen == 0 {
		return slices
	}
	page := totalLen / lens
	m := totalLen % lens
	for i := 0; i <= page; i++ {
		maxIndex := lens * (i + 1)
		if totalLen < maxIndex {
			s = slice[i*lens : lens*i+m]
		} else {
			s = slice[i*lens : maxIndex]
		}
		slices = append(slices, s)
	}
	return slices
}
