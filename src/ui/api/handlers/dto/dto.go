package dto

import (
	"backend_template/src/ui/api/handlers/dto/request"
	"backend_template/src/core/domain/errors"
	"backend_template/src/core/helpers/validator"
	"backend_template/src/core/utils"
	"encoding/json"
)

type Processable interface {
	request.Credentials |
		request.UpdatePassword |
		request.CreateAccount |
		request.UpdateAccountProfile |
		request.CreatePasswordReset |
		request.UpdatePasswordByPasswordReset
}

func Validate[T Processable](data interface{}) (*T, errors.Error) {
	if errs, ok := validator.ValidateDTO[T](data); !ok {
		return nil, errors.NewValidation(errs)
	}
	dataStr, err := json.Marshal(utils.FormatJSONData(data))
	if err != nil {
		return nil, errors.New(err)
	}
	var instance T
	json.Unmarshal(dataStr, &instance)
	return &instance, nil
}
