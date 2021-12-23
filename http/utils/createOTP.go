package utils

import (
	"awsCloud/database/config"
	"encoding/json"
	"math/rand"
	"time"
)

func CreateOTP(id string) (int, error) {
	num := rand.Intn(1000000)

	json, err := json.Marshal(int(num))
	if err != nil {
		return 0, err
	}

	err = config.RedisClient.Set(id, json, 2*time.Minute).Err()
	if err != nil {
		return 0, err
	}

	return num, nil
}