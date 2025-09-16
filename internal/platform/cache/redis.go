package cache

import (
	"context"
	"encoding/json"
	"time"

	domainEntity "github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/entity"
	"github.com/go-redis/redis/v8"
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
	jsonData, err := json.Marshal(healthCheck)
	if err != nil {
		return
	}

	rdb.Set(ctx, key, jsonData, 10*time.Second).Result()
}

func GetRDB() *redis.Client {
	return rdb
}
