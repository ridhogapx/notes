package config

import (
	"fmt"
	"testing"

	"github.com/gin-contrib/sessions/redis"
)

func TestRedisStore(t *testing.T) {
	// @RageNeko26
	// Below test should "pass" if store initiator is established
	// Example hard coded connection
	// Our objective is to cache with minimal expiration
	// Because long expiration will affect API Perfomance.
	conf := RedisConfig{
		Address: "172.17.0.3",
		Port:    "6379",
	}

	secret := "foo"

	_, err := redis.NewStore(10, "tcp", fmt.Sprintf("%s:%s", conf.Address, conf.Port), "", []byte(secret))

	if err != nil {
		t.Error("Failed to declare store:", err)
	}

	fmt.Println("Foo")
}
