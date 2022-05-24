package util

import (
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/common/log"

	. "hd_web/models"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/robfig/cron"
)

type MyJob struct{}

var AccountMap = map[string]string{
	//"15323335582": "aa123456",
	"13345539412": "liyu1201",
	//"18818693510": "aa123456",
}
var TokenMap = map[string]string{}

func (job MyJob) Run() {
	logs.Info("职位统计任务，单任务")
	GetJobs()
}

func CronStart() {
	runCronStr, _ := beego.AppConfig.String("runCron")
	runCron, err := strconv.ParseBool(runCronStr)
	if err == nil && runCron {
		c := cron.New()
		cron, _ := beego.AppConfig.String("cron")
		spec := cron
		c.AddFunc(spec, func() {
			GetJobs()
			logs.Info("职位统计任务已执行！")
		})
		//c.AddJob(spec, MyJob{})
		c.Start()

		//select {}//作用是在main函数中阻塞不退出
	}
}

func StartWinOrders() {

	for {

		// 遍历用户
		for u, p := range AccountMap {

			go func(username, password string) {

				// 获取用户token
				token := ""
				if val, has := TokenMap[username]; has {
					token = val
				} else {
					time.Sleep(3 * time.Minute)
					token = LoginWithPassword(username, password)
				}
				if token == "" {
					return
				}
				TokenMap[username] = token

				// 获取商品列表
				goods := GetGoods()
				if len(goods) == 0 {
					time.Sleep(time.Minute)
				}

				for _, good := range goods {
					go func(good Goods, token, uName string) {
						goodsId := strings.Replace(good.Url, "https://mini.hndutyfree.com.cn/#/pages/publicPages/goodDetails/index?goodsId=", "", -1)

						// 获取商品详情
						gd := FindGoodsDetail(goodsId, token)
						if gd == nil || gd.Data == nil {
							return
						}

						if gd.Code == 1024 {
							token = ""
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
							if order == nil {
								return
							}

							// 获取积分详情
							point := GetMemberPointClosed(token)

							// 积分满足则重新下单
							if point != nil && int(math.Floor(point.Data.Data)) >= order.Data.PointMax && good.UsePoints == 1 {
								pointMax = strconv.Itoa(order.Data.PointMax)
								pointRemain = strconv.Itoa(int(math.Floor(point.Data.Data)))
								order = GetPrepareOrderWithGoods(goodsId, countString, pointMax, pointRemain, token)
							}

							// 确认订单
							ConfirmOrderWithGoods(order, goodsId, countString, pointMax, token)

							// 支付确认
							PayConfirm(order, token, uName, gd.Data.ProductName)

						}

					}(good, token, username)
				}

			}(u, p)

			time.Sleep(1 * time.Second)
		}

	}

}

func GetGoods() []Goods {
	goods := new(Goods)
	goodsList, err := goods.GetAll()
	if err != nil {
		log.Error(err)
		time.Sleep(1 * time.Minute)
		return nil
	}
	if goodsList == nil {
		time.Sleep(1 * time.Minute)
		return nil
	}

	return goodsList
}
