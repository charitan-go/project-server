package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/charitan-go/project-server/ent"
	redispkg "github.com/charitan-go/project-server/pkg/redis"
)

const (
	SET_BY_ID_KEY = "project:id"
)

type ProjectRedisService interface {
	SetById(ctx context.Context, p *ent.Project) error
}

type projectRedisServiceImpl struct {
	redisSvc redispkg.RedisService
}

func NewProjectRedisService(redisSvc redispkg.RedisService) ProjectRedisService {
	return &projectRedisServiceImpl{redisSvc}
}

func (svc *projectRedisServiceImpl) getSetByIdKey(p *ent.Project) string {
	id := p.ReadableID.String()
	return fmt.Sprintf("%s:%s", SET_BY_ID_KEY, id)
}

// SetById implements ProjectRedisService.
func (svc *projectRedisServiceImpl) SetById(ctx context.Context, p *ent.Project) error {
	key := svc.getSetByIdKey(p)

	value, err := json.Marshal(p)
	if err != nil {
		log.Printf("Cannot parse project to string: %v\n", p)
		return err
	}

	err = svc.redisSvc.Set(ctx, key, value)
	if err != nil {
		log.Printf("Cannot set to redis: %v\n", err)
		return err
	}

	return nil
}
