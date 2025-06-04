package user

import (
	"reflect"

	"github.com/fengjx/amisgo"
	"github.com/fengjx/luchen"
	"github.com/fengjx/luchen/log"
	"go.uber.org/zap"

	"github.com/fengjx/lucky/common/types"
	"github.com/fengjx/lucky/logic/admin"
	"github.com/fengjx/lucky/logic/sys/internal/protocol"
)

func Init(hs *luchen.HTTPServer) {
	pwdInput := amisgo.NewInputText()
	pwdInput.Name = "new_pwd"
	pwdInput.Label = "密码"
	pwdInput.Required = true

	form := amisgo.NewForm()
	form.Name = "pwd-form"
	form.API = amisgo.NewAPI()
	form.API.Method = "POST"
	form.API.URL = "${API_BASEURL}/admin/sys/user/update-pwd"
	form.Body = []amisgo.Component{
		pwdInput,
	}

	opPwd := amisgo.NewAction().WithLabel("重置密码").WithIcon("fa fa-key").WithDialog(
		amisgo.NewDialog().WithTitle("重置密码").WithBody(form),
	)

	user, err := amisgo.NewAdminCRUD(
		"sys_user",
		"/sys/user",
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
				Name:       "salt",
				Label:      "密码盐",
				ShowTable:  false,
				ShowView:   false,
				CreateAble: false,
				UpdateAble: false,
			},
			{
				Name:                "username",
				Label:               "用户名",
				Required:            true,
				ShowTable:           true,
				ShowView:            true,
				UpdateAble:          false,
				CreateAble:          true,
				SearchAble:          true,
				SearchConditionType: amisgo.ConditionTypeLike,
			},
			{
				Name:                "status",
				Label:               "状态",
				Type:                amisgo.TypeSelect,
				ShowTable:           true,
				ShowView:            true,
				UpdateAble:          false,
				SearchAble:          true,
				SearchConditionType: amisgo.ConditionTypeEq,
				Options: &amisgo.Options{
					GetOptions: admin.GetDictOptions("sys_user.status"),
				},
			},
		}),
		amisgo.WithOperations([]*amisgo.Action{
			opPwd,
		}),
	)
	if err != nil {
		log.Panic("create user admin err", zap.Error(err))
	}
	admin.RegAdminCRUD(user)

	initEndpoint(hs)
}

func initEndpoint(hs *luchen.HTTPServer) {
	hs.Handle(&luchen.EndpointDefine{
		Name:     "UserAdmin.UpdatePwd",
		Path:     "/admin/sys/user/update-pwd",
		ReqType:  reflect.TypeOf(&protocol.UpdateUserPwdReq{}),
		RspType:  reflect.TypeOf(&types.AddRsp{}),
		Endpoint: makeUpdatePwdEndpoint(),
	})
}
