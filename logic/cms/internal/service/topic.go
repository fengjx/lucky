package service

import (
	"context"

	"github.com/fengjx/lucky/logic/cms/internal/data/dto"
	"github.com/fengjx/lucky/logic/sys/syspub"
)

var TopicSvc = &topicService{}

type topicService struct {
}

// ListAll 查询主图列表
func (s *topicService) ListAll(ctx context.Context) ([]*dto.Topic, error) {
	var topics []*dto.Topic
	err := syspub.ConfigAPI.GetConfig(syspub.ScopeApp, "cms.topics", &topics)
	if err != nil {
		return nil, err
	}
	return topics, nil
}
