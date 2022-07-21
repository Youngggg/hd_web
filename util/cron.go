package util

import (
	"github.com/beego/beego/v2/adapter/toolbox"
)

func InitTask() {
	tk1 := toolbox.NewTask("hd_15256002129", "* 35 17 * * *", StartOrders15256002129)
	tk2 := toolbox.NewTask("hd_13401159806", "* 35 17 * * *", StartOrders13401159806)
	tk3 := toolbox.NewTask("hd_13155347128", "* 35 17 * * *", StartOrders13155347128)

	toolbox.AddTask("hd_15256002129", tk1)
	toolbox.AddTask("hd_13401159806", tk2)
	toolbox.AddTask("hd_13155347128", tk3)
}
