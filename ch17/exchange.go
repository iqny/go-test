package ch17

import (
	"go-zh/ch17/topclient"
	"go-zh/ch17/topclient/request/taobaologisticsaddresssearch"
	"go-zh/ch17/topclient/request/tmallexchangereceivegetrequest"
)

type Exchange struct {
	cli *top_client.Client
}

func (e Exchange) GetExchanges(request map[string]interface{}) (response map[string]interface{}, err error) {
	req := tmallexchangereceivegetrequest.New()
	req.SetFields(request["fields"].(string))
	req.SetPageSize(request["page_size"].(string))
	req.SetPageNo(request["page_no"].(string))
	req.SetStartGmtModifiedTime(request["start_gmt_modified_time"].(string))
	req.SetEndGmtModifiedTime("end_gmt_modified_time")
	response, err = e.cli.Execute(req, request["session"].(string))
	return
}

func (e Exchange) GetLogisticsAddress(request map[string]interface{}) (response map[string]interface{}, err error) {
	req := taobaologisticsaddresssearch.New()
	req.SetRdef(request["rdef"].(string))
	response, err = e.cli.Execute(req, request["session"].(string))
	return
}

func (e Exchange) GetRow() {
	panic("implement me")
}

func (e Exchange) Agree() {
	panic("implement me")
}

func (e Exchange) Refuse() {
	panic("implement me")
}

func (e Exchange) ConfirmDelivery() {
	panic("implement me")
}

func (e Exchange) ConfirmReceiptGoods() {
	panic("implement me")
}

func (e Exchange) ReturnGoodsRefuse() {
	panic("implement me")
}
