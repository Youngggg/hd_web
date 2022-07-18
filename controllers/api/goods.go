package api

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"

	"hd_web/controllers"
	. "hd_web/models"
	"hd_web/util"
)

type GoodsAPIController struct {
	controllers.AdminBaseController
}

var token string

/**
 * 新增商品
 */
func (c *GoodsAPIController) AddGoods() {

	args := map[string]string{}
	body := c.Ctx.Input.RequestBody //接收raw body内容
	_ = json.Unmarshal(body, &args)

	url := args["url"]
	count := args["count"] // 只能接收url后面的参数，不能接收body中的参数
	isDiscount := args["is_discount"]
	usePoints := args["use_points"]
	useCoupon := args["use_coupon"]
	createdAt := time.Now()
	updatedAt := time.Now()

	var goods = new(Goods)
	goods.Url = url
	goods.Count, _ = strconv.Atoi(count)
	goods.IsDiscount, _ = strconv.Atoi(isDiscount)
	goods.UsePoints, _ = strconv.Atoi(usePoints)
	goods.UseCoupon, _ = strconv.Atoi(useCoupon)
	goods.Status = 1
	goods.CreatedAt = createdAt
	goods.UpdatedAt = updatedAt

	id, err := goods.Save()
	if nil != err {
		logs.Error(err)
		c.Data["json"] = controllers.ErrorData(err)
	} else {
		c.Data["json"] = controllers.SuccessData(id)
	}

	util.Datagoods = nil

	_ = c.ServeJSON()
}

func (c *GoodsAPIController) GoodsGet() {
	idstr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idstr)
	good := new(Goods)
	good.Id = uint(id)
	userObj, err := good.GetById()
	if err != nil {
		logs.Error(err)
		c.Data["json"] = controllers.ErrorData(err)
	}
	c.Data["json"] = controllers.SuccessData(userObj)
	c.ServeJSON()
}

/**
 * 修改商品
 */
func (c *GoodsAPIController) UpdateGoods() {
	goodsId, _ := c.GetInt("goodsId")
	url := c.GetString("url")
	count, _ := c.GetInt("count")
	isDiscount, _ := c.GetInt("is_discount")
	usePoints, _ := c.GetInt("use_points")
	useCoupon, _ := c.GetInt("use_coupon")
	updatedAt := time.Now()

	var goods = new(Goods)
	goods.Id = uint(goodsId)
	goods.Url = url
	goods.Count = count
	goods.IsDiscount = isDiscount
	goods.UsePoints = usePoints
	goods.UseCoupon = useCoupon
	goods.UpdatedAt = updatedAt
	upId, err := goods.Update()
	if nil != err {
		logs.Error(err)
		c.Data["json"] = controllers.ErrorData(err)
	} else {
		c.Data["json"] = controllers.SuccessData(upId)
	}
	util.Datagoods = nil
	c.ServeJSON()
}

/**
 * 商品列表接口
 */
func (c *GoodsAPIController) GoodsList() {

	start, _ := c.GetInt("start")
	perPage, _ := c.GetInt("perPage")

	goods := new(Goods)
	var goodsVOList = make([]GoodsVO, 10)
	list, total, err := goods.GetAllByCondition(start, perPage)
	if nil != err {
		logs.Error(err)
		c.Data["json"] = controllers.ErrorData(err)
	} else {
		for index, g := range list {
			goodsVo := new(GoodsVO)
			goodsVo.Goods = g
			goodsVo.CreatedAt = g.CreatedAt.Format("2006-01-02 15:04:05")
			goodsVo.UpdatedAt = g.UpdatedAt.Format("2006-01-02 15:04:05")

			goodsId := strings.Replace(g.Url, "https://mini.hndutyfree.com.cn/#/pages/publicPages/goodDetails/index?goodsId=", "", -1)
			if token == "" {
				token = util.LoginWithPassword("18818693510", "aa123456")
			}
			goodsDetail := util.FindGoodsDetail(goodsId, token)

			// 登录失效重试
			if goodsDetail != nil && goodsDetail.Code == 1024 {
				token = util.LoginWithPassword("18818693510", "aa123456")
				goodsDetail = util.FindGoodsDetail(goodsId, token)
			}

			if goodsDetail != nil && goodsDetail.Data != nil {
				goodsVo.Image = goodsDetail.Data.SmallImage
				goodsVo.Name = goodsDetail.Data.ProductName
			}
			goodsVOList = append(goodsVOList[:index], *goodsVo)
		}
		data := map[string]any{
			"result": goodsVOList,
			"total":  total,
		}
		c.Data["json"] = controllers.SuccessData(data)
	}

	_ = c.ServeJSON()
}

/**
 * 禁用商品
 */
func (c *GoodsAPIController) DisableGoods() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	logs.Error(id)
	var goods = new(Goods)
	goods.Id = uint(id)
	goods.Status = 0
	id64, err := goods.UpdateStatus()
	if nil != err {
		logs.Error(err)
		c.Data["json"] = controllers.ErrorData(err)
	} else {
		c.Data["json"] = controllers.SuccessData(id64)
	}
	util.Datagoods = nil
	c.ServeJSON()
}

/**
 * 启用商品
 */
func (c *GoodsAPIController) EnableGoods() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	var goods = new(Goods)
	goods.Id = uint(id)
	goods.Status = 1
	id64, err := goods.UpdateStatus()
	if nil != err {
		logs.Error(err)
		c.Data["json"] = controllers.ErrorData(err)
	} else {
		c.Data["json"] = controllers.SuccessData(id64)
	}
	util.Datagoods = nil
	c.ServeJSON()
}

/**
 * 删除接口
 */
func (c *GoodsAPIController) DeleteGoods() {
	goodsId, _ := c.GetInt("goodsId")
	logs.Error(goodsId)
	goods := new(Goods)
	goods.Id = uint(goodsId)
	id64, err := goods.Delete()
	if nil != err {
		logs.Error(err)
		c.Data["json"] = controllers.ErrorData(err)
	} else {
		c.Data["json"] = controllers.SuccessData(id64)
	}
	util.Datagoods = nil
	c.ServeJSON()
}
