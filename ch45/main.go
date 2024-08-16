package main

import (
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"xorm.io/xorm"
	"xorm.io/xorm/caches"

	_ "github.com/go-sql-driver/mysql"
)

type Ar struct {
	SkuSn string `json:"sku_sn"`
}
type Res map[string]interface{}

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	r := gin.Default()
	ginpprof.Wrap(r)
	cacher := caches.NewLRUCacher(caches.NewMemoryStore(), 1000)
	var err error
	var engine *xorm.Engine
	engine, err = xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/local_test?charset=utf8&interpolateParams=true")
	if err != nil {
		log.Fatalln(err)
	}
	engine.SetDefaultCacher(cacher)
	r.GET("/", func(c *gin.Context) {
		ar := Ar{}
		engine.Where("sku_sn = ?", "202115009040100144").Get(&ar)
		c.String(http.StatusOK, ar.SkuSn)
	})
	r.POST("/dou", func(c *gin.Context) {
		/*res := Res{}
		err := c.ShouldBindJSON(&res)
		fmt.Println(err, res)*/
		c.String(http.StatusOK, `{
  "order_list_get_response": {
    "has_next": false,
    "order_list": [
      {
        "address": "1",
        "address_mask": "str",
        "after_sales_status": 1,
        "bonded_warehouse": "str",
        "buyer_memo": "1",
        "capital_free_discount": 1.0,
        "card_info_list": [
          {
            "card_no": "1",
            "mask_password": "1"
          }
        ],
        "cat_id_1": 1,
        "cat_id_2": 1,
        "cat_id_3": 1,
        "cat_id_4": 1,
        "city": "1",
        "city_id": 1,
        "confirm_status": 1,
        "confirm_time": "1",
        "consolidate_info": {
          "consolidate_type": 0
        },
        "country": "1",
        "country_id": 1,
        "created_time": "1",
        "delivery_home_value": 1.0,
        "delivery_install_value": 1.0,
        "delivery_one_day": 0,
        "discount_amount": 1.0,
        "duo_duo_pay_reduction": 0.0,
        "duoduo_wholesale": 1,
        "extra_delivery_list": [
          {
            "logistics_id": 0,
            "tracking_number": "str"
          }
        ],
        "free_sf": 1,
        "gift_delivery_list": [
          {
            "logistics_id": 0,
            "tracking_number": "str"
          }
        ],
        "gift_list": [
          {
            "goods_count": 0,
            "goods_id": 0,
            "goods_img": "str",
            "goods_name": "str",
            "goods_price": 0.0,
            "goods_spec": "str",
            "outer_goods_id": "str",
            "outer_id": "str",
            "sku_id": 0
          }
        ],
        "goods_amount": 1.0,
        "group_order_id": 0,
        "group_role": 0,
        "group_status": 1,
        "home_delivery_type": 1,
        "home_install_value": 1.0,
        "inner_transaction_id": "1",
        "invoice_status": 1,
        "is_lucky_flag": 1,
        "is_pre_sale": 1,
        "is_stock_out": 1,
        "item_list": [
          {
            "goods_count": 1,
            "goods_id": "1",
            "goods_img": "1",
            "goods_name": "1",
            "goods_price": 1.0,
            "goods_spec": "1",
            "outer_goods_id": "1",
            "outer_id": "1",
            "sku_id": "1"
          }
        ],
        "last_ship_time": "1",
        "logistics_id": 1,
        "mkt_biz_type": 1,
        "only_support_replace": 1,
        "open_address_id": "str",
        "order_change_amount": 0.0,
        "order_depot_info": {
          "depot_code": "1",
          "depot_id": 1,
          "depot_name": "1",
          "depot_type": 1,
          "ware_id": 1,
          "ware_name": "1",
          "ware_sn": "1",
          "ware_sub_info_list": [
            {
              "ware_id": 1,
              "ware_name": "1",
              "ware_quantity": 1,
              "ware_sn": "1"
            }
          ],
          "ware_type": 1
        },
        "order_sn": "1",
        "order_status": 1,
        "order_tag_list": [
          {
            "name": "return_freight_payer",
            "value": 1
          }
        ],
        "pay_amount": 1.0,
        "pay_no": "1",
        "pay_time": "1",
        "pay_type": "1",
        "platform_discount": 1.0,
        "postage": 1.0,
        "pre_sale_time": "1",
        "promise_delivery_time": "str",
        "promotion_detail_list": [
          {
            "discount_amount": 0.0,
            "promotion_type": 0
          }
        ],
        "province": "1",
        "province_id": 1,
        "receive_time": "1",
        "receiver_address": "1",
        "receiver_address_mask": "str",
        "receiver_name": "1",
        "receiver_name_mask": "str",
        "receiver_phone": "1",
        "receiver_phone_mask": "str",
        "refund_status": 1,
        "remark": "1",
        "remark_tag": 1,
        "remark_tag_name": "红色",
        "resend_delivery_list": [
          {
            "logistics_id": 0,
            "tracking_number": "str"
          }
        ],
        "return_freight_payer": 1,
        "risk_control_status": 0,
        "self_contained": 0,
        "seller_discount": 1.0,
        "service_fee_detail": [
          {
            "service_fee": 0.0,
            "service_name": "str"
          }
        ],
        "ship_additional_link_order": "str",
        "ship_additional_origin_order": "str",
        "shipping_time": "1",
        "shipping_type": 0,
        "step_order_info": {
          "advanced_paid_fee": 1.0,
          "step_discount_amount": 1.0,
          "step_paid_fee": 1.0,
          "step_trade_status": 1
        },
        "stock_out_handle_status": 1,
        "support_nationwide_warranty": 1,
        "town": "1",
        "town_id": 1,
        "tracking_number": "1",
        "trade_type": 1,
        "updated_at": "1",
        "urge_shipping_time": "str",
        "yyps_date": "str",
        "yyps_time": "str"
      }
    ],
    "total_count": 3
  }
}`)
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
