package util

import (
	"math"
	"strconv"
	"strings"
	"sync"
	"time"

	. "hd_web/models"
)

func StartOrders() error {

	// 初始化商品
	InitGoods()

	for {

		for username, password := range AccountMap {
			//go func(u, p string) {
			StartOrdersWithUser(username, password)
			//}(username, password)
		}

		time.Sleep(1 * time.Second)

	}

}

func StartOrdersWithUser(username, password string) {

	// 初始化token && 用户重新登录
	InitJ(username, password)

	var wg sync.WaitGroup

	// 循环商品
	for _, good := range goods {

		if J.IsReset {
			break
		}

		wg.Add(1)

		go func(good Goods, uName string) {

			defer wg.Done()

			goodsId := strings.Replace(good.Url, "https://mini.hndutyfree.com.cn/#/pages/publicPages/goodDetails/index?goodsId=", "", -1)

			token := ""
			if val, has := J.TokenMap[uName]; has {
				token = val
			} else {
				return
			}

			// 获取商品详情
			gd := FindGoodsDetail(goodsId, token)
			if gd == nil {
				return
			}

			if gd.Code == 1024 || gd.Message == "需重新登录" {
				J.IsReset = true
				return
			}

			if gd.Data == nil {
				return
			}

			// 判断是否使用折扣并且无折扣价
			if good.IsDiscount == 1 && gd.Data.EstimatePrice == 0 {
				return
			}

			// 商品数量大于0开抢
			if gd.Data.Count > 0 {
				pointMax := "0"
				pointRemain := "0"
				countString := strconv.Itoa(good.Count)

				// 获取订单详情
				order := GetPrepareOrderWithGoods(goodsId, countString, pointMax, pointRemain, token)
				if order == nil || order.Code != 0 || order.Data == nil {
					return
				}

				// 获取积分详情
				point := GetMemberPointClosed(token)

				// 积分满足则重新下单
				if good.UsePoints == 1 && point != nil && point.Data != nil && int(math.Floor(point.Data.Data)) > 0 {
					if int(math.Floor(point.Data.Data)) > order.Data.PointMax {
						pointMax = strconv.Itoa(order.Data.PointMax)
						pointRemain = strconv.Itoa(int(math.Floor(point.Data.Data)) - order.Data.PointMax)

					} else {
						pointMax = strconv.Itoa(int(math.Floor(point.Data.Data)))
						pointRemain = "0"
					}
					order = GetPrepareOrderWithGoods(goodsId, countString, pointMax, pointRemain, token)
				}
				if order == nil || order.Code != 0 || order.Data == nil {
					return
				}

				// 获取优惠券详情
				coupon := GetOrderCouponWithGoods(order, goodsId, countString, token)
				// 有优惠券则重新下单
				couponId := ""
				maxCouponCount := 0
				if good.UseCoupon == 1 && coupon != nil && coupon.Data != nil && coupon.Data.Total > 0 {
					for _, c := range coupon.Data.Data {
						if int(c.Amount) > maxCouponCount {
							maxCouponCount = int(c.Amount)
							couponId = c.CouponId
						}
					}

					GetNextPrepareOrderWithGoods(order.Data.MainOrderId, countString, strconv.Itoa(maxCouponCount), pointMax, pointRemain, goodsId, token)
				}

				// 确认订单
				ConfirmOrderWithGoods(order, goodsId, countString, pointMax, couponId, token)

				// 支付确认
				PayConfirm(order, token, username, gd)

			}

		}(good, username)

	}
}
