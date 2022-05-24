package models

type FindGoodsDetailRes struct {
	Code int                  `json:"code"`
	Data *FindGoodsDetailData `json:"data"`
}

type FindGoodsDetailData struct {
	GoodsId       string  `json:"goodsId"`
	NeedCheck     bool    `json:"needCheck"`
	SalesPrice    float64 `json:"salesPrice"`
	EstimatePrice float64 `json:"estimatePrice"`
	IsPackage     int     `json:"isPackage"`
	Active        int     `json:"active"`
	Count         int     `json:"count"`
	SmallImage    string  `json:"smallImage"`
	ProductName   string  `json:"productName"`
}

type LoginWithPasswordRes struct {
	Code int                    `json:"code"`
	Data *LoginWithPasswordData `json:"data"`
}

type LoginWithPasswordData struct {
	Token    string `json:"token"`
	Phone    string `json:"phone"`
	BindId   string `json:"bindId"`
	IsWorker int    `json:"isWorker"`
}

type GetPrepareOrderWithGoodsRes struct {
	Code int                              `json:"code"`
	Data *GetPrepareOrderWithGoodsResData `json:"data"`
}

type GetPrepareOrderWithGoodsResData struct {
	TotalCount         int            `json:"totalCount"`
	MainOrderId        string         `json:"mainOrderId"`
	OriginalAmount     float64        `json:"originalAmount"`
	DealAmount         float64        `json:"dealAmount"`
	DiscountAmount     float64        `json:"discountAmount"`
	NeedPayAmount      float64        `json:"needPayAmount"`
	PointMax           int            `json:"pointMax"`
	CouponAmount       int            `json:"couponAmount"`
	MemberCouponAmount int            `json:"memberCouponAmount"`
	NeedCheck          bool           `json:"needCheck"`
	MemberAddress      *MemberAddress `json:"memberAddress"`
	Mac                string         `json:"mac"`
	State              int            `json:"state"`
}

type MemberAddress struct {
	Id        string `json:"id"`
	AddressId string `json:"addressId"`
	MemberId  string `json:"memberId"`
	Name      string `json:"name"`
	Province  string `json:"province"`
	City      string `json:"city"`
	District  string `json:"district"`
	Address   string `json:"address"`
	Mobile    string `json:"mobile"`
}

type GetMemberPointRes struct {
	Code int                    `json:"code"`
	Data *GetMemberPointResData `json:"data"`
}

type GetMemberPointResData struct {
	Check       int     `json:"check"`
	Data        float64 `json:"data"`
	Closed      int     `json:"closed"`
	ClosedPoint float64 `json:"closedPoint"`
}

type ConfirmOrderWithGoodsRes struct {
	Code      int    `json:"code"`
	Data      bool   `json:"data"`
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
	Addition  struct {
		Message string `json:"message"`
	} `json:"addition"`
}

type PayConfirmRes struct {
	Code int `json:"code"`
	Data struct {
		MainOrderId string  `json:"mainOrderId"`
		EndTime     string  `json:"endTime"`
		CreateTime  string  `json:"createTime"`
		TimeRemain  int     `json:"timeRemain"`
		NeedPay     float64 `json:"needPay"`
	}
}

type Msg struct {
	MsgType  string    `json:"msgtype"`
	Markdown *Markdown `json:"markdown"`
	At       *At       `json:"at"`
}

type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type At struct {
	AtMobiles []string `json:"atMobiles,omitempty"`
	IsAtAll   bool     `json:"isAtAll"`
}
