package response

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadsarimin/indocyber-api/config"
	"github.com/muhammadsarimin/indocyber-api/models"
	"github.com/muhammadsarimin/indocyber-api/models/cerr"
	"gorm.io/gorm"
)

func Error(c *gin.Context, log *config.CustomLog, err error) {

	var statusCode int
	var res models.Response

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = cerr.GetError("003", "")
	}

	switch v := err.(type) {
	case *models.Error:
		statusCode = v.StatusCode
		res = models.Response{
			ResponseCode:    v.Code,
			ResponseMessage: v.Message,
		}
	case *json.UnmarshalTypeError:
		Error(c, log, cerr.GetError("001", v.Field))
		return
	default:
		statusCode = http.StatusInternalServerError
		res = models.Response{
			ResponseCode:    "500",
			ResponseMessage: "internal server error",
		}
	}

	payload, ok := c.Get("payload")
	if ok {
		log.Info("[ response ]", payload, res)
	} else {
		log.Info("[ response ]", nil, res)
	}

	c.JSON(statusCode, res)
}

func Success(c *gin.Context, log *config.CustomLog, data interface{}) {

	res := models.Response{
		ResponseCode:    "000",
		ResponseMessage: "Success",
		ResponseData:    data,
	}

	payload, ok := c.Get("payload")
	if ok {
		log.Info("[ response ]", payload, res)
	} else {
		log.Info("[ response ]", nil, res)
	}

	c.JSON(http.StatusOK, res)
}
