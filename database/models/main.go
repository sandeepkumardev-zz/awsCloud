package models

type SignInUser struct {
	Username string `dynamodbav:"username" validate:"required"`
	Password string `dynamodbav:"password" validate:"required"`
}

type User struct {
	Id          string `dynamodbav:"id"`
	Username    string `dynamodbav:"username" validate:"required,username,isValid"`
	Password    string `dynamodbav:"password" validate:"required,password"`
	PhoneNumber string `dynamodbav:"phone_number" validate:"required,phone,isNum"`
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type TokenDetails struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AtExpires    int64  `json:"at_expires"`
	RtExpires    int64  `json:"rt_expires"`
}

type Config struct {
	PORT           string
	API_PREFIX     string
	API_VERSION    string
	AWS_REGION     string
	AWS_SECRET     string
	AWS_SECRET_KEY string
	TABLE_NAME     string
	ACCESS_SECRET  string
	REFRESH_SECRET string
}
