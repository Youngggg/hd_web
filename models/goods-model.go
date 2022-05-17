package models

import (
	"strconv"
	"sync"

	"github.com/beego/beego/v2/core/logs"

	"github.com/beego/beego/v2/client/orm"
)

/**
 * 模型与数据库字段多少不一定要匹配
 */
type Goods struct {
	BaseModel
	Url        string `json:"url"`
	Count      int    `json:"count"`
	IsDiscount int    `json:"is_discount"` // 是否有折扣再抢 0否，1是
	UsePoints  int    `json:"use_points"`  // 是否使用积分 0否，1是
	UseCoupon  int    `json:"use_coupon"`  // 是否使用优惠券 0否，1是
	Status     int    `json:"status"`      // 状态 0下线，1上线
}

type GoodsVO struct {
	Goods
	Image     string `json:"image"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

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

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Goods))
	// 如果使用 orm.QuerySeter 进行高级查询的话，这个是必须的。
	// 反之，如果只使用 Raw 查询和 map struct，是无需这一步的。
}

// 添加用户
func (goods *Goods) Save() (int64, error) {
	//	var o Ormer
	o := orm.NewOrm()
	// 每次操作都需要新建一个Ormer变量，当然也可以全局设置
	// 需要 切换数据库 和 事务处理 的话，不要使用全局保存的 Ormer 对象。
	return o.Insert(goods)
}

// 通过id查找商品
func (goods *Goods) GetById() (*Goods, error) {
	o := orm.NewOrm()
	err := o.Read(goods, "id")
	return goods, err
}

// 获取商品列表
func (goods *Goods) GetAll() ([]Goods, error) {
	o := orm.NewOrm()
	var goodsList []Goods
	num, err := o.Raw("SELECT * FROM goods").QueryRows(&goodsList)
	logs.Info("查询到", num, "条数据")
	return goodsList, err

}

// 获取用户列表
func (goods *Goods) GetAllByCondition(start, perPage int) (goodsList []Goods, total int64, newError error) {
	o := orm.NewOrm()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		var sql = "SELECT * FROM goods "
		sql += " order by id desc"
		sql += " LIMIT ?, ?"
		_, err := o.Raw(sql, strconv.Itoa(start), strconv.Itoa(perPage)).QueryRows(&goodsList)
		if err != nil {
			newError = err
		}
	}()
	go func() {
		defer wg.Done()
		var sql = "SELECT COUNT(0) FROM goods "
		err := o.Raw(sql).QueryRow(&total)
		if err != nil {
			newError = err
		}
		logs.Info("mysql row affected nums: ", total)
	}()
	wg.Wait()
	return goodsList, total, newError
}

// 通过id修改用户
func (goods *Goods) Update() (int64, error) {
	o := orm.NewOrm()
	id, err := o.Update(goods, "url", "count", "is_discount", "use_points", "use_coupon", "updated_at") // 要修改的对象和需要修改的字段
	return id, err
}

// 修改商品状态
func (goods *Goods) UpdateStatus() (int64, error) {
	o := orm.NewOrm()
	id, err := o.Update(goods, "status") // 要修改的对象和需要修改的字段
	if err != nil {
		return id, err
	}
	return id, nil
}

// 通过id删除用户
func (goods *Goods) Delete() (int64, error) {
	o := orm.NewOrm()
	id, err := o.Delete(goods)
	return id, err
}
