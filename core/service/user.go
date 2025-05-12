package service

import (
	"context"
	"database/sql"
	"strings"
	"time"

	commonsUtils "github.com/luancpereira/APICheckout/apis/commons/utils"
	"github.com/luancpereira/APICheckout/core/database"
	"github.com/luancpereira/APICheckout/core/database/sqlc"
	"github.com/luancpereira/APICheckout/core/enums"
	coreError "github.com/luancpereira/APICheckout/core/errors"
	"github.com/luancpereira/APICheckout/core/utils"
)

type User struct{}

/*****
funcs for creations
******/

func (u User) Create(params sqlc.InsertUserParams, repeatPassword string) (ID int64, errorFields []coreError.CoreErrorField) {
	existsUserID, err := u.GetIDByEmail(params.Email)
	if err != nil && !strings.Contains(err.Error(), sql.ErrNoRows.Error()) {
		hasErrorUser := coreError.MakeErrorField(err, enums.USER_VALIDATION_FIELD_EMAIL, &errorFields)
		if hasErrorUser {
			return
		}
	}

	if existsUserID != 0 {
		coreErr := coreError.New("error.user.already.exists")
		coreError.MakeErrorField(coreErr, enums.USER_VALIDATION_FIELD_EMAIL, &errorFields)
		return
	}

	errorFields = u.validateToCreate(params, repeatPassword)
	if len(errorFields) > 0 {
		return
	}

	hashedPassword, err := utils.Crypt{}.MakeHash(params.Password)
	hasError := coreError.MakeErrorField(err, enums.USER_VALIDATION_FIELD_PASSWORD, &errorFields)
	if hasError {
		return
	}

	hashedTokenConfirmation, err := u.tokenConfirmationCreate()
	hasError = coreError.MakeErrorField(err, enums.USER_VALIDATION_FIELD_TOKEN, &errorFields)
	if hasError {
		return
	}

	params.Email = strings.TrimSpace(params.Email)
	params.Password = hashedPassword
	params.TokenConfirmation = hashedTokenConfirmation
	params.TokenConfirmationExpirationDate = time.Now().Add(24 * time.Hour)
	params.Permission = enums.USER_PERMISSION_TYPE_NORMAL
	params.CreatedAt = time.Now()

	model, err := database.TransactionReturnOneObject(func(querier sqlc.Querier) (model sqlc.InsertUserRow, err error) {
		model, err = querier.InsertUser(context.Background(), params)
		if err != nil {
			return
		}

		return
	})
	hasError = coreError.MakeErrorField(err, enums.USER_VALIDATION_FIELD_FORM, &errorFields)
	if hasError {
		return
	}

	ID = model.ID

	return
}

func (u User) Login(email, password string) (user sqlc.SelectUserForLoginRow, err error) {
	user, err = database.DB_QUERIER.SelectUserForLogin(context.Background(), email)
	if err != nil {
		err = database.Utils{}.CoreErrorDatabase(err)
		return
	}

	err = utils.Crypt{}.Check(password, user.Password)
	if err != nil {
		return
	}

	return
}

/*****
funcs for creations
******/

/*****
funcs for gets
******/

func (u User) GetIDByEmail(email string) (ID int64, err error) {
	ID, err = database.DB_QUERIER.SelectUserIDByEmail(context.Background(), email)
	if err != nil && !strings.Contains(err.Error(), sql.ErrNoRows.Error()) {
		err = database.Utils{}.CoreErrorDatabase(err)
		return
	}

	return
}

/*****
funcs for gets
******/

/*****
funcs for validations
******/

func (u User) validateToCreate(model sqlc.InsertUserParams, repeatPassword string) (errorFields []coreError.CoreErrorField) {
	err := u.validateEmail(model.Email, 0)
	coreError.MakeErrorField(err, enums.USER_VALIDATION_FIELD_EMAIL, &errorFields)

	err = u.validatePassword(model.Password, repeatPassword)
	coreError.MakeErrorField(err, enums.USER_VALIDATION_FIELD_PASSWORD, &errorFields)

	return
}

func (u User) validateEmail(email string, ID int64) (err error) {
	if !commonsUtils.ValidateEmail(email) {
		return coreError.New("error.validation.email.invalid")
	}

	existsID, err := u.GetIDByEmail(strings.TrimSpace(email))
	if err != nil && !strings.Contains(err.Error(), sql.ErrNoRows.Error()) {
		return database.Utils{}.CoreErrorDatabase(err)
	}

	if existsID != 0 && existsID != ID {
		return coreError.New("error.public.user.email.exists")
	}

	return nil
}

func (u User) validatePassword(password, repeatPassword string) (err error) {
	if password != repeatPassword {
		return coreError.New("error.public.user.password.mismatch")
	}

	if len(password) < enums.USER_PASSWORD_MIN_SIZE {
		return coreError.New("error.public.user.password.size")
	}

	expectedQuantity := 2
	if ok := commonsUtils.ValidateUpperCharacters(password, expectedQuantity); !ok {
		return coreError.New("error.public.user.password.uppers")
	}

	if ok := commonsUtils.ValidateSpecialCharacters(password, expectedQuantity); !ok {
		return coreError.New("error.public.user.password.special.characters")
	}

	return
}

func (u User) tokenConfirmationCreate() (confirmationTokenHashed string, err error) {
	confirmationToken := strings.ToUpper(commonsUtils.RandomStringAndNumber(5))
	confirmationTokenHashed, err = utils.Crypt{}.MakeHash(confirmationToken)

	return
}

/*****
funcs for validations
******/

/*****
other funcs
******/

/*****
other funcs
******/
