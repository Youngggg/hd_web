package util

import (
	"encoding/json"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/go-resty/resty/v2"

	. "hd_web/models"
)

var _resty *resty.Client

const (
	LoginWithPasswordURL     = "https://service.cdfhnms.com/mini/loginWithPassword"
	FindGoodsDetailUrl       = "https://service.cdfhnms.com/mini/findGoodsDetailByIdAlways"
	GetPrepareOrderUrl       = "https://service.cdfhnms.com/mini/getPrepareOrderWithGoods"
	GetMemberPointUrl        = "https://service.cdfhnms.com/mini/getMemberPointClosed"
	ConfirmOrderWithGoodsUrl = "https://service.cdfhnms.com/core/confirmOrderWithGoods"
	PayConfirmUrl            = "https://service.cdfhnms.com/infrastructure/payConfirm"

	hd_webhook = "https://oapi.dingtalk.com/robot/send?access_token=eeca2967f730f9a2c11d44f9d92d95cd4722d45348e5608e4556188863ec4058"
)

func init() {
	_resty = resty.New().
		SetTimeout(10 * time.Second).
		SetRetryCount(2).
		SetRetryWaitTime(1 * time.Second).
		SetRetryMaxWaitTime(1 * time.Second)
}

func GetRestyClient() *resty.Client {
	return _resty
}

// 登录
func LoginWithPassword(username, password string) string {

	result := LoginWithPasswordRes{}
	res, err := GetRestyClient().R().
		SetQueryParam("phone", username).
		SetQueryParam("password", password).
		SetResult(&result).
		Get(LoginWithPasswordURL)

	if err != nil {
		logs.Error(err)
	}

	if result.Data != nil {
		return result.Data.Token
	}

	logs.Info(res)
	return ""
}

// 获取商品详情
func FindGoodsDetail(goodsId, token string) *FindGoodsDetailRes {
	findGoodsDetailRes := FindGoodsDetailRes{}
	res, err := GetRestyClient().R().
		SetHeader("token", token).
		SetQueryParam("goodsId", goodsId).
		SetResult(&findGoodsDetailRes).
		Get(FindGoodsDetailUrl)

	if err != nil {
		logs.Error(err)
	}
	if findGoodsDetailRes.Data != nil {
		logs.Info(time.Now().Format("2006-01-02 15:04:05"), " | 商品Id: ", findGoodsDetailRes.Data.GoodsId, " | 商品数量: ", findGoodsDetailRes.Data.Count)
	} else {
		logs.Info(res)
	}
	return &findGoodsDetailRes
}

// 获取订单详情
func GetPrepareOrderWithGoods(goodsId, count, point, pointRemain, token string) *GetPrepareOrderWithGoodsRes {
	getPrepareOrderWithGoodsRes := GetPrepareOrderWithGoodsRes{}

	res, err := GetRestyClient().R().
		SetHeader("token", token).
		SetQueryParam("goodsId", goodsId).
		SetQueryParam("count", count).
		SetQueryParam("point", point).
		SetQueryParam("pointRemain", pointRemain).
		SetResult(&getPrepareOrderWithGoodsRes).
		Get(GetPrepareOrderUrl)
	if err != nil {
		logs.Error(err)
	}
	logs.Info(res)

	return &getPrepareOrderWithGoodsRes
}

// 获取积分详情
func GetMemberPointClosed(token string) *GetMemberPointRes {
	result := GetMemberPointRes{}
	res, err := GetRestyClient().R().
		SetHeader("token", token).
		SetHeader("stockId", "6922").
		SetResult(&result).
		Get(GetMemberPointUrl)

	if err != nil {
		logs.Error(err)
	}
	logs.Info(res)
	return &result
}

func ConfirmOrderWithGoods(order *GetPrepareOrderWithGoodsRes, goodsId, count, point, token string) {
	result := ConfirmOrderWithGoodsRes{}

	params := map[string]string{
		"prov":           order.Data.MemberAddress.Province,
		"city":           order.Data.MemberAddress.City,
		"area":           order.Data.MemberAddress.District,
		"receiveAddress": order.Data.MemberAddress.Address,
		"receiveName":    order.Data.MemberAddress.Name,
		"receivePhone":   order.Data.MemberAddress.Mobile,
		"stockId":        "6922",
		"point":          point,
		"mainOrderId":    order.Data.MainOrderId,
		"goodsId":        goodsId,
		"count":          count,
		"memberCoupons":  "",
		"mac":            order.Data.Mac,
		"terminalId":     "6",
	}
	res, err := GetRestyClient().R().
		SetHeader("token", token).
		SetHeader("stockId", "6922").
		SetHeader("content-type", "application/json").
		SetQueryParams(params).
		SetResult(&result).
		Post(ConfirmOrderWithGoodsUrl)

	if err != nil {
		logs.Error(err)
	}
	logs.Info(res)
}

func PayConfirm(order *GetPrepareOrderWithGoodsRes, token, username, productName string) {
	result := PayConfirmRes{}
	res, err := GetRestyClient().R().
		SetHeader("token", token).
		SetHeader("stockId", "6922").
		SetHeader("terminalId", "6").
		SetQueryParam("mainOrderId", order.Data.MainOrderId).
		SetResult(&result).
		Get(PayConfirmUrl)

	if err != nil {
		logs.Error(err)
	}
	if result.Code == 0 {
		DingdingWarning(username, productName)
	}
	logs.Info(res)
}

func DingdingWarning(username, productName string) {
	msg := Msg{
		MsgType: "markdown",
		Markdown: &Markdown{
			Title: "hd下单成功提示",
			Text: "> **hd下单成功账号:**" + username + "\n" +
				"\n" + "> **商品名称:**" + productName + "\n",
		},
		At: &At{
			AtMobiles: []string{"18818693510"},
			IsAtAll:   true,
		},
	}
	msgJson, err := json.Marshal(msg)
	if err != nil {
		logs.Error(err)
	}

	v, err := GetRestyClient().R().
		SetHeader("Content-Type", "application/json").
		SetBody(msgJson).
		Post(hd_webhook)
	if err != nil {
		logs.Error(err)
		return
	}
	logs.Info(v)
}
