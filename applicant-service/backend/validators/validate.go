package validators

import (
	"applicant-service/models"
	"github.com/go-playground/validator/v10"
)

func ValidateExams(fl validator.FieldLevel) bool {
	exams, ok := fl.Field().Interface().(models.Exams)
	if !ok {
		return false
	}
	return exams.Physics > 0 || exams.Informatics > 0
}
