package admin

import (
	"github.com/fengjx/amisgo"
	"github.com/fengjx/go-halo/utils"

	"github.com/fengjx/lucky/integration/db"
	"github.com/fengjx/lucky/logic/sys/syspub"
)

var (
	// APIPrefix 管理后台接口前缀
	APIPrefix = "/admin"

	Router *amisgo.AdminRouter
)

func init() {
	// 设置默认数据库执行器
	amisgo.UseDefaultExecutor(db.GetDefaultDB())
	// 设置不能编辑的全局字段
	amisgo.UseOmitsEditColumns("ctime", "utime")

	Router = amisgo.NewAdminRouter()
	_ = Router.SetMenu([]*amisgo.Menu{
		{
			Label:    "Home",
			URL:      "/",
			Redirect: "/sys",
			Visible:  true,
		},
		{
			Label:   "系统",
			Visible: true,
			Children: []*amisgo.Menu{
				{
					Label:    "内容管理",
					Icon:     "fa fa-file-text",
					URL:      "/cms",
					Redirect: "/cms/news",
					Visible:  true,
					Children: []*amisgo.Menu{
						{
							MenuID:  "cms_news",
							Label:   "新闻管理",
							Icon:    "fa fa-file-text",
							URL:     "/cms/news",
							Visible: true,
						},
					},
				},
				{
					Label:    "系统管理",
					Icon:     "fa fa-wrench",
					URL:      "/sys",
					Redirect: "/sys/user",
					Visible:  true,
					Children: []*amisgo.Menu{
						{
							MenuID:  "sys_user",
							Label:   "用户管理",
							Icon:    "fa fa-user",
							URL:     "/sys/user",
							Visible: true,
						},
						{
							MenuID:  "sys_dict",
							Label:   "字典管理",
							Icon:    "fa fa-bars",
							URL:     "/sys/dict",
							Visible: true,
						},
						{
							MenuID:  "sys_config",
							Label:   "配置管理",
							Icon:    "fa fa-bars",
							URL:     "/sys/config",
							Visible: true,
						},
					},
				},
			},
		},
	})
}

// GetDictOptions 根据字典分组获取 GetOptions
func GetDictOptions(group string) amisgo.GetOptions {
	return func() ([]*amisgo.Option, error) {
		dict := syspub.DictAPI.GetGroupDict(group)
		ops := make([]*amisgo.Option, 0, len(dict))
		for _, d := range dict {
			ops = append(ops, &amisgo.Option{
				Label: d.Label,
				Value: utils.ToString(d.Value),
			})
		}
		return ops, nil
	}
}

// RegAdminCRUD 注册 crud 组件
func RegAdminCRUD(admins ...*amisgo.AdminCRUD) {
	Router.RegAdminCRUD(admins...)
}
