package utils

import (
	coreError "github.com/luancpereira/APICheckout/core/errors"
	"golang.org/x/crypto/bcrypt"
)

type Crypt struct{}

func (Crypt) MakeHash(value string) (hashedValue string, err error) {
	hashedValueByte, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		err = coreError.New("error.crypt.hash.value", err.Error())
		return
	}

	hashedValue = string(hashedValueByte)

	return
}

func (Crypt) Check(value string, hashValue string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hashValue), []byte(value))
	if err != nil {
		err = coreError.New("error.crypt.compare.values", err.Error())
		return
	}

	return
}
