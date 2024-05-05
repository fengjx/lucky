package service

import (
	"context"

	"github.com/fengjx/luchen/log"
	"go.uber.org/zap"

	"github.com/fengjx/lucky/logic/sys/internal/dao"
	"github.com/fengjx/lucky/logic/sys/internal/data/dto"
	"github.com/fengjx/lucky/logic/sys/syspub"
)

var ConfigSvc = newConfigService()

type configService struct {
	configMap map[string][]*syspub.ConfigDTO
}

func newConfigService() *configService {
	inst := &configService{
		configMap: make(map[string][]*syspub.ConfigDTO),
	}
	return inst
}

func (svc *configService) Refresh(ctx context.Context) {
	list, err := dao.SysConfigDao.ListAll(ctx)
	if err != nil {
		log.ErrorCtx(ctx, "list all sys_config err", zap.Error(err))
		return
	}
	configMap := map[string][]*syspub.ConfigDTO{}
	for _, item := range list {
		configMap[item.Scope] = append(configMap[item.Scope], dto.BuildConfigDTO(item))
	}
	svc.configMap = configMap
	log.InfofCtx(ctx, "refresh sys_config, size: %d", len(list))
}

func (svc *configService) ScopeConfig(scopes ...string) map[string][]*syspub.ConfigDTO {
	res := map[string][]*syspub.ConfigDTO{}
	for _, scope := range scopes {
		res[scope] = svc.configMap[scope]
	}
	return res
}

func (svc *configService) GetAllConfig() map[string][]*syspub.ConfigDTO {
	return svc.configMap
}
