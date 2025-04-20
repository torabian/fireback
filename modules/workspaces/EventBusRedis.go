package workspaces

// Redis implementation of the Socket Instance Manager

import (
	"context"
	"encoding/json"
	"log"

	redis "github.com/redis/go-redis/v9"
)

type RedisManager struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisManager(redisURL string) *RedisManager {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisURL,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("❌ Failed to connect to Redis at %s: %v", config.RedisEventsUrl, err)
	}

	return &RedisManager{
		client: rdb,
		ctx:    context.Background(),
	}
}

func (r *RedisManager) key(instanceId string) string {
	return "instance:" + instanceId + ":users"
}

func (r *RedisManager) AddUser(instanceId, userID string) error {
	return r.client.SAdd(r.ctx, r.key(instanceId), userID).Err()
}

func (r *RedisManager) RemoveUser(instanceId, userID string) error {
	return r.client.SRem(r.ctx, r.key(instanceId), userID).Err()
}

func (r *RedisManager) ListUsers(instanceId string) ([]string, error) {
	return r.client.SMembers(r.ctx, r.key(instanceId)).Result()
}

func (r *RedisManager) IsUserIn(instanceId, userID string) (bool, error) {
	return r.client.SIsMember(r.ctx, r.key(instanceId), userID).Result()
}

func (r *RedisManager) FireEvent(q QueryDSL, event Event) {
	content := event.Json()

	if err := r.client.Publish(ctx, EVENT_BUS_TOPIC, content).Err(); err != nil {
		log.Println("Publish error:", err)
	}
}

func (r *RedisManager) Subscribe(ctx context.Context, channel string) {

	sub := r.client.Subscribe(ctx, channel)
	ch := sub.Channel()

	for {

		select {
		case msg := <-ch:

			var event Event
			if err := json.Unmarshal([]byte(msg.Payload), &event); err != nil {
				log.Println("Invalid message:", err)
				continue
			}

			// Now, we need a good logic to determine who needs to get the message.
			// A lot of messages arrived, what to do with them?
			RouteEvent(event)
		}
	}
}
