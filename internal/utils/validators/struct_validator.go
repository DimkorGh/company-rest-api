package validators

import "github.com/go-playground/validator/v10"

type StructValidatorInt interface {
	Validate(structForCheck interface{}) error
}

type StructValidator struct {
	goValidator *validator.Validate
}

func NewStructValidator(goValidator *validator.Validate) *StructValidator {
	return &StructValidator{
		goValidator: goValidator,
	}
}

func (sv *StructValidator) Validate(structForCheck interface{}) error {
	err := sv.goValidator.Struct(structForCheck)
	if err != nil {
		return &StructValidatorError{
			ErrorMessage: err.Error(),
		}
	}

	return nil
}

type StructValidatorError struct {
	ErrorMessage string
}

func (sve *StructValidatorError) Error() string {
	return sve.ErrorMessage
}
