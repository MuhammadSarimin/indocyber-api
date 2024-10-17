package cerr

import "github.com/muhammadsarimin/indocyber-api/models"

var errs = []models.Error{
	{
		StatusCode: 400,
		Code:       "001",
		Message:    "invalid field format",
	},
	{
		StatusCode: 400,
		Code:       "002",
		Message:    "invalid field mandatory",
	},
	{
		StatusCode: 400,
		Code:       "003",
		Message:    "invalid id",
	},
	{
		StatusCode: 404,
		Code:       "004",
		Message:    "data not found",
	},
	{
		StatusCode: 500,
		Code:       "004",
		Message:    "internal server error",
	},
}

func GetError(code string, field string) error {
	for _, err := range errs {
		if err.Code == code {
			if err.Code == "001" || err.Code == "002" {
				err.Message = err.Message + " " + field
			}
			return &err
		}
	}
	return &models.Error{
		StatusCode: 500,
		Code:       "05",
		Message:    "internal server error",
	}
}
