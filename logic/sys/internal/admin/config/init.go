package config

import (
	"github.com/fengjx/amisgo"
	"github.com/fengjx/luchen"
	"github.com/fengjx/luchen/log"
	"go.uber.org/zap"

	"github.com/fengjx/lucky/logic/admin"
)

func Init(hs *luchen.HTTPServer) {
	config, err := amisgo.NewAdminCRUD(
		"sys_config",
		"/sys/config",
		amisgo.WithAPIPrefix(admin.APIPrefix),
		amisgo.WithDBName("lucky"),
		amisgo.WithFields([]*amisgo.Field{
			{
				Name:      "id",
				Label:     "主键",
				ShowTable: true,
				ShowView:  true,
			},
			{
				Name:                "status",
				Label:               "状态",
				Type:                amisgo.TypeSelect,
				ShowTable:           true,
				ShowView:            true,
				CreateAble:          true,
				UpdateAble:          true,
				SearchAble:          true,
				SearchConditionType: amisgo.ConditionTypeEq,
				Options: &amisgo.Options{
					GetOptions: admin.GetDictOptions("sys_config.status"),
				},
			},
		}),
	)
	if err != nil {
		log.Panic("create config admin err", zap.Error(err))
	}
	admin.RegAdminCRUD(config)
}
