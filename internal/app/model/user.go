package model

import (
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

//User ...
type User struct {
  ID                int    `json:"id"`
  Email             string `json:"email"`
  Password          string `json:"password,onitempty"`
  EncryptedPassword string `json:"-"`
}

// Validate - валидация User
func (u *User) Validate() error {
  //поле Email обязательное и имеет формать эл.почты
  //Password обязательный, не меннее 6 и не более 100 символов
  return validation.ValidateStruct(
    u,
    validation.Field(&u.Email, validation.Required, is.Email),
    validation.Field(&u.Password, validation.By(requiredIf(u.EncryptedPassword == "")), validation.Length(6,100)),//By() - кастомная валидация, см. файл validations.go
  )
}

//BeforeCreate ...
func (u *User) BeforeCreate() error {
  //Если пароль не пустой, то хэшируем его
  if len(u.Password) > 0 {
    enc, err := encryptString(u.Password)
    if err != nil {
      return err
    }

    u.EncryptedPassword = enc

  }
  return nil
}

// Sanitaze - убираем данные о пароле, которые не нужно показывать
func (u *User) Sanitaze() {
  u.Password = ""
}

// encryptString - ширование строки
func encryptString(s string) (string, error) {
  b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
  if err != nil {
    return "", err
  }

  return string(b), nil
}
