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
	GetById(ctx context.Context, id string) (*ent.Project, error)
}

type projectRedisServiceImpl struct {
	redisSvc redispkg.RedisService
}

func NewProjectRedisService(redisSvc redispkg.RedisService) ProjectRedisService {
	return &projectRedisServiceImpl{redisSvc}
}

func (svc *projectRedisServiceImpl) getByIdKey(id string) string {
	return fmt.Sprintf("%s:%s", SET_BY_ID_KEY, id)
}

// SetById implements ProjectRedisService.
func (svc *projectRedisServiceImpl) SetById(ctx context.Context, p *ent.Project) error {
	key := svc.getByIdKey(p.ReadableID.String())

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

func (svc *projectRedisServiceImpl) GetById(ctx context.Context, id string) (*ent.Project, error) {
	key := svc.getByIdKey(id)

	result, err := svc.redisSvc.Get(ctx, key)
	if err != nil {
		log.Printf("Cannot get from redis: %v\n", err)
		return nil, err
	}

	// TODO: Full impl
	log.Printf("Result from get redis: %v\n", result)
	return nil, nil
}
