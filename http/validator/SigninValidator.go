package validator

import (
	"awsCloud/database/models"
	"fmt"
	"strings"
)

func SignInValidator(user *models.SignInUser) error {
	// remove spaces
	user.Username = strings.TrimSpace(user.Username)
	user.Password = strings.TrimSpace(user.Password)

	if user.Username == "" {
		return fmt.Errorf("username is a required field")
	}

	if user.Password == "" {
		return fmt.Errorf("password is a required field")
	}

	return nil
}
