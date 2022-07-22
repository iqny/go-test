package exchange_interface

type Exchange interface {
	// GetExchanges 获取换货列表
	GetExchanges(req map[string]interface{})(res map[string]interface{},err error)
	// GetLogisticsAddress 卖家默认收货地址
	GetLogisticsAddress(req map[string]interface{})(res map[string]interface{},err error)
	// GetRow 获取单笔交易信息
	GetRow()
	// Agree 同意
	Agree()
	// Refuse 拒绝
	Refuse()
	// ConfirmDelivery 卖家确认发货
	ConfirmDelivery()
	// ConfirmReceiptGoods 卖家确认收货
	ConfirmReceiptGoods()
	// ReturnGoodsRefuse 拒绝收货
	ReturnGoodsRefuse()
}
