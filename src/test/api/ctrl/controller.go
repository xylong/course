package ctrl

import "github.com/shenyisyn/goft-gin/goft"

var (
	// Controllers 控制器
	Controllers=make([]goft.IClass,0)
)

// RegisterController 注册控制器
func RegisterController(class goft.IClass)  {
	Controllers=append(Controllers,class)
}
