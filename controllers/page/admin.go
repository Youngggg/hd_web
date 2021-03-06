package page

import (
	"log"
	"net/http"
	"runtime"

	"hd_web/controllers"
	. "hd_web/models"

	"github.com/PuerkitoBio/goquery"
	"github.com/beego/beego/v2/core/logs"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

/**
 * 给interface起个别名，这样是不是简短很多了
 */
type any = interface{}

type AdminPageController struct {
	controllers.AdminBaseController
}

//登录路由
func (c *AdminPageController) Login() {
	c.TplName = "admin/login.html"
}

/**
 * 登出
 */
func (c *AdminPageController) LoginOut() {
	c.DelSession("loginUser")
	c.Redirect("/admin/page/login", 302)
}

func (c *AdminPageController) Welcome() {
	c.TplName = "admin/index.html"
}

func (c *AdminPageController) UserAdd() {
	c.TplName = "admin/user-add.html"
}

/**
 * 用户列表路由
 */
func (c *AdminPageController) UserList() {
	c.Data["Title"] = "用户列表"
	c.TplName = "admin/user-list.html"
}

func (c *AdminPageController) GoodsAdd() {
	c.TplName = "admin/goods-add.html"
}

/**
 * 用户列表路由
 */
func (c *AdminPageController) GoodsList() {
	c.Data["Title"] = "商品列表"
	c.TplName = "admin/goods-list.html"
}

/**
 * 日志列表路由
 */
func (c *AdminPageController) Logs() {
	c.Data["Title"] = "日志列表"
	c.TplName = "admin/logs.html"
}

/**
 * 创建文章路由
 */
func (c *AdminPageController) CreateArticle() {
	category := &Category{}
	list, err := category.GetAllCategories()
	if err != nil {
		c.Data["categories"] = make(map[string]string)
	}
	logs.Debug("%+v", list)
	c.Data["categories"] = list
	c.TplName = "admin/article-add.html"
}

/**
 * 文章列表路由
 */
func (c *AdminPageController) ArticleList() {
	c.TplName = "admin/article-list.html"
}

/**
 * 编辑文章路由
 * /admin/page/articles
 */
func (c *AdminPageController) ArticleEdit() {
	article := new(Article)
	artId, _ := c.GetInt("id")
	article.Id = uint(artId)
	article.GetById()

	category := &Category{}
	list, err := category.GetAllCategories()
	if err != nil {
		c.Data["categories"] = make(map[string]string)
	}
	logs.Debug("%+v", list)
	c.Data["categories"] = list
	c.Data["article"] = article
	c.TplName = "admin/article-edit.html"
}

// 通讯录
func (c *AdminPageController) PNameView() {
	c.TplName = "admin/pname.html"
}

/**
 * 留言查看，审核 页面
 * /admin/page/messages
 */
func (c *AdminPageController) Messages() {
	article := new(Article)
	artId, _ := c.GetInt("id")
	article.Id = uint(artId)
	article.GetById()
	c.Data["article"] = article
	c.TplName = "admin/messages.html"
}

//JobCount 工作数量统计
func (c *AdminPageController) JobCount() {
	c.Data["title"] = "报表"
	c.TplName = "admin/job-count.html"
}

// IconList 列表
func (c *AdminPageController) IconList() {
	c.TplName = "admin/icons.html"
}

// Picture 图片检索
func (c *AdminPageController) Picture() {
	url := c.GetString("url")
	if url != "" {
		//请求html数据
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		} else {
			logs.Info("请求成功")
		}
		//转换数据为HTML对象模型
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		//查找元素
		text, _ := doc.Find("body").Html()
		c.Data["url"] = url
		c.Data["htmlContent"] = text
	} else {
		c.Data["htmlContent"] = ""
	}
	c.TplName = "admin/picture.html"
}

// ListPicture 图片列表路由
func (c *AdminPageController) ListPicture() {
	c.TplName = "admin/picture-list.html"
}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func (c *AdminPageController) SystemInfo() {
	osDic := make(map[string]interface{}, 0)
	osDic["goOs"] = runtime.GOOS
	osDic["arch"] = runtime.GOARCH
	osDic["mem"] = runtime.MemProfileRate
	osDic["compiler"] = runtime.Compiler
	osDic["version"] = runtime.Version()
	osDic["numGoroutine"] = runtime.NumGoroutine()

	dis, _ := disk.Usage("/")
	diskTotalGB := int(dis.Total) / GB
	diskFreeGB := int(dis.Free) / GB
	diskDic := make(map[string]interface{}, 0)
	diskDic["total"] = diskTotalGB
	diskDic["free"] = diskFreeGB

	mem, _ := mem.VirtualMemory()
	memUsedMB := int(mem.Used) / GB
	memTotalMB := int(mem.Total) / GB
	memFreeMB := int(mem.Free) / GB
	memUsedPercent := int(mem.UsedPercent)
	memDic := make(map[string]interface{}, 0)
	memDic["total"] = memTotalMB
	memDic["used"] = memUsedMB
	memDic["free"] = memFreeMB
	memDic["usage"] = memUsedPercent

	cpuDic := make(map[string]interface{}, 0)
	cpuDic["cpuNum"], _ = cpu.Counts(false)

	c.Data["os"] = osDic
	c.Data["mem"] = memDic
	c.Data["cpu"] = cpuDic
	c.Data["disk"] = diskDic

	c.TplName = "admin/system-info.html"
}

func (c *AdminPageController) Roles() {
	c.TplName = "admin/roles.html"
}

// 文章分类
func (c *AdminPageController) CategoryList() {
	c.TplName = "admin/category-list.html"
}
