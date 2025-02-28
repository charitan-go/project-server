package redispkg

import "go.uber.org/fx"

var RedisModule = fx.Module("redis",
	fx.Provide(
		NewRedisService,
	),
)
