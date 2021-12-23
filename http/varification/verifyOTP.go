package varification

import (
	"awsCloud/database/config"
	"fmt"
)

func VerifyOTP(id string, otp string) error {
	val, err := config.RedisClient.Get(id).Result()
	if err != nil {
		return fmt.Errorf("something went wrong")
	}

	if val == otp {
		return nil
	}

	return fmt.Errorf("invalid OTP")
}
