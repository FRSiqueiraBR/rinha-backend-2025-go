package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	domainEntity "github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/entity"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func Get(key string) (domainEntity.HealthCheck, bool) {
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return domainEntity.HealthCheck{}, false
	} else if err != nil {
		return domainEntity.HealthCheck{}, false
	}

	var healthCheck domainEntity.HealthCheck
	err = json.Unmarshal([]byte(val), &healthCheck)
	if err != nil {
		return domainEntity.HealthCheck{}, false
	}

	return healthCheck, true
}

func Set(key string, healthCheck domainEntity.HealthCheck) {
	json, err := json.Marshal(healthCheck)
	if err != nil {
		return
	}

	rdb.Set(ctx, key, json, 10*time.Second).Result()
}
