package util

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"

	. "hd_web/models"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/robfig/cron"
)

type MyJob struct{}

type MJob struct {
	TokenMap map[string]string
	IsReset  bool
}

var (
	AccountMap = map[string]string{
		//"15323335582": "aa123456",
		//"13514105572": "aa123456",
		//"13401159806": "aa123456",
		//"13345539412": "liyu1201",
		"18818693510": "aa123456",
	}

	J *MJob

	ResetChan = make(chan bool)
)

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
		for username, password := range AccountMap {

			InitJ(username, password)

			// 获取商品列表
			goods := GetGoods()
			if len(goods) == 0 {
				time.Sleep(time.Minute)
			}

			fmt.Printf("J: %p", J)

			var wg sync.WaitGroup
			for _, good := range goods {

				if J.IsReset {
					break
				}

				go func(good Goods, uName string, job *MJob) {

					wg.Add(1)

					goodsId := strings.Replace(good.Url, "https://mini.hndutyfree.com.cn/#/pages/publicPages/goodDetails/index?goodsId=", "", -1)

					token := ""
					if val, has := J.TokenMap[uName]; has {
						token = val
					} else {
						wg.Done()
						return
					}

					// 获取商品详情
					gd := FindGoodsDetail(goodsId, token)
					if gd == nil {
						wg.Done()
						return
					}

					if gd.Code == 1024 || gd.Message == "需重新登录" {
						J.IsReset = true
						//ResetChan <- true
						fmt.Println("reset: ", J, job, ResetChan)
						wg.Done()
						return
					}

					if gd.Data == nil {
						return
					}

					// 判断是否使用折扣并且无折扣价
					if good.IsDiscount == 1 && gd.Data.EstimatePrice == 0 {
						wg.Done()
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
							wg.Done()
							return
						}

						// 获取积分详情
						point := GetMemberPointClosed(token)

						// 积分满足则重新下单
						if good.UsePoints == 1 && point != nil && point.Data != nil && int(math.Floor(point.Data.Data)) >= order.Data.PointMax {
							pointMax = strconv.Itoa(order.Data.PointMax)
							pointRemain = strconv.Itoa(int(math.Floor(point.Data.Data)))
							order = GetPrepareOrderWithGoods(goodsId, countString, pointMax, pointRemain, token)
						}

						// 确认订单
						ConfirmOrderWithGoods(order, goodsId, countString, pointMax, token)

						// 支付确认
						PayConfirm(order, token, username, gd)

						wg.Done()

					}

				}(good, username, J)

			}

			wg.Wait()
			time.Sleep(1 * time.Second)

		}

	}

}

func GetGoods() []Goods {
	goods := new(Goods)
	goodsList, err := goods.GetAll()
	if err != nil {
		logs.Error(err)
		time.Sleep(1 * time.Minute)
		return nil
	}
	if goodsList == nil {
		time.Sleep(1 * time.Minute)
		return nil
	}

	return goodsList
}

func InitJ(username, password string) {

	if J == nil {
		J = &MJob{}
		J.TokenMap = map[string]string{}
		J.TokenMap[username] = LoginWithPassword(username, password)
	}

	if J.IsReset {
		time.Sleep(2 * time.Minute)
		J = &MJob{}
		J.TokenMap = map[string]string{}
		J.TokenMap[username] = LoginWithPassword(username, password)
	}
	//select {
	//case reset := <-ResetChan:
	//
	//	fmt.Println("reset chan: ", reset)
	//	if reset {
	//		time.Sleep(2 * time.Minute)
	//		J = &MJob{}
	//		J.TokenMap = map[string]string{}
	//		J.TokenMap[username] = LoginWithPassword(username, password)
	//	}
	//
	//default:
	//
	//	// 获取用户token
	//	if J == nil {
	//		J = &MJob{}
	//		J.TokenMap = map[string]string{}
	//		J.TokenMap[username] = LoginWithPassword(username, password)
	//	} else {
	//
	//	}

}
