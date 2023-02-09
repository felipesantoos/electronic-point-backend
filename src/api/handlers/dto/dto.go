package dto

import (
	"dit_backend/src/api/handlers/dto/request"
	"dit_backend/src/core/domain/errors"
	"dit_backend/src/core/helpers/validator"
	"dit_backend/src/core/utils"
	"encoding/json"
)

type Processable interface {
	request.Credentials |
		request.CreateAccount |
		request.UpdateAccountProfile |
		request.UpdatePassword |
		request.CreatePasswordReset |
		request.UpdatePasswordByPasswordReset
}

func Validate[T Processable](data interface{}) (map[string]interface{}, errors.Error) {
	if errs, ok := validator.ValidateDTO[T](data); !ok {
		return nil, errors.NewValidation(errs)
	}
	var instance map[string]interface{}
	dataStr, err := json.Marshal(utils.FormatJSONData(data))
	if err != nil {
		return nil, errors.New(err)
	}
	json.Unmarshal(dataStr, &instance)
	return instance, nil
}
