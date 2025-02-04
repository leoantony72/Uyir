package handler

import (
	"context"

	"github.com/leoantony72/Uyir/database"
	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var Db *redis.Client = database.StartRedis()
