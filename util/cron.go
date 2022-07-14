package util

import (
	"github.com/beego/beego/v2/adapter/toolbox"
)

func InitTask() {
	tk1 := toolbox.NewTask("hd_15256002129", "0 */1 * * * *", StartOrders)

	toolbox.AddTask("hd_15256002129", tk1)
}
