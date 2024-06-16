package protocol

import "github.com/fengjx/lucky/logic/sys/syspub"

// AppDataResp 配置数据
type AppDataResp struct {
	Config map[string][]*syspub.ConfigDTO `json:"config"`
	Dict   map[string][]*syspub.DictDTO   `json:"dict"`
}

// App 管理后台配置
type App struct {
	Pages []*Menu `json:"pages"`
}

// Menu 页面配置
// 参考 https://baidu.github.io/amis/zh-CN/components/app#%E5%B1%9E%E6%80%A7%E8%AF%B4%E6%98%8E
type Menu struct {
	Label     string  `json:"label"`               // 菜单名称
	Icon      string  `json:"icon,omitempty"`      // 菜单图标
	URL       string  `json:"url,omitempty"`       // 页面路由路径
	Link      string  `json:"link,omitempty"`      // 跳转外部链接
	Redirect  string  `json:"redirect,omitempty"`  // 当命中当前页面时，跳转到目标页面。
	Visible   bool    `json:"visible"`             // 有些页面可能不想出现在菜单中，可以配置成 false，另外带参数的路由无需配置，直接就是不可见的。
	SchemaAPI string  `json:"schemaApi,omitempty"` // 页面配置
	Children  []*Menu `json:"children,omitempty"`  // 子菜单
}
