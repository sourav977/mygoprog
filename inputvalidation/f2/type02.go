package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"passwd"`
}

func main() {
	v := validator.New()
	_ = v.RegisterValidation("passwd", passwdValidator)

	a := User{
		Email:    "something@gmail.com",
		Name:     "A girl has no name",
		Password: "1234",
	}
	err := v.Struct(a)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println(e)
		}
	}

}

func passwdValidator(fl validator.FieldLevel) bool {
	return len(fl.Field().String()) > 6
}
