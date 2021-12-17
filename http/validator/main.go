package validator

import (
	"awsCloud/database/models"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

func SigninValidator(user *models.User) error {
	// remove spaces
	user.Username = strings.TrimSpace(user.Username)
	user.Password = strings.TrimSpace(user.Password)

	// validator translator
	translator := en.New()
	uni := ut.New(translator, translator)
	trans, found := uni.GetTranslator("en")
	if !found {
		log.Fatal("translator not found")
	}

	v := validator.New()
	if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
		log.Fatal(err)
	}

	// validate the required fields
	_ = v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	// validate the username field
	// first it will validate that username contains only letters & numbers
	_ = v.RegisterTranslation("isvalid", trans, func(ut ut.Translator) error {
		return ut.Add("isvalid", "{0} contains only characters & numbers.", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("isvalid", fe.Field())
		return t
	})
	_ = v.RegisterValidation("isvalid", func(fl validator.FieldLevel) bool {
		var IsLetter = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString
		return IsLetter(fl.Field().String())
	})
	// now it will validate the username length
	_ = v.RegisterTranslation("username", trans, func(ut ut.Translator) error {
		return ut.Add("username", "{0} length must be between 8 to 16.", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("username", fe.Field())
		return t
	})
	_ = v.RegisterValidation("username", func(fl validator.FieldLevel) bool {
		if len(fl.Field().String()) < 8 {
			return false
		} else if len(fl.Field().String()) > 16 {
			return false
		} else {
			return true
		}
	})

	// validate the password field
	_ = v.RegisterTranslation("passwrd", trans, func(ut ut.Translator) error {
		return ut.Add("passwrd", "{0} length must be between 6 to 12.", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("passwrd", fe.Field())
		return t
	})
	_ = v.RegisterValidation("passwrd", func(fl validator.FieldLevel) bool {
		if len(fl.Field().String()) < 6 {
			return false
		} else if len(fl.Field().String()) > 12 {
			return false
		} else {
			return true
		}
	})

	err := v.Struct(user)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			return fmt.Errorf(e.Translate(trans))
		}
	}

	return nil
}
