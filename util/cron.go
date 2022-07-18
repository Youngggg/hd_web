package util

import (
	"github.com/beego/beego/v2/adapter/toolbox"
)

func InitTask() {
	tk1 := toolbox.NewTask("hd_15256002129", "0 20 15 * * *", StartOrders15256002129)
	tk2 := toolbox.NewTask("hd_13401159806", "0 19 15 * * *", StartOrders13401159806)
	tk3 := toolbox.NewTask("hd_13155347128", "0 18 15 * * *", StartOrders13155347128)

	toolbox.AddTask("hd_15256002129", tk1)
	toolbox.AddTask("hd_13401159806", tk2)
	toolbox.AddTask("hd_13155347128", tk3)
}