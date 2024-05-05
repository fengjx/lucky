package service

import (
	"context"

	"github.com/fengjx/luchen/log"
	"go.uber.org/zap"

	"github.com/fengjx/lucky/logic/sys/internal/dao"
	"github.com/fengjx/lucky/logic/sys/internal/data/dto"
	"github.com/fengjx/lucky/logic/sys/syspub"
)

var DictSvc = newDictService()

type dictService struct {
	dictMap map[string][]*syspub.DictDTO
}

func newDictService() *dictService {
	inst := &dictService{
		dictMap: map[string][]*syspub.DictDTO{},
	}
	return inst
}

func (svc *dictService) Refresh(ctx context.Context) {
	list, err := dao.SysDictDao.ListAll(ctx)
	if err != nil {
		log.ErrorCtx(ctx, "list all sys_config err", zap.Error(err))
		return
	}
	dictMap := map[string][]*syspub.DictDTO{}
	for _, item := range list {
		dictMap[item.Group] = append(dictMap[item.Group], dto.BuildDictDTO(item))
	}
	svc.dictMap = dictMap
	log.InfofCtx(ctx, "refresh sys_dict, size: %d", len(list))
}

func (svc *dictService) GetGroupDict(groups ...string) map[string][]*syspub.DictDTO {
	if len(groups) == 0 {
		return svc.dictMap
	}
	res := map[string][]*syspub.DictDTO{}
	for _, group := range groups {
		if dicts, ok := svc.dictMap[group]; ok {
			res[group] = dicts
		}
	}
	return res
}
