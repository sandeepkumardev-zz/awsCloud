package utils

import (
	"awsCloud/database/config"
	"encoding/json"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func CreateOTP(id string) (int, error) {
	num := rand.Intn(1000000)

	json, err := json.Marshal(int(num))
	if err != nil {
		return 0, err
	}
	rdExTime, _ := strconv.Atoi(os.Getenv("REDIS_EXPIRATION_TIME"))
	err = config.RedisClient.Set(id, json, time.Duration(rdExTime)*time.Minute).Err()
	if err != nil {
		return 0, err
	}

	return num, nil
}
